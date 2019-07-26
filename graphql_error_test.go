package graphql

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/matryer/is"
)

type testCase struct {
	errBody string
	errType string
}

func (tc testCase) run(is *is.I) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf(`{"errors": [%s]}`, tc.errBody))
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := NewClient(srv.URL)

	var responseData map[string]interface{}
	err := client.Run(ctx, &Request{q: "query {}"}, &responseData)
	is.True(err != nil)
	is.Equal(ErrorType(err), tc.errType)
}

func TestAppsyncErr(t *testing.T) {
	is := is.New(t)

	for _, tc := range []testCase{
		{`{"message": "error", "errorType": "UnauthorizedException"}`, ErrTypeUnauthorized},
		{`{"message": "error", "errorType": "UnknownOperationException"}`, ErrTypeUnknownOperation},
		{`{"message": "error", "errorType": "foobar"}`, "foobar"},
		{`{"message": "error"}`, ""},
	} {
		tc.run(is)
	}
}
