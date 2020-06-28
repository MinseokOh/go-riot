package lol

import "errors"

//https://developer.riotgames.com/docs/portal#web-apis_response-codes
func Error(code int) (err error) {
	switch code {
	case 400:
		err = NewError("Bad Request")
		break
	case 401:
		err = NewError("Unauthorized")
		break
	case 403:
		err = NewError("Forbidden")
		break
	case 404:
		err = NewError("Not Found")
		break
	case 415:
		err = NewError("Unsupported Media Type")
		break
	case 429:
		err = NewError("Rate Limit Exceeded")
		break
	case 500:
		err = NewError("Internal Server Error")
		break
	case 503:
		err = NewError("Service Unavailable")
		break
	}

	return nil
}

func NewError(msg string) (err error) {
	err = errors.New(msg)
	return
}
