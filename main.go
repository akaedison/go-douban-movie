package main

import (
	"awesomeProject/go-douban-movie/parse"
)

func main() {
	//table.CreateTable()
	parse.GetPages("https://movie.douban.com/top250")
}
