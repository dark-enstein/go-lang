package code93

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code93"
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
	code93, _ := code93.Encode(r.Text, true, true)
	//fmt.Println(code93)

	code93, _ = barcode.Scale(code93, 200, 200)

	file, _ := os.Create(r.Dest + "code93-code.png")
	defer file.Close()

	err := png.Encode(file, code93)
	if err != nil {
		return err
	} else {
		fmt.Println("code93 code generated ")
	}
	return nil
}
