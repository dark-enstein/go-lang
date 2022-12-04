package code39

import (
	"fmt"
	"github.com/boombuler/barcode/code39"
	"image/png"
	"os"
)

type RequestValue struct {
	Type    string
	Present bool
	Text    string
	Dest    string
}

func Generate(r RequestValue) error {
	code39Code, _ := code39.Encode(r.Text, true, true)
	//fmt.Println(code39Code)

	file, _ := os.Create(r.Dest + "code39.png")
	defer file.Close()

	err := png.Encode(file, code39Code)
	if err != nil {
		return err
	} else {
		fmt.Println("Code39 code generated ")
	}
	return nil
}
