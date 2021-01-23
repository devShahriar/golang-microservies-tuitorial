package handlers

import (
	"fmt"
	"path/filepath"
	"reflect"
)

type User struct {
	name string `key:"shudip"`
	id   int    `keyValidate:"goru"`
}

func main() {
	u := &User{
		name: "chudip",
		id:   12,
	}

	fi, _ := reflect.TypeOf(u).Elem().FieldByName("name") // get the tag value in a struct
	fmt.Println(string(fi.Tag))
	path, _ := filepath.Abs("test.go")

	fmt.Println(filepath.Join("/data/product-db.go", path))
}
