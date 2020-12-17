package validation

import "strings"

type Error struct {
	fieldErrors map[string][]string
}

func NewError(field string, errors ...string) *Error {
	return &Error{
		fieldErrors: map[string][]string{
			field: errors,
		},
	}
}

func CollectErrors(errs ...*Error) *Error {
	var err *Error
	for _, e := range errs {
		err = err.Merge(e)
	}
	return err
}

func (v *Error) Error() string {
	msg := strings.Builder{}
	msg.WriteString("there were validation errors: ")
	for fieldName, errors := range v.fieldErrors {
		msg.WriteString(fieldName + " ")
		end := len(errors) - 1
		for i, err := range errors {
			msg.WriteString(err)
			if i == end-2 {
				msg.WriteString(" and ")
			} else if i < end {
				msg.WriteString(", ")
			}
		}
	}
	return msg.String()
}

func (v *Error) Errors() map[string][]string {
	return v.fieldErrors
}

func (v *Error) Merge(v2 *Error) *Error {
	if v == nil && v2 == nil {
		return nil
	}
	if v == nil {
		return v2
	}
	if v2 == nil {
		return v
	}
	err := *v
	for field, errors := range v2.fieldErrors {
		if errs, ok := err.fieldErrors[field]; ok {
			for _, msg := range errors {
				found := false
			DedupeLoop:
				for _, m := range errs {
					if m == msg {
						found = true
						break DedupeLoop
					}
				}
				if !found {
					err.fieldErrors[field] = append(err.fieldErrors[field], msg)
				}
			}
		} else {
			err.fieldErrors[field] = errors
		}
	}
	return &err
}

func AllowNil(i interface{}, cb func(interface{}) *Error) *Error {
	if i == nil {
		return nil
	}
	return cb(i)
}

type Validator func(interface{}) *Error
