package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ArtifacthubFetch(RepositoryName string, PackageName string) {
	url := fmt.Sprintf("https://artifacthub.io/api/v1/packages/helm/%s/%s", RepositoryName, PackageName)
	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	// var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// var bodyString = bodyBytes
		var bodyJson map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &bodyJson); err != nil {
			log.Fatal(err)
		}
		fmt.Println(bodyJson["version"])
		// out, _ := json.Marshal(bodyString)
		// fmt.Println(bodyString)
	}
}
