package graphql

import "fmt"

const (
	ErrTypeUnauthorized     string = "UnauthorizedException"
	ErrTypeUnknownOperation string = "UnknownOperationException"
)

type GraphQLError struct {
	Message   string  `json:"message"`
	ErrorType *string `json:"errorType"`
	HttpCode  *int
}

func (e GraphQLError) Error() string {
	return "graphql: " + e.Message
}

func (e GraphQLError) Type() string {
	if e.ErrorType == nil {
		return ""
	}
	return *e.ErrorType
}

func (e GraphQLError) Code() int {
	if e.HttpCode == nil {
		return 0
	}
	return *e.HttpCode
}

func ErrorType(err error) string {
	gqlErr, ok := err.(GraphQLError)
	if !ok {
		return ""
	}

	return gqlErr.Type()
}

func ErrorHttpCode(err error) int {
	gqlErr, ok := err.(GraphQLError)
	if !ok {
		return 0
	}

	return gqlErr.Code()
}

func newGraphQLErrWithCode(code int) GraphQLError {
	return GraphQLError{
		Message:  fmt.Sprintf("server returned a non-200 status code: %v", code),
		HttpCode: &code,
	}
}
