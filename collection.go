package collection

import (
	"encoding/json"
	"reflect"
)

func TypeOf(obj interface{}) string {
	return reflect.TypeOf(obj).String()
}

func Collect(bytes []byte) map[string]interface{} {
	m := map[string]interface{}{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil
	}
	return m
}

func JsonBytes(obj interface{}) []byte {
	res, _ := json.Marshal(obj)
	return res
}

func JsonString(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

func Dict(obj interface{}) map[string]interface{} {
	return obj.(map[string]interface{})
}

func List(obj interface{}) []interface{} {
	return obj.([]interface{})
}

func T(obj interface{}, indices ...interface{}) interface{} {
	for _, index := range indices {
		_, isInt := index.(int)
		if isInt {
			obj = obj.([]interface{})[index.(int)]
		} else {
			obj = obj.(map[string]interface{})[index.(string)]
		}
	}
	return obj
}
