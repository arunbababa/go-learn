package main

import (
	"fmt"
)

var i int = 5

func iniFunc() {
	var ini string = "arunba"
	fmt.Println(ini)
	// 
}

func main() {
	fmt.Println(i)

	// すごい、グローバルというかまぁスコープ外の変数は利用できない！たとえ上位のスコープデモ！！
	// これエラー出る
	// fmt.Println(ini)

	// すごい、変数は必ず使わないといけない！！
	var useItPlease int = 100
	fmt.Println(useItPlease)
}