package proxy

import (
	"chproxy/app/converter"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func NewServer() *http.Server {
	director := NewDirectory()
	modifier := NewModifyResponse()

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	return &http.Server{
		Addr:    ":9000",
		Handler: rp,
	}
}

func NewDirectory() func(*http.Request) {
	return func(req *http.Request) {
		// log.Printf("%#v\n", req.URL)

		if datRequest(req) {
			req.URL.Scheme = "http"
			req.URL.Path = converter.ConvertRequestPath(req.URL.Path)

			log.Printf("%#v\n", req)
		}
	}
}

func NewModifyResponse() func(*http.Response) error {
	return func(res *http.Response) error {
		log.Printf("%#v\n", res)
		log.Printf("%#v\n", res.Request.URL.String())

		if datProxyRequest(res.Request) {
			res.Body = io.NopCloser(converter.ConvertResponse(res.Body))

			log.Printf("%#v\n", res)
		}

		return nil
	}
}

func datRequest(req *http.Request) bool {
	return strings.HasSuffix(req.URL.Host, "5ch.net") && strings.HasSuffix(req.URL.Path, ".dat")
}

func datProxyRequest(req *http.Request) bool {
	return strings.HasSuffix(req.URL.Host, "5ch.net") && strings.HasPrefix(req.URL.Path, "/test/read.cgi/")
}
