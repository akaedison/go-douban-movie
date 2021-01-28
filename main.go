package main

import (
	"awesomeProject/go-douban-movie/parse"
	_ "github.com/akaedison/go-douban-movie/parse"
)

func main() {
	//table.CreateTable()
	parse.GetDouban("https://movie.douban.com/top250")
	/*fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)*/
}
