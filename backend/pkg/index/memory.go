package index

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"sort"
	"strconv"
)

// Memory defines datastore in memory.
type Memory struct {
	data []int
}

// NewMemory creates a new datastore with reading data from file.
func NewMemory(f fs.FS, path string) (*Memory, error) {
	// I decided to use fs.FS with "path" interface instead of only "path" to make it easier for unit testing.
	// Otherwise in unit tests new file has to be created and deleted after every test.
	file, err := f.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	var data []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("cannot parse number from file: %w", err)
		}

		data = append(data, int(value))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return &Memory{data: data}, nil
}

// FindIndex read index by value from memory.
func (m *Memory) Find(_ context.Context, value int) (int, error) {
	// I decided to pass context.Context even is not used because is a
	// good practice to pass context when we do calls to external resources,
	// or when we want to trace function execution and other details.
	//
	// I assume this logic would be done via some database engine.
	// That's why I added it here and not in handler.
	// Usually if here would be an sql engine I would write SQL query below instead of logic below.
	count := len(m.data)

	index := sort.Search(count, func(i int) bool {
		return m.data[i] >= value
	})

	if index < count && m.data[index] == value {
		return index, nil
	}

	margin := 0.10
	bound := int(float64(value) * margin)
	min := value - bound
	max := value + bound

	if index > 0 && m.data[index-1] >= min {
		return index - 1, nil
	}

	if index < count && m.data[index] <= max {
		return index, nil
	}

	return 0, ErrIndexNotFound
}
