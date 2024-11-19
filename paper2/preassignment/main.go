package main

import (
	"fmt"
	"time"
)

// 使用你已经学过的知识，尝试用Go编写一个计时器程序
// 当用户：
// 输入0时，重置计时器
// 输入1时，开始计时
// 输入2时，暂停或继续计时
// 00:00:00
// 定义 当前时间
// todo 只能实现一个无限计时的功能
type stime struct {
	second int
	minute int
	hour   int
}

func (cur *stime) Cal(snd int) {
	if snd < 60 {
		cur.hour += 0
		cur.minute += 0
		cur.second = snd
	} else {
		cur.minute += 1
		if cur.minute == 60 {
			cur.hour += 1
			cur.minute = 0
		}
		cur.second = 0
	}
}
func (cur stime) Print() {
	fmt.Printf("\r%02d:%02d:%02d", cur.hour, cur.minute, cur.second)
}
func main() {

	var current stime
	for i := 1; i <= 60; i++ {
		time.Sleep(time.Second)
		if i == 60 {
			i = 0
		}
		current.Cal(i)
		current.Print()
	}
}
