package entities

import (
	"fmt"
)

type (
	templateFields map[string]bool
)

var templateList = make(map[uint16]templateFields)

func addFieldToTemplate(id uint16, name string) {
	if _, exists := templateList[id]; !exists {
		templateList[id] = make(map[string]bool)
	}
	templateList[id][name] = true
}

func getTemplateFields(id uint16) (templateFields, error) {
	if list, exist := templateList[id]; exist {
		return list, nil
	}
	err := fmt.Errorf("template: template id does not exist.")
	return nil, err
}
