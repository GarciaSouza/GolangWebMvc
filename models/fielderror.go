package models

//FieldError Custom field error
type FieldError struct {
	Err       error
	FieldName string
}

func (fe *FieldError) Error() string {
	return fe.FieldName + ": " + fe.Err.Error()
}
