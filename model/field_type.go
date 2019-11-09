package model

type FieldType string

const (
	FieldTypeNumber         FieldType = "number"
	FieldTypeSingleChoice   FieldType = "singleChoice"
	FieldTypeMultipleChoice FieldType = "multipleChoice"
	FieldTypeText           FieldType = "text"
)
