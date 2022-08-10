package foo

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func Foo() {
	fmt.Println(uuid.NewV4().String())
}
