package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			fmt.Printf("%q\n",s)
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}


func main() {
	data := []string{"one","","three"}
	fmt.Printf("%q\n",nonempty(data))
	fmt.Printf("%q\n",data)
}
