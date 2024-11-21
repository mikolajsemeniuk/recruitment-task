package index_test

import (
	"context"
	"errors"
	"testing"
	"testing/fstest"

	"github.com/mikolajsemeniuk/recruitment-task/pkg/index"
)

func TestMemory_Find(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		data  string
		value int
		index int
		err   error
	}{
		{
			name:  "Exact match at beginning",
			data:  "0\n3\n4\n",
			value: 0,
			index: 0,
			err:   nil,
		},
		{
			name:  "Exact match in the middle",
			data:  "0\n10\n20\n100\n1100\n1200\n2000\n2500\n",
			value: 1100,
			index: 4,
			err:   nil,
		},
		{
			name:  "Exact match at the end",
			data:  "0\n2500\n",
			value: 2500,
			index: 1,
			err:   nil,
		},
		{
			name:  "Match within lower range",
			data:  "0\n10\n2700\n",
			value: 3000,
			index: 2,
			err:   nil,
		},
		{
			name:  "Match within upper range",
			data:  "2700\n",
			value: 3000,
			index: 0,
			err:   nil,
		},
		{
			name:  "Match within one element upper range",
			data:  "3200\n3700\n",
			value: 3600,
			index: 1,
			err:   nil,
		},
		{
			name:  "Not match within lower range",
			data:  "0\n10\n2699\n",
			value: 3000,
			index: 0,
			err:   index.ErrIndexNotFound,
		},
		{
			name:  "Not match within upper range",
			data:  "3301\n",
			value: 3000,
			index: 0,
			err:   index.ErrIndexNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := "input.txt"
			f := fstest.MapFS{
				path: &fstest.MapFile{
					Data: []byte(c.data),
				},
			}

			m, err := index.NewMemory(f, path)
			if err != nil {
				t.Fatalf("Failed to create Memory: %v", err)
			}

			index, err := m.Find(context.TODO(), c.value)
			if index != c.index {
				t.Errorf("index got: %v, want: %v", index, c.index)
			}

			if !errors.Is(err, c.err) {
				t.Errorf("err got: %v, want %v", err, c.err)
			}
		})
	}
}
