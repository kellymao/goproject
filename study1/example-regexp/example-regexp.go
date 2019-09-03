package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+\\.[0-9]+"   //模式

	if ok, _ := regexp.MatchString(pattern, s); ok {
		fmt.Println("匹配到了")
	}

	reg, _ := regexp.Compile(pattern)
	all := reg.FindAllString(s, -1)

	fmt.Println(all)
	fmt.Println(len(all))


/*
   匹配到了
   [2578.34 4567.23 5632.18]
   3


 */

	//将匹配到的部分替换为 "##.#"
	str := reg.ReplaceAllString(s, "##.#")
	fmt.Println(str)  // John: ##.# William: ##.# Steve: ##.#
}
