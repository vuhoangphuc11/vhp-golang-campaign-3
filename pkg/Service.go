package pkg

import (
	"fmt"
	"time"
)

func SayLove(n int) {

	fmt.Println("==========START SAY LOVE==========")
	defer fmt.Println("==========END SAY LOVE==========")
	for i := 1; i < n+1; i++ {
		if i == 4 {
			continue
		}
		fmt.Printf("I love u %d\n", i)

	}
}

func Introduce() {
	fmt.Println("==========START INTRODUCE==========")
	defer fmt.Println("==========END INTRODUCE==========")
	const fullName = "VU HOANG PHUC"
	const tech = "GOLANG DEVELOPER"
	const info = "09/11/2001"
	fmt.Println(fullName)
	fmt.Println(tech)
	fmt.Println(info)

}

func CheckMarks(marks float32) string {
	if marks >= 5 {
		return "Complete"
	}
	if marks >= 8 {
		return "Good"
	}
	if marks >= 9.5 {
		return "Excellent"
	} else {
		return "Fail"
	}
}

func CalculateBMI(weight, height float32) {
	bmi := weight / (height * 2)
	if bmi < 18.5 {
		fmt.Printf("BMI:", bmi, "- you are skinny")
	}
	if bmi >= 18.5 && bmi <= 22.9 {
		fmt.Println("BMI:", bmi, "- you are normal")
	}
	if bmi >= 23 && bmi <= 24.9 {
		fmt.Println("BMI:", bmi, "- Are you at risk of obesity")
	}
	if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("BMI:", bmi, "- you are obese type 1")
	}
	if bmi >= 30 {
		fmt.Println("BMI:", bmi, "you are obese type 2")
	}
}

func SwapValue(a *int, b int) int {
	fmt.Println("After swap value  a =", *a)
	*a = b
	fmt.Print("Before swap value  a = ")
	return *a
}

func GetCurrentTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
