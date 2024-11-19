package main

import (
	"fmt"
	"go_papers/paper2/lv1/utils"
)

const max = 100

func main() {
	students := make([]utils.Student, 6, max)
	i := 0
	for ; i <= 5; i++ {
		students[i].Age = i + 10
		students[i].Score = (float64(i) + 90) / 100
		students[i].Sex = 'm'
	}
	students[0].Name = "xiaoming"
	students[1].Name = "xiaowang"
	students[2].Name = "xiaoliii"
	students[3].Name = "xiaohei"
	students[4].Name = "xiaozhang"
	students[5].Name = "xiaolong"
	fmt.Println("这期学生信息如下:")
	for _, v := range students {
		fmt.Printf("name:%s\tage:%d\tsex:%c\tscore:%.2f\n", v.Name, v.Age, v.Sex, v.Score/0.01)
	}
}
