package main

import (
	"flag"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
	"regexp"
)

type tracingBlob struct {
	allCMDArguments      []string
	flagParsereqInit     requestStatus
	flagParsereqFinal    requestStatus
	flagParseReturnVal   requestValue
	flagParseResponseVal requestValue
	codeGenCompileRegexp *regexp.Regexp
	codeGenRegexpMatch   bool
	codeGenDestInit      string
	codeGenDestFinal     string
	codeGenTypeReq       string
}

var tracer tracingBlob

type requestStatus struct {
	Type    *string
	Present *string
	Dest    *string
}

type requestValue struct {
	Type    string
	Present bool
	Text    string
	Dest    string
}

func flagParse() requestValue {
	var req requestStatus
	tracer.flagParsereqInit = req
	req.Type = flag.String("type", "qr", "generate qr codes")
	req.Present = flag.String("string", "none", "enter input string to be encoded")
	req.Dest = flag.String("dest", ".", "destination to save generated code")

	flag.Parse()
	tracer.flagParsereqFinal = req

	if *req.Present == "none" {
		tracer.flagParseReturnVal = requestValue{*req.Type, false, *req.Present, *req.Dest}
		return requestValue{*req.Type, false, *req.Present, *req.Dest}
	}

	tracer.flagParseReturnVal = requestValue{*req.Type, true, *req.Present, *req.Dest}
	return requestValue{*req.Type, true, *req.Present, *req.Dest}
}

func codeGen(r requestValue) error {
	re := regexp.MustCompile(".*/$")
	tracer.codeGenCompileRegexp = re
	tracer.codeGenRegexpMatch = re.Match([]byte(r.Dest))
	if !re.Match([]byte(r.Dest)) {
		tracer.codeGenDestInit = r.Dest
		fmt.Println("Before dest path: ", r.Dest)
		r.Dest = r.Dest + "/"
		fmt.Println("After dest path: ", r.Dest)
	}
	tracer.codeGenDestFinal = r.Dest
	tracer.codeGenTypeReq = r.Type
	if r.Type == "qr" {
		//gen qr code
		qrCode, _ := qr.Encode(r.Text, qr.M, qr.Auto)

		qrCode, _ = barcode.Scale(qrCode, 200, 200)

		file, _ := os.Create(r.Dest + "qrcode.png")
		defer file.Close()

		err := png.Encode(file, qrCode)
		if err != nil {
			return err
		} else {
			fmt.Println("File generated ")
		}
	}
	return nil
}

func main() {
	tracer.allCMDArguments = os.Args
	request := flagParse()
	tracer.flagParseResponseVal = request
	if request.Present == false {
		fmt.Println("No text passed")
		os.Exit(1)
	}
	codeGen(request)

	fmt.Println(tracer)
}
