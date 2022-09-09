package main

import (
	"fmt"
	m "github.com/vuhoangphuc11/vhp-golang-campaign-3/model"
	p "github.com/vuhoangphuc11/vhp-golang-campaign-3/pkg"
)

func main() {
	a := 99
	p.Introduce()
	p.SayLove(5)
	p.CalculateBMI(74, 1.6)
	p.GetCurrentTime()
	fmt.Println("Your academic ability is:", p.CheckMarks(7.9))
	user := m.User{FullName: "Vu Hoang Phuc", Age: 21, Email: "vuhoangphuc11@gmail.com", Phone: "0969141969"}
	fmt.Println(user)
	fmt.Println(p.SwapValue(&a, 1))

	user = m.User{FullName: "Vu Hoang Phuc", Phone: "03667788900", Email: "vuhoangphuc11@gmail.com", Age: 21}

	user.GetInfo()
	user.ChangeInfo()
	user.InputInfo(user)

}
