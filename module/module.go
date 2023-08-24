package module

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id     int		`json:"id"`
	Title  string	`json:"title"`
	Author string	`json:"author"`
}

//ประกาศ interface ของ book
type IBook interface {
	GetBook() *book
	SetBook(title string)
}

//ประกาศ constructor ของ book
//กำหนดให้ทุก function ของ book ให้เชื่อต่อกัย interface ของ book
func NewBook(title, author string) IBook {
	return &book{
		&Book{
			Id:   1,
			Title:  title,
			Author: author,
	}}
}


func main() {
books()


}


//เป็นการสร้างฟังก์ชันที่รับค่าอะไรก็ได้ แต่จะ return ค่าเป็น string
//จัดรูปแบบให้เป็น json แล้ว return ออกมา

func Debug(obj any) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}


func books() {
	//เรียกใช้งาน NewBook โดยส่งค่า title และ author ไป
	b := NewBook("golang", "google")
	//เรียกใช้งานฟังก์ชัน GetBook เปลี่ยนผลแล้วแสดงผลออกมา
	b.SetBook("ABC")
	//เรียกใช้ Debug โดยส่งค่า b.GetBook() ไปเป็นรูปแบบ json
	Debug(b.GetBook())
}

//ประกาศ struct ของ book
type book struct{
	*Book
}

//ฟังชั่น getbook จะ return ค่าออกมาเป็น book
func (b *book) GetBook() *book {
	return b
}
//ฟังชั่น setbook จะ set ค่าให้กับ title
func (b *book) SetBook(Title string) {
	b.Title = Title
}