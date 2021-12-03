package model_test

import (
	"chproxy/app/model"
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
				Date:  "2021/11/30(火) 21:41:25",
				ID:    "author1-id",
				BE:    "author1-be",
				Text:  "foobar",
			},
			{
				Name:  "author2",
				EMail: "author2@example.com",
				Date:  "2021/11/30(火) 21:42:00",
				ID:    "author2-id",
				Text:  "改行は\n変換する",
			},
		},
	}

	got := thread.ToDat()

	want := `author1<>author1@example.com<>2021/11/30(火) 21:41:25 ID:author1-id BE:author1-be<>foobar<>test-title
author2<>author2@example.com<>2021/11/30(火) 21:42:00 ID:author2-id<>改行は <br> 変換する<>
あぼーん<>あぼーん<>あぼーん<>あぼーん<>
`

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differ.\n%v\n", diff)
	}
}
