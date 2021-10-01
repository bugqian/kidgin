package Models

func Hello() int64 {
	var a int64
	db.Table("user").Count(&a)
	//RedisClient.Set("bugqian", "最帅", 100*time.Second)
	return a
}
