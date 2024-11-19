package main

import (
	"fmt"
	"go_papers/paper2/lv2/utils"
	"math/rand"
	"time"
)

const max = 100

func main() {

	classroom := &utils.Classroom{
		ClassName: "040923",
	}
	stu_s := make([]*utils.Student, 6, max)
	stu_s[0] = &utils.Student{
		Name: "xiaoming",
		Age:  18,
	}
	stu_s[1] = &utils.Student{
		Name: "xiaowang",
		Age:  19,
	}
	stu_s[2] = &utils.Student{
		Name: "xiaoli",
		Age:  17,
	}
	stu_s[3] = &utils.Student{
		Name: "xiaohei",
		Age:  16,
	}
	stu_s[4] = &utils.Student{
		Name: "xiaogu",
		Age:  14,
	}
	stu_s[5] = &utils.Student{
		Name: "xiaopi",
		Age:  18,
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		scores := make([]float64, 4)
		for j := 0; j < 4; j++ {
			scores[j] = 90 + rand.Float64()*(100-90)
		}
		UpdateScore(stu_s[i], scores)
		AddStudent(classroom, stu_s[i])
	}
	for _, v := range classroom.Students {
		fmt.Printf("name:%s \tage:%d \tscore:{yuwen:%.2f,"+
			"shuxue:%.2f,yingyu:%.2f,wuli:%.2f} \taverage:%.2f\n",
			v.Name, v.Age, v.Scores[0], v.Scores[1], v.Scores[2],
			v.Scores[3], CalculateAverage(v))
	}
}

func AddStudent(c *utils.Classroom, s *utils.Student) {
	c.Students = append(c.Students, s)
}

func UpdateScore(s *utils.Student, score []float64) {
	s.Scores = score
}

func CalculateAverage(s *utils.Student) float64 {
	var average float64
	for _, v := range s.Scores {
		average += v
	}
	w := len(s.Scores)
	average /= float64(w)
	return average
}
