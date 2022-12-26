package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	base_url = "https://pypi.org/project/"
)

func request(p string) (int, error) {
	requestString := base_url + p

	res, err := http.Get(requestString)

	statusCode := res.StatusCode

	return statusCode, err
}

func main() {
	requestedPackage := os.Args[1]
	statusCode, err := request(requestedPackage)

	if err != nil {
		fmt.Print("Error occured while sending the request")
	}

	if statusCode != 200 {
		if statusCode == 404 {
			fmt.Printf("The package requested %s doesn't exist on PyPi\n", requestedPackage)
			os.Exit(1)
		} else {
			fmt.Println("PyPi returns a status code which hasn't been handled by pype yet: ", statusCode)
			os.Exit(1)
		}

	} else {
		fmt.Printf("The package %s is a valid package\n", requestedPackage)
		os.Exit(0)
	}

}
