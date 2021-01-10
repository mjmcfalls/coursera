package main
import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"encoding/json"

)

func main(){
	// prompts the user to first enter a name, and then enter an address. 
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a name: ")
	text, _ := reader.ReadString('\n')
	name := strings.TrimSuffix(text, "\n")

	fmt.Printf("Enter an adress: ")
	rawaddress, _ := reader.ReadString('\n')
	address := strings.TrimSuffix(rawaddress, "\n")
	
	// create a map and add the name and address to the map using the keys “name” and “address”
	m := map[string]string{
		"Name": name,
		"Address": address,
	}
	 
	// use Marshal() to create a JSON object from the map
	outjson, _ := json.Marshal(m)
	// print the JSON object.
	fmt.Println(string(outjson))
}
