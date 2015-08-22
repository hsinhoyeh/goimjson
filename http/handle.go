package http

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful"
	imj "github.com/hsinhoyeh/goimjson"
)

const (
	root = "root"
)

// Service defines an interface for WebService Handler
type Handler interface {
	Get(request *restful.Request, response *restful.Response)
	Post(request *restful.Request, response *restful.Response)
}

func New(h Handler) *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// TODO: partial query
	service.Route(service.GET("/{version}").To(h.Get))
	service.Route(service.POST("/").To(h.Post))
	return service
}

type docType map[string]interface{}

// ServiceHandler is a http service handlers for goimjson
type ServiceHandler struct {
	j *imj.ImJSON
}

// New allocates an instance of Service
func NewHandler(j *imj.ImJSON) *ServiceHandler {
	return &ServiceHandler{j: j}
}

// Get handles the Get request
func (s *ServiceHandler) Get(request *restful.Request, response *restful.Response) {
	// read path and version
	ver := request.PathParameter("version")

	data, err := s.j.Get(ver, root)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	if data == nil && err == nil {
		response.WriteError(http.StatusNotFound, errors.New("[imjsom] not found"))
		return
	}
	response.WriteEntity(data.Interface())
}

type postResponse struct {
	Ver string `json:"ver"`
}

// Post handles the Post request
func (s *ServiceHandler) Post(request *restful.Request, response *restful.Response) {
	d := &docType{}
	err := request.ReadEntity(d)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	ver := s.j.Set(root, d)
	if ver == imj.InvalidVersion {
		response.WriteError(http.StatusInternalServerError, errors.New("[imjson] invalid version"))
		return
	}
	response.WriteEntity(postResponse{ver})
}
