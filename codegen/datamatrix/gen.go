package datamatrix

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
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
	datamatrixCode, _ := datamatrix.Encode(r.Text)
	//fmt.Println(datamatrixCode)

	datamatrixCode, _ = barcode.Scale(datamatrixCode, 200, 200)

	file, _ := os.Create(r.Dest + "datamatrix.png")
	defer file.Close()

	err := png.Encode(file, datamatrixCode)
	if err != nil {
		return err
	} else {
		fmt.Println("Datamatrix code generated ")
	}
	return nil
}
