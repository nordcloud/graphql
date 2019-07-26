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
	errType AppsyncErrorType
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
	is.Equal(AppsyncErr(err), tc.errType)
}

func TestAppsyncErr(t *testing.T) {
	is := is.New(t)

	for _, tc := range []testCase{
		{`{"message": "error", "errorType": "UnauthorizedException"}`, ErrAppsyncUnauthorized},
		{`{"message": "error", "errorType": "UnknownOperationException"}`, ErrAppsyncUnknownOperation},
		{`{"message": "error", "errorType": "foobar"}`, ErrAppsyncUnknown},
		{`{"message": "error"}`, ErrAppsyncUnknown},
	} {
		tc.run(is)
	}
}
