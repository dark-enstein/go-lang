package code128

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"image/png"
	"log"
	"os"
)

type RequestValue struct {
	Type    string
	Present bool
	Text    string
	Dest    string
}

func Generate(r RequestValue) error {
	// Generate a new writer for Code 128 barcode
	// this format allows you to encode all ASCII characters!
	writer := oned.NewCode128Writer()
	// with the writer, we can start encoding!
	c128Code, err := writer.Encode(r.Text, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		log.Fatalf("impossible to encode barcode: %s", err)
	}
	// create a file that will hold our barcode
	file, err := os.Create(r.Dest + "code128.png")
	if err != nil {
		log.Fatalf("impossible to create file: %s", err)
	}
	defer file.Close()
	// Encode the image in PNG
	err = png.Encode(file, c128Code)
	if err != nil {
		log.Fatalf("impossible to encode barcode in PNG: %s", err)
	} else {
		fmt.Println("Code128 code generated ")
	}
	return nil
}
