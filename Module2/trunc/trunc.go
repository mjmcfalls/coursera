package main
import (
	"fmt"
	"strings"
)

func main(){
	var input string
	fmt.Printf("Enter a floating point number (ex: 12.3):")
	fmt.Scan(&input)
	fmt.Printf("Truncated value: %s", strings.Split(input,".")[0])
}
