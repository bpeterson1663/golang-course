package main

import (
	"fmt"
	"sort"

	"github.com/eiannone/keyboard"
)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)
	go listenForKeyPress()
	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	//declares an array
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "bird"

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My new car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
	pointers()
	slices()
	maps()

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.says()
	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPressChan <- char //send infor to a channel
	}

	interfaces()
}

func pointers() {
	x := 10

	myFirstPointer := &x //the location in memory of x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15 //go to the location and set to 15

	fmt.Println("x is now", x) // x is now 15
}

func slices() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is ", animals[0])

	fmt.Println("First two are ", animals[0:2])
	fmt.Println("THe slice", len(animals), "elements long")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals)) //cat, dog, fish, horse

	animals = deleteFromSlice(animals, 2) //should remove fish

	fmt.Println("No more cat", animals)

}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func maps() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}
}

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) says() { //a is called an reciever. Way to add a function to type
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan //receive info into a channel
		fmt.Println("you pressed", string(key))
	}
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

type AnimalInterface interface {
	Says() string
	HowManyLegs() int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}
func interfaces() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	fmt.Println(dog)
}

func riddle(a AnimalInterface) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type CarComposition struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("number of passengers: ", v.NumberOfPassengers)
	fmt.Println("Number of wheels: ", v.NumberOfWheels)
}

func (c CarComposition) show() {
	fmt.Println("Make", c.Make)
	fmt.Println("Model: ", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric", c.IsElectric)
	fmt.Println("Is Hybrid", c.IsHybrid)
	c.Vehicle.showDetails()
}

func composition() {

	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := CarComposition{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()
}
