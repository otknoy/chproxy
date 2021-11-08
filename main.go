package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	director := func(req *http.Request) {
		log.Printf("%#v\n", req)
	}

	modifier := func(res *http.Response) error {
		log.Printf("%#v\n", res)

		s := "sample"

		res.Body = io.NopCloser(strings.NewReader("sample"))
		res.ContentLength = int64(len(s))

		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	server := http.Server{
		Addr:    ":9000",
		Handler: rp,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
