package pkg

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseYAML() {
	data, err := os.ReadFile("ChartFiles/Chart.yaml")
	check(err)
	fmt.Println(string(data))
	var dataJson map[string]interface{}
	yaml.Unmarshal(data, &dataJson)
	map2, _ := json.Marshal(dataJson["dependencies"])
	fmt.Println((map2[0]))
	fmt.Println((string(map2)))

	// dependencies := dataJson["dependencies"] //map[string]interface{}
	// teste := json.Marshal(dependencies)

	// fmt.Println(dataJson["dependencies"])
	// for key, element := range dependencies["version"] {
	// 	fmt.Println("key:", key, "element:", element)
	// }

}
