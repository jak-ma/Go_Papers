package utils

import "fmt"

const max = 100

type Floatstack struct {
	base []float64
	top  int
}

type Charstack struct {
	base []byte
	top  int
}

func InitOpnd(s *Floatstack) {
	s.base = make([]float64, 0, max)
	if s.base == nil {
		fmt.Println("内存分配失败...")
		return
	}
	s.top = 0
}

func InitOptr(s *Charstack) {
	s.base = make([]byte, 0, max)
	if s.base == nil {
		fmt.Println("内存分配失败...")
		return
	}
	s.top = 0
}

func (s *Floatstack) Push(v float64) {
	if s.top == max-1 {
		fmt.Println("栈满...")
		return
	}
	s.base = append(s.base, v)
	s.top++
}

func (s *Charstack) Push(v byte) {
	if s.top == max-1 {
		fmt.Println("栈满...")
		return
	}
	s.base = append(s.base, v)
	s.top++
}

func (s *Floatstack) Pop() (float64, error) {
	if s.top == 0 {
		return 0, fmt.Errorf("empty FloatStack")
	}
	s.top--
	value := s.base[s.top]
	return value, nil
}

func (s *Charstack) Pop() (byte, error) {
	if s.top == 0 {
		return 0, fmt.Errorf("empty CharStack")
	}
	s.top--
	value := s.base[s.top]
	s.base = s.base[:s.top]
	return value, nil
}

func (s Charstack) GetOptrTop() byte {
	if s.top == 0 {
		return ')'
	}
	value := s.base[s.top-1]
	return value
}

func IsDigit(v byte) bool {
	return v >= '0' && v <= '9'
}

func Operate(op byte, a, b float64) float64 {
	var c float64
	switch op {
	case '+':
		c = a + b
	case '-':
		c = a - b
	case '*':
		c = a * b
	case '/':
		if b == 0 {
			fmt.Println("发生除0错误")
			return 0
		} else {
			c = a / b
		}
	}
	return c
}

// 借用了下数据结构实验课的函数
func GetIndex(t byte) int {
	index := 0
	switch t {
	case '+':
		index = 0
	case '-':
		index = 1
	case '*':
		index = 2
	case '/':
		index = 3
	case '(':
		index = 4
	case ')':
		index = 5
	case '=':
		index = 6
	default:
		break
	}
	return index
}

// 借用了下数据结构实验课的符号表
func GetPriority(t1, t2 byte) byte {
	priority := [7][7]byte{
		{'>', '>', '<', '<', '<', '>', '>'},
		{'>', '>', '<', '<', '<', '>', '>'},
		{'>', '>', '>', '>', '<', '>', '>'},
		{'>', '>', '>', '>', '<', '>', '>'},
		{'<', '<', '<', '<', '<', '=', '0'},
		{'>', '>', '>', '>', '0', '>', '>'},
		{'<', '<', '<', '<', '<', '0', '='},
	}
	index1 := GetIndex(t1)
	index2 := GetIndex(t2)
	return priority[index1][index2]
}
