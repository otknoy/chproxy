package converter_test

import (
	"chproxy/app/converter"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertRequestPath_should_convert_dat_path_to_html_path(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			"/hoge/dat/1635648562.dat",
			"/test/read.cgi/hoge/1635648562",
		},
	}

	for _, tt := range tests {
		got := converter.ConvertRequestPath(tt.in)

		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Error(diff)
		}
	}
}
