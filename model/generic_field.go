package model

/* GenericFieldOption */
type GenericFieldOption map[string]interface{}

func (option GenericFieldOption) GetLabel() string {
	return option["label"].(string)
}
func (option GenericFieldOption) GetValue() interface{} {
	return option["value"]
}

/* GenericField */
type GenericField map[string]interface{}

func (item GenericField) GetID() string {
	return item["id"].(string)
}

func (item GenericField) GetType() string {
	return item["type"].(string)
}

func (item GenericField) HasOptions() bool {
	_, hasOptions := item["options"].([]interface{})
	return hasOptions
}

func (item GenericField) GetOptions() []GenericFieldOption {
	if !item.HasOptions() {
		return nil
	}
	options := make([]GenericFieldOption, len(item["options"].([]interface{})))
	for i, optionInterface := range item["options"].([]interface{}) {
		options[i] = GenericFieldOption{
			"label": optionInterface.(map[string]interface{})["label"].(string),
			"value": optionInterface.(map[string]interface{})["value"],
		}
	}
	return options
}

func (item GenericField) GetOptionDisplayValue(value string) string {
	if !item.HasOptions() {
		return ""
	}
	options := item["options"].([]map[string]interface{})
	for _, option := range options {
		if option["value"].(string) == value {
			return option["label"].(string)
		}
	}
	return ""
}
