package internal

import (
	"fmt"

	"github.com/openaq/openaq-go"
)

type APIKeyNotSetError struct{}

func (e *APIKeyNotSetError) Error() string {
	return "API Key not set, run `openaq settings set api-key [api-key]` to set value"
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "Your API Key is in invalid"
}

type ForbiddenError struct{}

func (e *ForbiddenError) Error() string {
	return "Forbidden: you are not allowed to access this resource"
}

type TooManyRequestsError struct{}

func (e *TooManyRequestsError) Error() string {
	return "Too Many Requests: You have exceeded the rate limit"
}

type InternalServerError struct{}

func (e *InternalServerError) Error() string {
	return "Something went wrong with the OpenAQ API"
}

func ErrorCheck(err error) error {

	switch v := err.(type) {
	case *openaq.APIError:
		switch errorCode := v.Code; errorCode {
		case 401:
			return &InvalidCredentialsError{}
		case 403:
			return &ForbiddenError{}
		case 429:
			return &TooManyRequestsError{}
		case 500:
			return &InternalServerError{}
		}
	default:
		fmt.Println(err.(*openaq.APIError).Code)

	}
	return err
}
