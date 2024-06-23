package lib

import "fmt"

type Rec struct {
	Name    string
	Val     int
	RecVal1 *Rec
	RecVal2 *Rec
}

func (rec *Rec) Calc() {
	fmt.Println("#")
	var val1 = 0
	var val2 = 0
	if rec.RecVal1 != nil {
		rec.RecVal1.Calc()
		val1 = rec.RecVal1.Val
	}
	if rec.RecVal2 != nil {
		rec.RecVal2.Calc()
		val2 = rec.RecVal2.Val
	}

	rec.Val = rec.Val + val1 + val2
}
