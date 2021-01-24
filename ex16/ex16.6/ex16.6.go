//ex16/ex16.6/ex16.6.go
package main

import (
	"fmt"

	"github.com/tuckersGo/musthaveGo/ex16/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}
