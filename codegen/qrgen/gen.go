package qrgen

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

type RequestValue struct {
	Type    string
	Present bool
	Text    string
	Dest    string
}

func QrGenerate(r RequestValue) error {
	//gen qrgen code
	qrCode, _ := qr.Encode(r.Text, qr.M, qr.Auto)
	fmt.Println(qrCode)

	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create(r.Dest + "qrcode.png")
	defer file.Close()

	err := png.Encode(file, qrCode)
	if err != nil {
		return err
	} else {
		fmt.Println("QR code generated ")
	}
	return nil
}
