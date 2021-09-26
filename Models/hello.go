package Models

import "fmt"

func Hello() {
	var a int64
	db.Table("user").Count(&a)
	fmt.Println(a)
}
