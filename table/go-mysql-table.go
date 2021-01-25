package table

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB

	user         string = "root"
	pass         string = "zm159539"
	dbname       string = "go"
	host         string = "47.96.24.50"
	port         string = "3306"
	table_douban string = "douban_movie"
)

type DoubanMovie struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Title    string `gorm:"column:title"`
	Subtitle string `gorm:"column:subtitle"`
	Other    string `gorm:"column:other"`
	Desc     string `gorm:"column:desc"`
	Year     string `gorm:"column:year"`
	Area     string `gorm:"column:area"`
	Tag      string `gorm:"column:tag"`
	Star     string `gorm:"column:star"`
	Comment  string `gorm:"column:comment"`
	Quote    string `gorm:"column:quote"`
}

func (DoubanMovie) TableName() string {
	return table_douban
}

func CreateTable() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败")
	}
	err = db.AutoMigrate(&DoubanMovie{})
	if err != nil {
		panic("创建表失败")
	} else {
		fmt.Printf("表%v创建成功", table_douban)
	}
}
