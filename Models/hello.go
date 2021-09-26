package Models

func Hello() int64 {
	var a int64
	db.Table("user").Count(&a)
	return a
}
