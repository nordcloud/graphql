package graphql

const (
	ErrorTypeUnauthorized     string = "UnauthorizedException"
	ErrorTypeUnknownOperation string = "UnknownOperationException"
)

type AppsyncErrorType uint32

const (
	ErrAppsyncUnknown AppsyncErrorType = iota
	ErrAppsyncUnauthorized
	ErrAppsyncUnknownOperation
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

func AppsyncErr(err error) AppsyncErrorType {
	gqlErr, ok := err.(GraphQLError)
	if !ok {
		return ErrAppsyncUnknown
	}

	if gqlErr.ErrorType != nil {
		switch *gqlErr.ErrorType {
		case ErrorTypeUnauthorized:
			return ErrAppsyncUnauthorized
		case ErrorTypeUnknownOperation:
			return ErrAppsyncUnknownOperation
		}
	}

	return ErrAppsyncUnknown
}
