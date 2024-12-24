package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCategoryID = errors.New("invalid category ID")
