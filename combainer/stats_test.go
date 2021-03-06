package combainer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	c1 := &clientStats{}

	c1.AddSuccessAggregate()
	assert.EqualValues(t, c1.successAggregate, 1)
	c1.AddFailedAggregate()
	assert.EqualValues(t, c1.failedAggregate, 1)
	stats := c1.GetStats()
	assert.EqualValues(t, stats.AggregateTotal, 2)

	c1.AddSuccessParsing()
	assert.EqualValues(t, c1.successParsing, 1)
	c1.AddFailedParsing()
	assert.EqualValues(t, c1.failedParsing, 1)
	stats = c1.GetStats()
	assert.EqualValues(t, stats.ParsingTotal, 2)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c1.AddSuccessParsing()
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			c1.AddSuccessParsing()
			wg.Done()
		}()
		c1.AddSuccessParsing()

		go func() {
			c1.AddFailedParsing()
			wg.Done()
		}()
		c1.AddFailedParsing()
	}
	wg.Wait()

	stats = c1.GetStats()
	assert.EqualValues(t, stats.ParsingTotal, 4012)

	c2 := &clientStats{}
	c1.CopyStats(c2)
	c2.last = stats.Heartbeated
	assert.EqualValues(t, stats, c2.GetStats())
}
