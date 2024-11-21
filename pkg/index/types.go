package index

import (
	"fmt"
	"net/http"
	"strconv"
)

type FindInput struct {
	Number int
}

func NewFindInput(r *http.Request) (FindInput, error) {
	value, err := strconv.ParseInt(r.PathValue("value"), 10, 32)
	if err != nil {
		return FindInput{}, fmt.Errorf("cannot parse value: %w", err)
	}

	// I decided to use int and not uint for parsing non-negative values to give user more meaningful error message.
	if value < 0 {
		return FindInput{}, ErrValueNegative
	}

	return FindInput{Number: int(value)}, nil
}

type FindOutput struct {
	Index int `json:"index"`
}
