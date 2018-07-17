package main
import "fmt"
func main() {
	fmt.Println("Hello, Go!")
	test(1, "hello", 12.5)
}

func test(params...interface{})(){
	test_a, test_b, test_c := params[0], params[1], params[2]
	fmt.Println(test_a)
	fmt.Println(test_b)
	fmt.Println(test_c)
}