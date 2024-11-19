package main

import (
	"fmt"
	"time"
)

type Task interface {
	Execute() error
}

// PrintTask实现Task接口
type PrintTask struct {
	Message string
}

func (p PrintTask) Execute() error {
	fmt.Println(p.Message)
	return nil
}

// CalculationTask实现Task接口
type CalculationTask struct {
	A float64
	B float64
}

func (c CalculationTask) Execute() error {
	fmt.Println(c.A + c.B)
	return nil
}

// SleepTask实现Task接口
type SleepTask struct {
	Duration int
}

func (s SleepTask) Execute() error {
	fmt.Println("sleep start")
	time.Sleep(time.Duration(s.Duration) * time.Second)
	fmt.Println("sleep end")
	return nil
}

// 创建Scheduler结构体
type Scheduler struct {
	Tasks []Task
}

func (sc *Scheduler) AddTask(task Task) {
	sc.Tasks = append(sc.Tasks, task)
}
func (sc Scheduler) RunAll() {
	for _, v := range sc.Tasks {
		v.Execute()
	}
}

func main() {
	// 演示一下过程
	scheduler := &Scheduler{}

	scheduler.AddTask(PrintTask{Message: "non error"})
	scheduler.AddTask(CalculationTask{A: 50, B: 2})
	scheduler.AddTask(SleepTask{Duration: 3})

	scheduler.RunAll()
}
