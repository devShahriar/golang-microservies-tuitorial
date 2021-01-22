package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
)

type User struct {
	name string `validate:"shudip"`
	id   int    `keyValidate:"goru"`
}

func main() {
	u := &User{
		name: "chudip",
		id:   12,
	}

	fi, _ := reflect.TypeOf(u).Elem().FieldByName("name")
	a := strings.Split(string(fi.Tag), ":")
	fmt.Println(a)
	path, _ := filepath.Abs("test.go")

	fmt.Println(filepath.Join("/data/product-db.go", path))
}
