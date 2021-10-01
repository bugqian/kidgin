package Models

import "time"

type Article struct {
	Id         int    `gorm:"primaryKey" json:"id" form:"id"`
	Title      string `json:"title" form:"title" gorm:"column:title" binding:"required"`
	Desc       string `json:"desc" form:"desc" gorm:"column:desc" binding:"required"`
	Content    string `json:"content" form:"content" column:"content" binding:"required"`
	CreateTime int64  `json:"create_time" form:"create_time" column:"create_time"`
}

func (t *Article) Add(article Article) error {
	article.CreateTime = time.Now().Unix()
	err := db.Create(&article).Error
	return err
}

func (t *Article) Update(article Article, id int) error {
	err := db.Select("title", "desc", "content").Where("id = ?", id).Updates(article).Error
	return err
}

func (t *Article) Info(id int) (Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	return article, err
}

func (t *Article) Delete(id int) error {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	return err
}
