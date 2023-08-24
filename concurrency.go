package main

import (
	"fmt"
	"sync"
)

//concurrency คือ การทำงานพร้อมกัน หรือ การทำงานก่อนหลังกัน

// jobs
type number struct {
	a int 
	b int
}

// results
type sumn struct {
	r int
}

func workers (jobsCh <-chan number, resultsCh chan<- sumn) {
	for job := range jobsCh {
		resultsCh <- sumn{sumation(job.a, job.b)}
	}
	}

func sumation(a, b int) int {
	return a + b
}

func WorkerPool() {
	  //Worker Pool
	  //ประกาศ struct ที่จะใช้เก็บค่า ชื่อ number
	  nums := []number{
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},

	}

	// สร้าง chan ที่จะใช้เก็บค่า ชื่อ jobsCh และ resultsCh
	// โดย jobsCh จะเก็บค่า number และ resultsCh จะเก็บค่า sumn
	// len(nums) คือ จำนวนของ nums
	jobsCh := make(chan number, len(nums))
	resultsCh := make(chan sumn, len(nums))

	// วนลูปเพื่อส่งค่า number ที่อยู่ใน nums ไปที่ jobsCh
	for _, n := range nums {
		//jobsCh <- n คือ การส่งค่า n ไปที่ jobsCh
		jobsCh <- n

	}
	// ปิด jobsCh
	close(jobsCh)


	//numWorker := 2 คือ จำนวนของ go routines ที่จะทำงานพร้อมกัน
	numWorker := 2
	// w := 0 คือ การกำหนดค่าเริ่มต้นของ w และให้น้อยวกว่า numWorker และเพิ่มค่า w ทีละ 1
	for w := 0; w < numWorker; w++ {
		//go routines
		go workers (jobsCh , resultsCh )
	}

	results := make([]sumn, 0)
	for a := 0; a < len(nums); a++ {
		temp := <-resultsCh
		results = append(results, temp)
	}
	fmt.Println(results)

}



func WaitGroup() {
	//Basic -> sysnc.WaitGroup
	a := []int{1, 2, 3, 4, 5}
// sync.WaitGroup คือ การรอให้ทุกอย่างทำงานเสร็จก่อน จึงจะทำงานต่อไป
	var wg sync.WaitGroup
	for i := range a {
		//wg.Add(1) คือ การบอกว่า งานนี้จะทำงานพร้อมกัน 1 งาน
		wg.Add(1)
		go func (i int) {
			//defer  คือ การทำงานที่จะทำหลังจากที่ function นั้นทำงานเสร็จ
			// wg.Done() คือ การบอกว่า งานนี้เสร็จแล้ว
			defer wg.Done()
			fmt.Printf("%v", a[i])
		}(i)
	}
	wg.Wait()

}

func Chan() {
		//Chan 
		numberCh := make(chan int)
		msgCh := make(chan string)
	
	//number
		go func(numberCh chan<- int) {
			numberCh <- 10
		}(numberCh) 
	
		
	//msg
		go func(magCh chan<- string) {
			msgCh <- "Hello World"
		}(msgCh)
	
		number := <- numberCh
		msg := <- msgCh
	
		fmt.Println(number)
		fmt.Println(msg)
}