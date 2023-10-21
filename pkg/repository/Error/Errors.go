package Error

import "errors"

var (
	RecordNotFound  = errors.New("record not found")
	RecordNotCreate = errors.New("failed to create entity")
)
