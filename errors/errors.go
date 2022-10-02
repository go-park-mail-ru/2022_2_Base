package BaseErrors

import "errors"

var ErrBadRequest400 = errors.New("Bad request - Problem with the request")
var ErrUnauthorized401 = errors.New("Unauthorized - Access token is missing or invalid")
var ErrNotFound404 = errors.New("Not found - Requested entity is not found in database")
var ErrConflict409 = errors.New("Conflict - User already exists")
var ErrServerError500 = errors.New("Internal Server Error - Request is valid but operation failed at server side")
