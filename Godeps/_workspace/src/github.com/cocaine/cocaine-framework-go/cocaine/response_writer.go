package cocaine

import (
	"net/http"
	"strconv"
)

// ResponseWriter implements http.ResponseWriter interface. It implements cocaine integration.
type ResponseWriter struct {
	cRes         *Response
	req          *http.Request
	handlerHeader http.Header
	written       int64 // number of bytes written in body
	contentLength int64 // explicitly-declared Content-Length; or -1
	status        int   // status code passed to WriteHeader
	wroteHeader	  bool
	logger		  *Logger
}


func (w *ResponseWriter) Header() http.Header {
	return w.handlerHeader
}

func (w *ResponseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		w.logger.Err("http: multiple response.WriteHeader calls")
		return
	}
	w.wroteHeader = true
	w.status = code

	if cl := w.handlerHeader.Get("Content-Length"); cl != "" {
		v, err := strconv.ParseInt(cl, 10, 64)
		if err == nil && v >= 0 {
			w.contentLength = v
		} else {
			w.logger.Errf("http: invalid Content-Length of %q", cl)
			w.handlerHeader.Del("Content-Length")
		}
	}
	w.cRes.Write(WriteHead(code, HttpHeaderToCocaineHeader(w.handlerHeader)))
}

func (w *ResponseWriter) finishRequest() {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	if w.req.MultipartForm != nil {
		w.req.MultipartForm.RemoveAll()
	}

}

// bodyAllowed returns true if a Write is allowed for this response type.
// It's illegal to call this before the header has been flushed.
func (w *ResponseWriter) bodyAllowed() bool {
	if !w.wroteHeader {
		panic("")
	}
	return w.status != http.StatusNotModified
}

func (w *ResponseWriter) Write(data []byte) (n int, err error) {
	return w.write(len(data), data, "")
}

func (w *ResponseWriter) WriteString(data string) (n int, err error) {
	return w.write(len(data), nil, data)
}

// either dataB or dataS is non-zero.
func (w *ResponseWriter) write(lenData int, dataB []byte, dataS string) (n int, err error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	if lenData == 0 {
		return 0, nil
	}
	if !w.bodyAllowed() {
		return 0, http.ErrBodyNotAllowed
	}

	w.written += int64(lenData) // ignoring errors, for errorKludge
	if w.contentLength != -1 && w.written > w.contentLength {
		return 0, http.ErrContentLength
	}
	if dataB != nil {
		w.cRes.Write(dataB)
		return len(dataB), nil
	} else {
		w.cRes.Write(dataS)
		return len(dataS), nil
	}
}
