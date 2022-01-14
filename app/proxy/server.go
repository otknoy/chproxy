package proxy

import (
	"chproxy/app/converter"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
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

		req.URL.Scheme = "http"
		req.URL.Path = converter.ConvertRequestPath(req.URL.Path)

		log.Printf("%#v\n", req)
	}
}

func NewModifyResponse() func(*http.Response) error {
	return func(res *http.Response) error {
		res.Body = io.NopCloser(converter.ConvertResponse(res.Body))

		// log.Printf("%#v\n", res)

		// s := "sample"

		// res.Body = io.NopCloser(strings.NewReader(s))
		// res.ContentLength = int64(len(s))

		return nil
	}
}
