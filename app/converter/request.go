package converter

import (
	"strings"
)

func ConvertRequestPath(path string) string {
	board, thread := func() (string, string) {
		l := strings.Split(path, "/")
		return l[1], strings.Split(l[3], ".")[0]
	}()

	return "/test/read.cgi/" + board + "/" + thread
}
