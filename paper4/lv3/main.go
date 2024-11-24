package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Author struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  `json:"author"`
	Tags    []string `json:"tags"`
}

func main() {
	author := Author{"lxy", "hello friend"}
	post := Post{"This is my title", "This is my content", author, []string{"科技", "体育"}}
	cj, err := json.Marshal(post)
	if err != nil {
		fmt.Println("序列化出错:", err)
		return
	}
	fmt.Println(string(cj))
	news := strings.Trim(string(cj), "{}")
	fmt.Println(news)
	// 这里只做了去掉最外面的大括号的处理然后直接打印
	// 把每个键值对分离出来并依次打印有点做不来 --03题
	// 反序列化Author字符串
	cjauthor, err := json.Marshal(author)
	if err != nil {
		fmt.Println("author序列化失败，err:", err)
	}
	var p Author
	err = json.Unmarshal(cjauthor, &p)
	if err != nil {
		fmt.Println("反序列化出错:", err)
	}
	fmt.Println(p.Name, "\n", p.Bio)
}
