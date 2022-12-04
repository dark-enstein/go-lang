package aztec

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/aztec"
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
	aztecCode, _ := aztec.Encode([]byte(r.Text), aztec.DEFAULT_EC_PERCENT, aztec.DEFAULT_LAYERS)
	//fmt.Println(aztecCode)

	aztecCode, _ = barcode.Scale(aztecCode, 200, 200)

	file, _ := os.Create(r.Dest + "aztec-code.png")
	defer file.Close()

	err := png.Encode(file, aztecCode)
	if err != nil {
		return err
	} else {
		fmt.Println("Aztec code generated ")
	}
	return nil
}
