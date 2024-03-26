package main

import (
	"fmt"
	"math"
)

type Vecto struct {
	X, Y float64
}

type MyFloat float64

func swap(x, y int) (int, int) {
	return y, x
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return v - lim
	}
	// return v-lim ----- Khong xai v o day duoc
}
func nearestSqrt(x float64) float64 {
	z := 2.0
	for ((z*z - x) / (2 * z)) != 0 {
		if math.Abs((z*z-x)/(2*z)) < 0.000001 {
			return z
		}
		z -= (z*z - x) / (2 * z)
	}
	return z
}
func transfer(x int) string {
	switch x {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	default:
		return "Number greater than three"
	}
	// Co the su dung nhu sau:
	// switch {
	// case x==1:
	// 	return "something"
	// }
}

// New chapter

func (v Vecto) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (m MyFloat) Abs() float64 {
	return math.Sqrt(float64(m))
}
func Abs2(f MyFloat) float64 {
	return math.Sqrt(float64(f))
}
func (v *Vecto) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// De thay doi gia tri cua vecto truyen vao thi t deu can khai bao pointer o khai bao function
// Nhung khi goi thi` method khong can pointer, con function thi can`
type inter interface {
	Abs() float64
}

func prin(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is integer\n", v)
	case float64:
		fmt.Printf("%v is float64\n", v)
	case string:
		fmt.Printf("%v is string\n", v)
	default:
		fmt.Printf("%v is something else\n", v)
	}
}

type Operator struct {
	Name     string
	Codename string
	Class    string
	NBanner  int
}

func (op Operator) String() string {
	return fmt.Sprintf("Operator %v has codename %v  and belong to %v section, he/she has %v year(s) of experience", op.Name, op.Codename, op.Class, op.NBanner)
}

func main() {
	defer fmt.Println("It's end here.")
	var a, b int = 2, 1
	fmt.Printf("Your number before swapping: %v and %v\n", a, b)
	a, b = swap(3, 4)
	fmt.Printf("Your number after swapping: %v and %v\n", a, b)
	fmt.Printf("Your number:%d, %d, %v\n", a, uint(a), float64(a))
	fmt.Printf("Your Pow:%v\n", pow(2, 5, 10))
	fmt.Printf("Your Nearest Sqrt:%v compare to %v\n", nearestSqrt(5), math.Sqrt(5))
	fmt.Println("Transfer number", transfer((4)))
	// New chapter
	v := Vecto{3, 4}
	pointer := &v
	pointer.X = 2
	fmt.Println("Vecto", v)
	var s [2]string
	s[0] = "Halo"
	s[1] = "Hola"
	day := [5]int{0, 1, 2, 3, 4}
	fmt.Println("Day so", day[1:4])
	pointer.X = 3
	fmt.Println("Vecto length", v.Abs())
	v.Scale(2)
	fmt.Println("Vecto after length", v.Abs())
	// Method duoc khai bao nhan gia tri bth, khong phai pointer thi van truyen pointer vao duoc
	// Function duoc khai bao gia tri bth thi chi duoc nhan gia tri value. Exp:
	// p := &Vertex{4, 3}
	// fmt.Println(p.Abs())
	// fmt.Println(AbsFunc(*p))
	var k inter
	a1 := MyFloat(math.Sqrt(9))
	a2 := Vecto{3, 4}
	k = a1
	k = a2
	fmt.Println("k is", k)
	prin(21)
	prin(21.1)
	prin("Halo")
	prin(21 + 16i)
	hairdresser := Operator{"Susie Glitter", "Goldenglow", "Caster", 4}
	inquisitor := Operator{"Irene", "Irene", "Guard", 3}
	fmt.Println(hairdresser)
	fmt.Println(inquisitor)
}