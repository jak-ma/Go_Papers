package utils

type Student struct {
	Name   string
	Age    int
	Scores []float64
}

type Classroom struct {
	ClassName string
	Students  []*Student
}
