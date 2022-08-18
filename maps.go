package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      34,
		"moneyPenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["moneyPenny"])
	fmt.Println(m["vaibhav"])
	//check if value exist in map
	v, ok := m["vaibhav"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["james"]; ok {
		fmt.Println("this is present in map", v)
	}

	//add to map
	m["todd"] = 34
	// for range loop
	for k, val := range m {
		fmt.Println(k, val)
	}
	fmt.Println(m)

	//delete from map
	//delete(<map_name>,"key")
	delete(m, "james")
	fmt.Println(m)
	//type of if condition
	if v, ok := m["todd"]; ok {
		fmt.Println(v)
		delete(m, "todd")
		delete(m, "moneyPenny")
	}
	fmt.Println(m)
}
