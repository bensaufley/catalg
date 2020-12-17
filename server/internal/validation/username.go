package validation

func Username(username string) *Error {
	if username == "" {
		return NewError("username", "cannot be blank")
	}

	return nil
}
