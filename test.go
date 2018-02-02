package main
 
import (
	"C"
    "fmt"
    "runtime"
    "math/rand"
    "time"
)

func test(src string, dest int) int {

	if src == "src" {
		return 1
	}else if src == "dest" {
		return dest
	}
	return -1
}

func translate(src, dest int) {
	src = src + dest
	dest = src - dest
	src = src - dest
	fmt.Println(src , dest)
}

func translateP(src, dest *int) {
	*src = *src + *dest
	*dest = *src - *dest
	*src = *src - *dest
	fmt.Println(*src , *dest)
}

func boring(msg string, c chan string) {
   for i := 0; ; i++ {
   		fmt.Println("boring while")
        c <- fmt.Sprintf("%s %d", msg, i)
        time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	var i = 1.3
	var j = 1

	var k float64

	k = 2.01

	//const c int = 11

    fmt.Println("Hellow World!", runtime.Version())
	fmt.Println(i, j)
	fmt.Println(k)
	//fmt.Println(c)

	var sum int

	for i := 0; i < 100; i++ {
		sum += i
	}

	fmt.Println(sum)

	if (true) {
		fmt.Println("true")
	}

	var p string

	p = "B"

	switch p {
		case "A":
			fmt.Println("A")
		case "B":
			fmt.Println("B")
		case "C":
			fmt.Println("C")
		default:
			fmt.Println("default")
	}

	fmt.Println(test("dest", 123))

	var r, t int = 123, 321
	translateP(&r, &t)
	fmt.Println(r , t)


	fmt.Println(C.count)
    fmt.Println(C.funcx())

	c := make(chan string)
	go boring("boring!", c)
 	for i := 0; i < 5; i++ {
    	fmt.Printf("You say: %q\n", <-c)
    }
    fmt.Println("You're boring; I'm leaving.")
}   