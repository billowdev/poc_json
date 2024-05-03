package main

import (
	"encoding/json"
	"fmt"
	"poc_json/models"
	"reflect"
)

func ReadJSONB(jsonData []byte, target interface{}) error {
	err := json.Unmarshal(jsonData, target)
	if err != nil {
		return err
	}
	return nil
}

func NewJSONB(data interface{}) (models.JSONB, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}

	return models.JSONB(dataMap), nil
}
func GORMStructToMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct")
	}

	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i).Interface()
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		result[jsonTag] = fieldValue
	}

	return result, nil
}
