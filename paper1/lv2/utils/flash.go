package utils

import (
	"fmt"
)

// lv2.1
func Cau1(op rune, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			fmt.Println("发生除零错误")
			return -1
		} else {
			return a / b
		}
	default:
		fmt.Println("输入字符非四则运算操作符")
		return -1
	}
}
func CauTool1() {
	fmt.Println("欢迎使用go语言计算器")
	fmt.Println("请输入两个整数和一个操作符")
	fmt.Println("输入exit退出程序")
	var (
		a  int
		b  int
		op rune
	)
	for {
		fmt.Println("输入第一个整数：")
		fmt.Scan(&a)
		fmt.Println("输入操作符：")
		fmt.Scanf("%c", &op)
		fmt.Println("输入第二个整数：")
		fmt.Scan(&b)
		fmt.Printf("%d%c%d=%d\n", a, op, b, Cau1(op, a, b))
		var j string
		fmt.Println("是否继续？(exit退出):")
		fmt.Scan(&j)
		if j == "exit" {
			break
		}
	}
	fmt.Println("感谢使用,再见！")
}

// lv2.2 TODO 烂尾了 数据结构的没学到精髓 第一次c转go 改来改去 逻辑还是有点问题 后面熟悉了再尝试写一次 还是很有信心的
//func Cau2(opnd *Floatstack, optr *Charstack, exp []byte) float64 {
//	num := 0.0
//	i := 0
//	var o1, o2, res float64
//	var c byte
//	for {
//		if i >= len(exp) && optr.GetOptrTop() == '=' {
//			break
//		}
//		if IsDigit(exp[i]) {
//			if exp[i+1] == '.' {
//				var c = 1.0
//				num = float64(exp[i] - '0')
//				i = i + 2
//				for i < len(exp) && IsDigit(byte(exp[i])) {
//					c *= 10.0
//					num = num + float64(exp[i]-'0')/c
//					i++
//				}
//				opnd.Push(num)
//			} else {
//				for i < len(exp) && IsDigit(byte(exp[i])) {
//					num = num*10 + float64(exp[i]-'0')
//					i++
//				}
//				opnd.Push(num)
//			}
//			num = 0
//		} else {
//			c = GetPriority(optr.GetOptrTop(), byte(exp[i]))
//			switch c {
//			case '<':
//				optr.Push(byte(exp[i]))
//				i++
//			case '>':
//				o1, _ = opnd.Pop()
//				o2, _ = opnd.Pop()
//				c, _ = optr.Pop()
//				res = Operate(c, o2, o1)
//				opnd.Push(res)
//			case '=':
//				_, _ = optr.Pop()
//				i++
//			}
//		}
//	}
//	r, _ := opnd.Pop()
//	return r
//}
