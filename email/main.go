package main

import (
	"fmt"
	"os"
	"regexp"
)

func parseArgs() []string {
	emails := os.Args[1:]
	return emails
}

func emailValidate(s []string) ([]string, []bool) {
	var result []string
	var boolong []bool
	re := regexp.MustCompile(`[a-zA-Z0-9_.]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,4}`)
	for _, val := range s {
		isValid := re.Match([]byte(val))
		if isValid != true {
			passIn := fmt.Sprintf("Email %s is not a valid email", val)
			result = append(result, passIn)
			boolong = append(boolong, isValid)
		} else {
			passIn := fmt.Sprintf("Email %s is a valid email", val)
			result = append(result, passIn)
			boolong = append(boolong, isValid)
		}
	}
	return result, boolong

}

func main() {
	emails := parseArgs()
	results, _ := emailValidate(emails)
	for _, val := range results {
		fmt.Println(val)
	}
}
