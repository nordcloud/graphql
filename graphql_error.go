package graphql

const (
	ErrTypeUnauthorized     string = "UnauthorizedException"
	ErrTypeUnknownOperation string = "UnknownOperationException"
)

type GraphQLError struct {
	Message   string  `json:"message"`
	ErrorType *string `json:"errorType"`
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

func ErrorType(err error) string {
	gqlErr, ok := err.(GraphQLError)
	if !ok {
		return ""
	}

	return gqlErr.Type()
}
