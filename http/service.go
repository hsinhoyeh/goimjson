package http

import (
	"net/http"

	"github.com/emicklei/go-restful"
	imj "github.com/hsinhoyeh/goimjson"
)

// ListenAndServe listen the given address
func ListenAndServe(addr string) error {
	imjobj, err := imj.New()
	if err != nil {
		return err
	}
	restful.Add(
		New(
			NewLogHandler(
				NewHandler(
					imjobj,
				),
			),
		),
	)
	if err = http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
