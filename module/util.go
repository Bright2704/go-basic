package module

import (
	"errors"
	"fmt"
)

func Sum(a, b int) int{
	return a + b
} 

func Boo(c bool) bool {
	return c
}

func sum(name string, a ,b int) (int, string) {
	return a + b , "Hello: " + name 
}

func changer(x int) {
	// Local variable จะเปลี่ยนแปลงได้แค่ในฟังก์ชันนั้นๆ
	x = 20
	return
}


func genErr() error {
	return errors.New("error")
}

func loop() {
	// forloop 
	for i := 0; i < 10 ; i ++ {
		fmt.Println(i)
	}

	// whileloop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

/////////////////////////////////////////////////
func ifelse() {
		var score int = 70
	if score >80 { 
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D") 
	} else {
		fmt.Println("F")
	}

	// err จะใช้ได้เพราะในฟังก์ชัน genErr() มีการ return error จะเรียกใช้ต่อไม่ได้
	if err := genErr(); err != nil {
		panic(err.Error())
	}

}
/////////////////////////////////////

func Variable() {
	
	 a := 10
	fmt.Println(a)

	//ประกาศตัวแปร any คือ ตัวแปรที่เก็บค่าได้ทุกประเภท
	var b any 
	//กำหนดค่าให้กับตัวแปร any
	b = 20
	//แปลงค่า b เป็น c และมีสเตตัส OK ใช้ตรวจสอบว่า b มีค่าเป็น string หรือไม่
	c, OK := b.(string)
	//ถ้าไม่ใช่ string จะ panic
	if !OK {
		panic("b is not a string")
	}
	//แสดงค่า c
	fmt.Println(c)
}

///////////////////////////////////
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

	// a := 10
	// b := 5

	// fmt.Println(a, b)

	// swap(&a, &b)
	// fmt.Println(a, b)

}
////////////////////////////////////////////////Slice , Array

func arrayWathcer(data *[]int) {
	// การเปลี่ยนแปลงข้อมูลใน Array จะเปลี่ยนแปลงข้อมูลใน Slice ด้วย
	// กำหนดตัวที่ 0 ให้เป็น 0
		//data[0] = 0

	// ไม่สามารถพิ่มขนาดได้เพราะ data มีขนาด 4 ตัวแล้ว(data []int)
	//data = append(data, 5)

	//กำหนด Slice ให้เป็น Pointer แล้วค่อยเพิ่มข้อมูล

		*data = append(*data, 5)



}

func ARY() {
	//Array
	a := [3]int{1,2,3}
	fmt.Println(a[1])

	//Slice
	//สร้าง b ที่เป็น Slice ของ int ที่มีขนาด 3 ตัว
	b := make([]int, 3)
	//กำหนดค่าให้ b
	b = []int{1,2,3}
	//fmt.Println(b[1])
	//เพิ่มข้อมูลใน Slice
	b = append(b, 4)
	fmt.Println(b)

// เรียก arraywathcer แล้วส่ง b ที่เป็น &b ไป คือส่ง Pointer ไป
//เป็นการ Pass by Reference เพิ่มที่อยู่ของ b ไปให้ arrayWathcer
	arrayWathcer(&b)
//จะกลายเป็น [1 2 3 4 5]
	fmt.Println(b)


	//การวนลูป Slice
	C := []int{1,2,3,4,5}

	//การเรียกใช้แบบปกติ
	// for i := range C {
	// 	fmt.Println(C[i])
	// }

	//การเรียกแบบไม่เอา value
	for _, v := range C {
		fmt.Println(v)
	}
	//จะได้ค่า [1 2 3 4 5]เหมือนกัน 

}
