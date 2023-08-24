package main

import "fmt"

//พิมเล็ก จะเป็น private พิมใหญ่จะเป็น public

type GameOrMovie interface {
	getPrice() int
}

type Game struct {
	Title string
	Platform string
	Price int
}

type Movie struct {
	Title string
	Price int
}

func (g Game) getPrice() int {
	return g.Price
}

func (g Movie) getPrice() int {
	return g.Price
}

func sum[V GameOrMovie](objs []V) int {
	var result int
	for _, obj := range objs {
		result += obj.getPrice()
	}
	return result
}

func M() {
	games := []Game{
		{
			Title:   "GTA V",
			Platform: "PS4",
			Price:    1790,
		},
		{
			Title:   "MineCraft",
			Platform: "PC",
			Price:    490,
		}}
	
		movies := []Movie{
			{
				Title: "Avengers",
				Price: 350,
			},
			{
				Title: "Spiderman",
				Price: 250,
			},
			}
	
	
		fmt.Println(sum(games))
		fmt.Println(sum(movies))
}

