package main

import (
    "encoding/json"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type AnyArray map[string]interface{}


func (a AnyArray) NewGet_Slice(parts []string) interface{} {
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

func (a AnyArray) Get_Slice(parts []string) interface{} {
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


func (a AnyArray) Set_Slice(parts []string, value interface{}) {
    var current AnyArray = a
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
    var current AnyArray = a
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
        copy := reflect.New(typ).Interface()
        if err := json.Unmarshal(data, &copy); err != nil {
            return err
        }
        value = copy
    }

    a.Set(dest, value)
    return nil
}

func (a AnyArray) Copy_Slice(src []string, dest []string) error {
    value := a.Get_Slice(src)
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
        copy := reflect.New(typ).Interface()
        if err := json.Unmarshal(data, &copy); err != nil {
            return err
        }
        value = copy
    }

    a.Set_Slice(dest, value)
    return nil
}

func parsePath(path string) []string {
    parts := strings.Split(path, "[")
    res := []string{}
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

func main() {
    a := make(AnyArray)
    a.Set("xxx[0][1]", "hello")
    a.Set("xxx[0][Test]", "Test")
    a.Set("xxx[TEST][1]", "hellosss")
    a.Set("yyy[0]", 123)

    // 将 xxx 拷贝到 ccc
    if err := a.Copy("xxx[0]", "ccc"); err != nil {
        fmt.Println(err)
        return
    }
    
    a.Copy("yyy[0]", "cccsss");
    
    a.Copy("ccc[1]", "cccs");
    //a.Copy("cccs", "ccc");
    // 验证结果
    fmt.Println(a.Get("xxx[0][1]")) // Output: hello
    fmt.Println(a.Get("xxx[TEST][1]"))
    fmt.Println(a.Get("yyy[0]"))    // Output: 123
    fmt.Println(a.Get("ccc[1]")) // Output: hello
    fmt.Println(a.Get("cccs"))
    a.Set("test[0]", a.NewGet("ccc"))
    fmt.Println(a.Get("ccc"),1)
    Test:=a.NewGet("ccc").(AnyArray)
    Test["1"] = "001"
    fmt.Println(a.Get("ccc"),1)
    fmt.Println(a.Get("test[0][Test]"))
    fmt.Println(a.Get("ccc[Test]"))
    fmt.Println(a.Get("cccsss"))
    
    a.Copy("xxx", "ccc")
    fmt.Println(a.Get("ccc[TEST][1]"),0)
    fmt.Println(a.Get("[ccc][TEST][1]"),1)
    fmt.Println(a.Get("ccc"),0)
    fmt.Println(a.Get("[ccc]"),1)
    fmt.Println(a.Get("xxx"))
    fmt.Println(a.Get("ccc").(AnyArray).NewGet("[TEST]").(AnyArray).NewGet("[1]"))
}