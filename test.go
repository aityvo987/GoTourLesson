package main

import (
	"fmt"
	"math"
	"sync"
	"time"
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

type Student struct {
	Name  string
	ID    int
	Year  int
	Major string
}

func (st Student) String() string {
	return fmt.Sprintf("Student %v has ID %v  and belong to %v major, he/she is %v year(s) student", st.Name, st.ID, st.Major, st.Year)
}

//Generic

func Equal[T comparable](s []T, t T) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head *Element[T]
}

type Element[E any] struct {
	next *Element[E]
	val  E
}

func (lst *List[T]) Push(t T) {
	if lst.head == nil {

		lst.head = &Element[T]{val: t}
	} else {
		p := lst.head
		for p.next != nil {
			p = p.next
		}
		p.next = &Element[T]{val: t}
	}
}
func (lst *List[T]) Print() {
	e := lst.head
	for {
		fmt.Print("->", e.val)
		if e.next == nil {
			break
		}
		e = e.next
	}
	fmt.Println()

}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channel
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonanci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonanciS(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

//Mutex

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
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
	dat := Student{"Tien Dat", 20110011, 4, "Cong nghe phan mem"}
	tho := Student{"Tan Tho", 20110211, 4, "Cong nghe phan mem"}
	fmt.Println(dat)
	fmt.Println(tho)
	//Generic
	g := []int{1, 2, 3, 4, 5}
	fmt.Printf("Position of %v in %v is %v\n", 5, g, Equal(g, 5))
	var l List[int]
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Print()
	//Time
	go say("Halo")
	say("Hela")
	//Channel
	str := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)
	go sum(str, c)
	go sum(str[3:], c)
	c1, c2 := <-c, <-c
	fmt.Println(c1, c2)
	//Buffered channel
	//ch := make(chan int, 3)
	ch := make(chan int, 4)
	ch <- 4
	ch <- 3
	ch <- 2
	ch <- 1
	fmt.Println(<-ch)
	//Range and close
	chanrangeclose := make(chan int, 10)
	go fibonanci(cap(chanrangeclose), chanrangeclose)
	for i := range chanrangeclose {
		fmt.Printf("[%v]\n", i)
	}
	//Select
	chanS := make(chan int)
	quitS := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-chanS)
		}
		quitS <- 0
	}()
	fibonanciS(chanS, quitS)

	//Select default
	tick := time.Tick(1 * time.Second)
	boom := time.After(5 * time.Second)
detonate:
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("KABOOM!!!")
			// return
			break detonate
		default:
			fmt.Println(".")
			time.Sleep(250 * time.Millisecond)
		}
	}
	// Mutex
	csafe := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			go csafe.Inc("Even")
		} else {
			go csafe.Inc("Odd")
		}

	}
	time.Sleep(1 * time.Second)
	fmt.Println("Result key:", csafe.Value("Even"), csafe.Value("Odd"))
}
