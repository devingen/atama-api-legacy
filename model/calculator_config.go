package model

type CalculatorConfig struct {
	// field name that contains a unique id (id, email etc.)
	FieldID string

	// field name that contains the limit of item
	FieldLimit string

	// field name that contains the data of the first item type
	//FirstDataGroupIndex string

	// field name that contains the data of the second item type
	//SecondDataGroupIndex string
}
