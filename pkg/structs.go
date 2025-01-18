package pkg

import (
	"fmt"
	"log"
)

type Square struct {
	X int
	Y int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}
	s := Square{x, y,length}
	return &s, nil
}

func (s *Square) Move(dx int, dy int){
	s.X += dx
	s.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func StructExample(){
	square, err := NewSquare(5, 10, 11)
	if err != nil {
		log.Fatalf("Error: can't create square")
	}
	square.Move(4, 5)
	fmt.Printf("%+v\n", square)
	fmt.Println(square.Area())
}