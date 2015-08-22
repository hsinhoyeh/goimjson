package http

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	hasLog = flag.Bool("has_log", false, "has log to stdout")
)

type server struct {
	port int
	stop chan struct{}
}

// Stop stops server
func (s *server) Stop() {
	s.stop <- struct{}{}
}

func RunTestServer(t *testing.T) *server {
	srv := &server{
		port: 9000,
		stop: make(chan struct{}, 1),
	}
	// fork a goroutine to start server
	go func(srv *server) {
		ListenAndServe(fmt.Sprintf(":%d", srv.port))
		// TODO: stop service
	}(srv)

	time.Sleep(5 * time.Second)

	return srv
}

func TestHTTP(t *testing.T) {
	// start http server
	serv := RunTestServer(t)

	postData := map[string]interface{}{
		"foo": "foo",
		"bar": "bar",
	}

	marshaledPostData, _ := json.Marshal(postData)

	// Post http://localhost:9000
	// -d'{"foo":"foo","bar":"bar"}'
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://localhost:%d", serv.port),
		bytes.NewBuffer(marshaledPostData))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := httpDo(req)

	AssertStatusOk(t, resp)
	respBytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	postResp := &postResponse{}
	json.Unmarshal(respBytes, postResp)

	// Get http://localhost:9000/$ver
	req, _ = http.NewRequest(
		"GET",
		fmt.Sprintf("http://localhost:%d/%s", serv.port, postResp.Ver),
		nil)
	resp, _ = httpDo(req)
	respBytes, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	gotData := map[string]interface{}{}
	json.Unmarshal(respBytes, &gotData)

	assert.True(t, reflect.DeepEqual(postData, gotData))

	serv.Stop()
}

func httpDo(req *http.Request) (*http.Response, error) {
	if !*hasLog {
		return http.DefaultClient.Do(req)
	}

	reqBlob, _ := httputil.DumpRequest(req, true)
	fmt.Println("<----request")
	fmt.Println(string(reqBlob))
	fmt.Println("----")

	resp, err := http.DefaultClient.Do(req)

	respBlob, _ := httputil.DumpResponse(resp, true)
	fmt.Println("---->response")
	fmt.Println(string(respBlob))
	fmt.Println("----")

	return resp, err
}

func AssertStatusOk(t *testing.T, resp *http.Response) {
	assert.Equal(t, 200, resp.StatusCode)
}
