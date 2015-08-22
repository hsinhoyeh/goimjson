package http

import (
	"log"

	"github.com/emicklei/go-restful"
)

// LogHandler wrap a ServiceHandler to track each request and response via log
type LogHandler struct {
	s *ServiceHandler
}

// NewLogHandler new a service with logger
func NewLogHandler(s *ServiceHandler) *LogHandler {
	return &LogHandler{s: s}
}

// Get handles the Get request
func (l *LogHandler) Get(req *restful.Request, resp *restful.Response) {
	l.logHeader(req, resp)
	l.s.Get(req, resp)
}

// Post handles the Post request
func (l *LogHandler) Post(req *restful.Request, resp *restful.Response) {
	l.logHeader(req, resp)
	l.s.Post(req, resp)
}

func (l *LogHandler) logHeader(req *restful.Request, resp *restful.Response) {
	log.Printf("%s %s %s\n",
		req.Request.RemoteAddr,
		req.Request.Method,
		req.Request.URL)
}
