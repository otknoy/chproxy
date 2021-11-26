package converter_test

import (
	"bytes"
	"chproxy/app/converter"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestConvertResponse(t *testing.T) {
	html := loads(t, filepath.Join("testdata", "sample.html"))
	dat := loads(t, filepath.Join("testdata", "sample.dat"))

	got := converter.ConvertResponse(html)

	if diff := cmp.Diff(
		readerToString(dat),
		readerToString(got),
	); diff != "" {
		t.Error(diff)
	}
}

func loads(t *testing.T, filename string) io.Reader {
	t.Helper()

	f, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	return bytes.NewReader(f)
}

func readerToString(r io.Reader) string {
	b, _ := io.ReadAll(
		transform.NewReader(r, japanese.ShiftJIS.NewDecoder()),
	)

	return string(b)
}
