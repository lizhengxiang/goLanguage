package main
import "fmt"

func main(){
	r:=[...]int{99:-4}
	fmt.Println(r[99])
	fmt.Println(r[9])
	fmt.Println(len(r))
}
