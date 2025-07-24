package platform

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mwc *ComponentContext) Error(err error) {
	mwc.error = err

}
func (mwc *ComponentContext) GetError() error {
	return mwc.error

}

type MiddlewareComponent interface {
	Init()

	ProcessRequest(w http.ResponseWriter, r *http.Request)
}
