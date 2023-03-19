package main

import "fmt"

func replace(list map[string]string) {

	for key, val := range list {
		new_key := ""
		new_value := ""
		is_change := false

		for _, ch := range key {
			if ch == 'a' {
				is_change = true
				new_key = new_key + "ab"
				continue
			}
			new_key = new_key + string(ch)
		}

		for _, ch := range val {
			if ch == 'a' {
				new_value = new_value + "ab"
				continue
			}
			new_value = new_value + string(ch)
		}

		if is_change {
			delete(list, key)
			list[new_key] = new_value
			continue
		}

		list[key] = new_value

	}

}

func main() {

	x := map[string]string{"a": "A", "b": "ab", "c": "abby"}
	fmt.Println(x)

	replace(x)
	fmt.Println(x)

}
