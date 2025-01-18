package pkg

import "fmt"

type Vehicle interface {
	GetYear() (int, error)
	GetModel() (string, error)
	GetPrice() (int, error)
}

type Car struct {
	Model string
	Year int
	Price int
}

type Motorbike struct {
	Model string
	Year int
	Price int
}

func (c *Car) GetYear() (int, error) {
	if c == nil || c.Year == 0 {
		return -1, fmt.Errorf("Error while getting year")
	}
	return c.Year, nil
}

func (c *Car) GetModel() (string, error) {
	if c == nil || c.Model == "" {
		return "", fmt.Errorf("Error while getting year")
	}
	return c.Model, nil
}

func (c *Car) GetPrice() (int, error) {
	if c == nil || c.Price == 0 {
		return -1, fmt.Errorf("Error while getting year")
	}
	return c.Price, nil
}

func (c *Motorbike) GetYear() (int, error) {
	if c == nil || c.Year == 0 {
		return -1, fmt.Errorf("Error while getting year")
	}
	return c.Year, nil
}

func (c *Motorbike) GetModel() (string, error) {
	if c == nil || c.Model == "" {
		return "", fmt.Errorf("Error while getting year")
	}
	return c.Model, nil
}

func (c *Motorbike) GetPrice() (int, error) {
	if c == nil || c.Price == 0 {
		return -1, fmt.Errorf("Error while getting year")
	}
	return c.Price, nil
}

func BuildVehicle(model string, year int, price int) Vehicle {
	if model == "car" {
		return &Car{Model: model, Year: year, Price: price}
	}
	return &Motorbike{Model: model, Year: year, Price: price}
}

func DisplayVehicleInfo(v Vehicle) {
	fmt.Printf("Printing Vehicle info \n ")
	m, _ := v.GetModel()
	if m != "" && m == "car" {
		fmt.Printf("Vehicle is a car %+v \n", v)
	} else if m != "" && m == "motorbike" {
		fmt.Printf("Vehicle is a motorbike %+v \n", v)
	} else {
		fmt.Println("Vehicle info is nil")
	}
}

// func main(){
// 	c := &car{model: "Audi", year: 2020, price: 10000}
// 	DisplayVehicleInfo(c)
// }

