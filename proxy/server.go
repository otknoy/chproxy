package proxy

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func NewServer() *http.Server {
	director := func(req *http.Request) {
		log.Printf("%#v\n", req)
	}

	modifier := func(res *http.Response) error {
		log.Printf("%#v\n", res)

		s := "sample"

		res.Body = io.NopCloser(strings.NewReader(s))
		res.ContentLength = int64(len(s))

		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	return &http.Server{
		Addr:    ":9000",
		Handler: rp,
	}
}
