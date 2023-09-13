package user

import "errors"

var ErrRegisterFailedToHashPassword = errors.New("failed to hash your password")

var ErrRegisterAccountAlreadyExists = errors.New("email already registered")

var ErrRegisterFailedToRegisterAccount = errors.New("failed to register your account")
