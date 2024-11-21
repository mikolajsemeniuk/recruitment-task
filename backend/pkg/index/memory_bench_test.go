package index_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/mikolajsemeniuk/recruitment-task/pkg/index"
)

func BenchmarkFindExactMatch(b *testing.B) {
	var builder strings.Builder

	size := 1000000
	data := make([]int, size)

	for i := range size {
		val := i * 10
		data[i] = val
		builder.WriteString(fmt.Sprintf("%d\n", val))
	}

	fs := fstest.MapFS{"input.txt": &fstest.MapFile{Data: []byte(builder.String())}}

	store, err := index.NewMemory(fs, "input.txt")
	if err != nil {
		b.Fatalf("Failed to create Memory: %v", err)
	}

	value := data[len(data)/2]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = store.Find(context.TODO(), value)
	}
}

func BenchmarkFindNearMiss(b *testing.B) {
	var builder strings.Builder

	size := 1000000
	data := make([]int, size)

	for i := range size {
		val := i * 10
		data[i] = val
		builder.WriteString(fmt.Sprintf("%d\n", val))
	}

	fs := fstest.MapFS{"input.txt": &fstest.MapFile{Data: []byte(builder.String())}}

	store, err := index.NewMemory(fs, "input.txt")
	if err != nil {
		b.Fatalf("Failed to create Memory: %v", err)
	}

	value := data[len(data)/2] + int(float64(data[len(data)/2])*0.05)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = store.Find(context.TODO(), value)
	}
}

func BenchmarkFindMiss(b *testing.B) {
	var builder strings.Builder

	size := 1000000
	data := make([]int, size)

	for i := range size {
		val := i * 10
		data[i] = val
		builder.WriteString(fmt.Sprintf("%d\n", val))
	}

	fs := fstest.MapFS{"input.txt": &fstest.MapFile{Data: []byte(builder.String())}}

	store, err := index.NewMemory(fs, "input.txt")
	if err != nil {
		b.Fatalf("Failed to create Memory: %v", err)
	}

	value := data[len(data)-1] + 100000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = store.Find(context.TODO(), value)
	}
}
