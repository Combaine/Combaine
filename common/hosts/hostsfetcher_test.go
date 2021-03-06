package hosts

import (
	"sort"
	"testing"

	"github.com/combaine/combaine/utils"
	"github.com/stretchr/testify/assert"
)

func TestCommonHostsUtil(t *testing.T) {
	myname := utils.Hostname()
	hosts := Hosts{"DC1": {"host1", myname}}

	assert.Equal(t, hosts["DC1"], hosts.AllHosts())
	assert.NotContains(t, hosts.RemoteHosts(), myname)

	local := Hosts{"DC1": {myname}}
	remote := Hosts{"DC1": {"host1"}}
	local.Merge(&remote)
	lHosts := local.AllHosts()
	sort.Strings(hosts["DC1"])
	sort.Strings(lHosts)
	assert.EqualValues(t, hosts["DC1"], lHosts)
}
