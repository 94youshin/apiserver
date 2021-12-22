package util_test

import (
	"github.com/youshintop/apiserver/pkg/util"
	"testing"
)

func TestGenShortId(t *testing.T) {
	shortId, err := util.GenShortId()
	if shortId == "" || err != nil {
		t.Error("failed GenShortId")
	}
	t.Log("GenShortId paas")
}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortId, err := util.GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StopTimer()

	for i := 0; i < b.N; i++ {
		util.GenShortId()
	}
}
