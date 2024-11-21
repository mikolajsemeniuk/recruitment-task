package index_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mikolajsemeniuk/recruitment-task/pkg/index"
)

type storage struct {
	find func(ctx context.Context, value int) (int, error)
}

func (s *storage) Find(ctx context.Context, value int) (int, error) {
	return s.find(ctx, value)
}

func TestHandler_Find(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		value          string
		storageFunc    func(ctx context.Context, value int) (int, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Ok",
			value:          "123",
			storageFunc:    func(_ context.Context, value int) (int, error) { return value, nil },
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"index\":123}\n",
		},
		{
			name:           "Invalid value in URL",
			value:          "abc",
			storageFunc:    func(_ context.Context, _ int) (int, error) { return 0, nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "cannot parse value: strconv.ParseInt: parsing \"abc\": invalid syntax\n",
		},
		{
			name:           "Negative value",
			value:          "-1",
			storageFunc:    func(_ context.Context, _ int) (int, error) { return 0, nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "value cannot be negative\n",
		},
		{
			name:           "Index not found",
			value:          "123",
			storageFunc:    func(_ context.Context, _ int) (int, error) { return 0, index.ErrIndexNotFound },
			expectedStatus: http.StatusNotFound,
			expectedBody:   "index not found\n",
		},
		{
			name:  "Storage internal error",
			value: "123",
			storageFunc: func(_ context.Context, _ int) (int, error) {
				return 0, errors.New("internal error") //nolint:err113
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "internal error\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			store := &storage{find: c.storageFunc}

			req := httptest.NewRequest(http.MethodGet, "/"+c.value, nil)

			recorder := httptest.NewRecorder()

			handler := index.NewHandler(store)

			handler.ServeHTTP(recorder, req)

			if recorder.Code != c.expectedStatus {
				t.Errorf("expected status want: %d, got: %d", c.expectedStatus, recorder.Code)
			}

			body := recorder.Body.String()
			if body != c.expectedBody {
				t.Errorf("expected body: %q, got: %q", c.expectedBody, body)
			}
		})
	}
}
