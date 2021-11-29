//인터페이스 사용법
//인터페이스에 속해있는 메서드의 서명이 같다면 서로 연관이 없더라도 사용할 수 있다. 즉 동적인 코딩이 가능하고 높은 확장성을 가질 수 있다.
package main

import "fmt"

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

type Rental struct {
	name         string
	feePerDay    float64
	periodlength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodlength)
}

type Items []Coster

func (ts Items) Cost() (c float64) {
	for _, t := range ts {
		c += t.Cost()
	}
	return
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Womens Walking Shoes", 50000, 3},
		10.00,
	}
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	items := Items{shirt, video, eventShoes}
	displayCost(items)
	displayCost(shoes)
	displayCost(eventShoes)
	displayCost(shirt)
	displayCost(video)
}
