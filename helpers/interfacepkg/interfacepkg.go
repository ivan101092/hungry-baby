package interfacepkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// InArray ...
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// Exist ...
func Exist(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

// InterfaceArrayToString ...
func InterfaceArrayToString(data []interface{}) (res string) {
	for i, d := range data {
		if i != 0 {
			res = res + ", "
		}
		res = res + fmt.Sprintf("%v", d)
	}

	return res
}

// InterfaceStringToString ...
func InterfaceStringToString(data interface{}, key string) string {
	if data == nil || key == "" {
		return ""
	}

	res := fmt.Sprintf("%v", data.(map[string]interface{})[key])
	if res == "<nil>" {
		res = ""
	}

	return res
}

// Marshal convert interface json to string
func Marshal(data interface{}) (res string) {
	name, err := json.Marshal(data)
	if err != nil {
		return res
	}
	res = string(name)

	return res
}

// Unmarshall convert string to interface json
func Unmarshall(data string) (res interface{}) {
	json.Unmarshal([]byte(data), &res)

	return res
}

// UnmarshalCb convert string to interface json
func UnmarshalCb(data string, res interface{}) {
	json.Unmarshal([]byte(data), &res)
}

// UnmarshalCbInterface convert interface to interface json
func UnmarshalCbInterface(data interface{}, res interface{}) {
	dataString := Marshal(data)
	json.Unmarshal([]byte(dataString), &res)
}

// MarshalMap convert map string interface json to string
func MarshalMap(data map[string]interface{}) (res string) {
	name, err := json.Marshal(data)
	if err != nil {
		return res
	}
	res = string(name)

	return res
}

// UnmarshallMap convert string map string interface json
func UnmarshallMap(data string) (res map[string]interface{}) {
	json.Unmarshal([]byte(data), &res)

	return res
}

// Convert ...
func Convert(data interface{}, cb interface{}) (err error) {
	dataString := Marshal(data)
	err = json.Unmarshal([]byte(dataString), &cb)

	return err
}
