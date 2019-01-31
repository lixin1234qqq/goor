package orhttp

import (
	"net/http"

	"github.com/baidu/gls"
	"github.com/baidu/goid"
)

// Wrap returns an http.Handler wrapping h
func Wrap(h http.Handler) http.Handler {
	if h == nil {
		panic("h == nil")
	}
	handler := &handler{
		handler: h,
	}
	return handler
}

// handler wraps an http.Handler
type handler struct {
	handler http.Handler
}

// ServeHTTP delegates to h.Handler
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	cGoId := goid.GoIDAsm()
	gls.SetGls(cGoId, make(map[interface{}]interface{}))
	defer func() {
		gls.RemoveGls(cGoId)
	}()
	gls.Set("request", req)
	h.handler.ServeHTTP(w, req)
}
