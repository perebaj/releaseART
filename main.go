package main

import (
	"fmt"

	"github.com/perebaj/releaseART/pkg"
)

func main() {
	dependecies_slice := pkg.ParseYAML()
	fmt.Println(dependecies_slice)
	pkg.ArtifacthubFetch("argo", "argo-cd")

}
