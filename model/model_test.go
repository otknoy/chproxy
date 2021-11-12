package model_test

import (
	"chproxy/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestThread(t *testing.T) {
	thread := &model.Thread{
		Title: "test-title",
		ResList: []model.Res{
			{
				Name:  "author1",
				EMail: "author1@example.com",
				Date:  "",
				ID:    "author1-id",
				Text:  "foobar",
			},
			{
				Name:  "author2",
				EMail: "author2@example.com",
				Date:  "",
				ID:    "author2-id",
				Text:  "hogehoge",
			},
		},
	}

	got := thread.ToDat()

	want := `author1<>author1@example.com<>ID:author1-id<>foobar<>test-title
author2<>author2@example.com<>ID:author2-id<>hogehoge<>
あぼーん<>あぼーん<>あぼーん<>あぼーん<>
`

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differ.\n%v\n", diff)
	}
}
