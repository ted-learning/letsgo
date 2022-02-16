package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	updateSlice(s1)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("arr=", arr)

	updateSlice(s2)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("arr=", arr)

	fmt.Println("Re slice")
	s2 = s2[:5]
	s2 = s2[2:]
	fmt.Println(s2)

	extendingSlice()

	sliceOps()
}

func extendingSlice() {
	fmt.Println("Extending Slice")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	s1 := arr[2:6]
	s2 := s1[3:5] // [s1[3],s1[4]]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s1))
	s5[0] = 100
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))
}

func printSlice(s []int) {
	fmt.Printf("value=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
func sliceOps() {
	fmt.Println("Creating Slice")
	var s []int
	printSlice(s)
	for i := 0; i < 100; i++ {
		s = append(s, i*2+1)
		printSlice(s)
	}
	fmt.Println(s)

	s2 := []int{2, 4, 6, 8}
	s3 := make([]int, 16)
	s4 := make([]int, 10, 100)
	printSlice(s2)
	printSlice(s3)
	printSlice(s4)

	fmt.Println("Copying Slice")
	copy(s4, s)
	printSlice(s4)
	printSlice(s4[8:20])

	fmt.Println("Deleting Slice")
	s4 = append(s4[:2], s4[3:]...)
	printSlice(s4)

	fmt.Println("Popping top")
	top := s4[0]
	s4 = s4[1:]
	fmt.Println(top, s4)

	fmt.Println("Popping tail")
	tail := s4[len(s4)-1]
	s4 = s4[:len(s4)-1]
	fmt.Println(tail, s4)
}
