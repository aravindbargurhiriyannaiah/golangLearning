package main

import (
	"fmt"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("http://google.com")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(-1)
	// }
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)

	t := triangle{
		base:   8,
		height: 10,
	}
	s := square{
		sideLength: 5,
	}
	calculateAndPrintArea(t)
	calculateAndPrintArea(s)
}

// func (logWriter) Write(byteSlice []byte) (int, error) {
// 	fmt.Println(string(byteSlice))
// 	fmt.Println("Wrote these many bytes ", len(byteSlice))
// 	return len(byteSlice), nil
// }

type shape interface {
	printArea()
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func calculateAndPrintArea(s shape) {
	s.printArea()
}

func (s square) printArea() {
	fmt.Printf("Square: sideLength = %v, area = %v\n", s.sideLength, s.sideLength*s.sideLength)
}

func (t triangle) printArea() {
	a := 0.5 * t.base * t.height
	fmt.Printf("Triangle: base = %v, height = %v area = %v\n", t.base, t.height, a)
}
