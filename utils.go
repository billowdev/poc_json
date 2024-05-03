package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func ReadJSONB(jsonData []byte, target interface{}) error {
	err := json.Unmarshal(jsonData, target)
	if err != nil {
		return err
	}
	return nil
}

func NewJSONB(data interface{}) (JSONB, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}

	return JSONB(dataMap), nil
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

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Value-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()
	valueString, err := json.Marshal(j)
	return string(valueString), err
}
func (j *JSONB) Scan(value interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Scan-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()

	if data, ok := value.([]byte); ok {
		if err := json.Unmarshal(data, j); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected type for JSONB: %T", value)
	}

	return nil
}
