package vm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type AnyArray map[string]interface{}

func (a AnyArray) NewGetSlice(parts []string) interface{} {
	var current interface{} = a
	for _, part := range parts {
		switch t := current.(type) {
		case AnyArray:
			current = deepCopy(t[part])
		case []interface{}:
			index, _ := strconv.Atoi(part)
			current = deepCopy(t[index])
		default:
			return nil
		}
	}
	return current
}

func (a AnyArray) GetSlice(parts []string) interface{} {
	var current interface{} = a
	for _, part := range parts {
		switch t := current.(type) {
		case AnyArray:
			current = t[part]
		case []interface{}:
			index, _ := strconv.Atoi(part)
			current = t[index]
		default:
			return nil
		}
	}
	return current
}
func (a AnyArray) NewGet(path string) interface{} {
	parts := parsePath(path)
	var current interface{} = a
	for _, part := range parts {
		switch t := current.(type) {
		case AnyArray:
			current = deepCopy(t[part])
		case []interface{}:
			index, _ := strconv.Atoi(part)
			current = deepCopy(t[index])
		default:
			return nil
		}
	}
	return current
}

func (a AnyArray) Get(path string) interface{} {
	parts := parsePath(path)
	var current interface{} = a
	for _, part := range parts {
		switch t := current.(type) {
		case AnyArray:
			current = t[part]
		case []interface{}:
			index, _ := strconv.Atoi(part)
			current = t[index]
		default:
			return nil
		}
	}
	return current
}

func (a AnyArray) SetSlice(parts []string, value interface{}) {
	var current = a
	for _, part := range parts[:len(parts)-1] {
		if _, ok := current[part].(AnyArray); !ok {
			current[part] = make(AnyArray)
		}
		current = current[part].(AnyArray)
	}
	current[parts[len(parts)-1]] = value
}

func (a AnyArray) Set(path string, value interface{}) {
	parts := parsePath(path)
	var current = a
	for _, part := range parts[:len(parts)-1] {
		if _, ok := current[part].(AnyArray); !ok {
			current[part] = make(AnyArray)
		}
		current = current[part].(AnyArray)
	}
	current[parts[len(parts)-1]] = value
}

func (a AnyArray) Copy(src string, dest string) error {
	value := a.Get(src)
	if value == nil {
		return fmt.Errorf("key '%s' not found", src)
	}

	// 使用反射获取值的类型，如果是结构体或者指针，使用 json 序列化和反序列化实现深拷贝
	typ := reflect.TypeOf(value)
	if typ.Kind() == reflect.Struct || typ.Kind() == reflect.Ptr {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		copys := reflect.New(typ).Interface()
		if err := json.Unmarshal(data, &copys); err != nil {
			return err
		}
		value = copys
	}

	a.Set(dest, value)
	return nil
}

func (a AnyArray) CopySlice(src []string, dest []string) error {
	value := a.GetSlice(src)
	if value == nil {
		return fmt.Errorf("key '%s' not found", src)
	}

	// 使用反射获取值的类型，如果是结构体或者指针，使用 json 序列化和反序列化实现深拷贝
	typ := reflect.TypeOf(value)
	if typ.Kind() == reflect.Struct || typ.Kind() == reflect.Ptr {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		copys := reflect.New(typ).Interface()
		if err := json.Unmarshal(data, &copys); err != nil {
			return err
		}
		value = copys
	}

	a.SetSlice(dest, value)
	return nil
}

func parsePath(path string) []string {
	parts := strings.Split(path, "[")
	var res []string
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		res = append(res, strings.TrimSuffix(part, "]"))
	}
	return res
}

/*
	func parsePath(path string) []string {
	    parts := strings.Split(path, "[")
	    for i, part := range parts {
	        parts[i] = strings.TrimSuffix(part, "]")
	    }
	    return parts
	}
*/
func deepCopy(i interface{}) interface{} {
	switch i := i.(type) {
	case []interface{}:
		cp := make([]interface{}, len(i))
		for i, v := range i {
			cp[i] = deepCopy(v)
		}
		return cp
	case AnyArray:
		cp := make(AnyArray, len(i))
		for i, v := range i {
			cp[i] = deepCopy(v)
		}
		return cp
	case map[string]interface{}:
		cp := make(map[string]interface{})
		for k, v := range i {
			cp[k] = deepCopy(v)
		}
		return cp
	default:
		return i
	}
}

func interfaceToAny(value map[string]interface{}) AnyArray {
	var Res = make(AnyArray)
	for k, v := range value {
		switch v := v.(type) {
		case map[string]interface{}:
			Res[k] = interfaceToAny(v)

		default:
			Res[k] = v
		}

	}
	return Res
}

func JsonTmp(value interface{}) interface{} {
	switch v := value.(type) {
	case AnyArray:
		var Tmp = make([]interface{}, 0)
		loack := true
		for i := 0; i <= len(v)-1; i++ {
			String := strconv.Itoa(i)
			value, Type := v[String]
			if loack && Type {
				switch value := value.(type) {
				case AnyArray:
					Tmp = append(Tmp, JsonTmp(value))
				default:
					Tmp = append(Tmp, value)
				}
			}
			if !Type {
				loack = false
				break
			}
		}
		if loack {
			return Tmp
		} else {
			Res := make(AnyArray)
			for k, v := range v {
				Res[k] = JsonTmp(v)
			}
			return Res
		}
	default:
		return v
	}
}
