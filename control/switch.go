package main

import "fmt"

func main() {
	//switchBasic()
	//switchFallthrough()
	//switchOmitExpression(90)
	switchType("a")
	switchType(1)
	switchType(nil)
	switchType(true)
}

func switchBasic() {
	score := 90
	grade := ""
	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60, 50:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(score, grade)
}

func switchFallthrough() {
	score := 90
	grade := ""
	switch score {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "B"
	case 70, 60, 50:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(score, grade)
}

func switchOmitExpression(score int) {
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 50:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(score, grade)
}

func switchType(x interface{}) {
	switch i := x.(type) {
	case nil:
		fmt.Printf("%v type is %T\n", x, i)
	case int:
		fmt.Printf("%v type is int\n", x)
	case float64:
		fmt.Printf("%v type is float64\n", x)
	case string:
		fmt.Printf("%v type is string\n", x)
	default:
		fmt.Printf("%v is unknown type\n", x)

	}
}
