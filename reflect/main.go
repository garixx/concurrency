package main

import (
	"bytes"
	"fmt"
	"reflect"
)

const (
	quote   = "\""
	twoDots = ": "
)

type User struct {
	Name string
	Age  int64
}

type City struct {
	Name       string
	Population int64
	GDP        int64
	Mayor      string
}

func main() {
	var u User = User{"bob", 10}

	res, err := JSONEncode(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	c := City{"sf", 5000000, 567896, "mr jones"}
	res, err = JSONEncode(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func JSONEncode(v interface{}) ([]byte, error) {
	// TODO: check if v is a struct else return error - done
	e := reflect.ValueOf(v)
	if e.Kind() != reflect.Struct {
		return []byte{}, fmt.Errorf("expected struct but got %T", v)
	}

	// TODO: iterate over v`s reflect value using NumField() - done
	// use type switch to create string result of "{field}" + ": " + "{value}"
	// start with just 2 types - reflect.String and reflect.Int64
	buf := bytes.Buffer{}
	last := e.NumField() - 1

	buf.WriteString("{")
	for i := 0; i < last; i++ {
		WriteField(e, i, &buf)
		buf.WriteString(", ")
	}
	WriteField(e, last, &buf)
	buf.WriteString("}")

	return buf.Bytes(), nil
}

func WriteField(e reflect.Value, i int, buf *bytes.Buffer) {
	field := e.Type().Field(i)
	value := fmt.Sprintf("%v", e.Field(i).Interface())

	buf.WriteString(quote)
	buf.WriteString(field.Name)
	buf.WriteString(quote)
	buf.WriteString(twoDots)

	switch field.Type.Kind() {
	case reflect.String:
		buf.WriteString(quote)
		buf.WriteString(value)
		buf.WriteString(quote)
	//case reflect.Int64:
	default:
		buf.WriteString(value)
	}
}
