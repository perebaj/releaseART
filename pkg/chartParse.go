package pkg

import (
	"os"

	"github.com/perebaj/releaseART/interfaces"
	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseYAML() []string {
	yfile, err := os.ReadFile("ChartFiles/Chart.yaml")
	check(err)

	var chart interfaces.Chart
	err2 := yaml.Unmarshal(yfile, &chart)
	check(err2)

	var dependencies_slice []string
	//for each element of Dependencies save repository in a list
	for _, dep := range chart.Dependencies {
		dependencies_slice = append(dependencies_slice, dep.Repository)
	}
	return dependencies_slice

}
