package main

import (
	"github.com/akaedison/go-douban-movie/parse"
)

func main() {
	//table.CreateTable()
	parse.getDouban("https://movie.douban.com/top250")
}
