package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "letsgo",
		"course":  "golang",
		"side":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[int]string
	fmt.Println(m3)

	fmt.Println("Get values")
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println(m["side"])
	k, ok := m["abc"]
	fmt.Println(k, ok)
	k2, ok2 := m["course"]
	fmt.Println(k2, ok2)

	fmt.Println("Deleting value")
	delete(m, "course")
	getValue(m, "course")
}

func getValue(m map[string]string, k string) {
	if v, ok := m[k]; ok {
		fmt.Println("value is :", v)
	} else {
		fmt.Println("key does not exist")
	}
}
