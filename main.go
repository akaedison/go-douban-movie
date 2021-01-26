package main

import (
	"awesomeProject/go-douban-movie/parse"
	_ "github.com/akaedison/go-douban-movie/parse"
)

func main() {
	//table.CreateTable()
	parse.GetDouban("https://movie.douban.com/top250")
}
