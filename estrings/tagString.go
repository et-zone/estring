package estrings

import (
	"errors"
	"reflect"
)

type TagData map[string]interface{}

// Update the key of map according to tag,
// There is no need to copy data, so the performance impact can be ignored
// src：source Data， desStruct：Target data type
func (t *TagData) ChTagToJsonByTagName(desStruct interface{}, tag string) error {
	if t == nil {
		return errors.New("tagData is nil,the tagData neet init")
	}

	if tag == "" || tag == "json" || len(*t) == 0 {
		return nil
	}

	tVal := reflect.TypeOf(desStruct)

	if tVal.Kind() != reflect.Struct {
		return errors.New("desStruct must struct type ")
	}

	l := tVal.NumField()
	for i := 0; i < l; i++ {
		jsonName := tVal.Field(i).Tag.Get("json")
		tagName := tVal.Field(i).Tag.Get(tag)

		if jsonName == "-" || tagName == "-" {
			continue
		}
		if jsonName == "" {
			jsonName = tVal.Field(i).Name
		}
		if tagName == "" {
			tagName = tVal.Field(i).Name
		}
		if jsonName == tagName {
			continue
		}

		(*t)[jsonName] = (*t)[tagName]
		delete(*t, tagName)
	}
	return nil
}
