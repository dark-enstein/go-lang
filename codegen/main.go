package main

import (
	"flag"
	"fmt"
	"github.com/dark-enstein/go-lang/codegen/aztec"
	"github.com/dark-enstein/go-lang/codegen/code128"
	"github.com/dark-enstein/go-lang/codegen/code39"
	"github.com/dark-enstein/go-lang/codegen/code93"
	"github.com/dark-enstein/go-lang/codegen/datamatrix"
	"github.com/dark-enstein/go-lang/codegen/qrgen"
	"log"
	"os"
	"regexp"
)

type tracingBlob struct {
	allCMDArguments      []string
	flagParsereqInit     requestStatus
	flagParsereqFinal    requestStatus
	flagParseReturnVal   RequestValue
	flagParseResponseVal RequestValue
	codeGenCompileRegexp *regexp.Regexp
	codeGenRegexpMatch   bool
	codeGenDestInit      string
	codeGenDestFinal     string
	codeGenTypeReq       string
}

var Tracer tracingBlob

type requestStatus struct {
	Type *string

	Present *string
	Dest    *string
}

type RequestValue struct {
	Type    string
	Present bool
	Text    string
	Dest    string
}

func flagParse() RequestValue {
	var req requestStatus
	Tracer.flagParsereqInit = req
	req.Type = flag.String("type", "qr", "generate codes. Options: qr, aztec, code39, code93, code128ID, datamatrix")
	req.Present = flag.String("string", "none", "enter input string to be encoded")
	req.Dest = flag.String("dest", ".", "destination to save generated code")

	flag.Parse()
	Tracer.flagParsereqFinal = req

	if *req.Present == "none" {
		Tracer.flagParseReturnVal = RequestValue{*req.Type, false, *req.Present, *req.Dest}
		return RequestValue{*req.Type, false, *req.Present, *req.Dest}
	}

	Tracer.flagParseReturnVal = RequestValue{*req.Type, true, *req.Present, *req.Dest}
	return RequestValue{*req.Type, true, *req.Present, *req.Dest}
}

func codeGen(r RequestValue) error {
	re := regexp.MustCompile(".*/$")
	Tracer.codeGenCompileRegexp = re
	Tracer.codeGenRegexpMatch = re.Match([]byte(r.Dest))
	if !re.Match([]byte(r.Dest)) {
		Tracer.codeGenDestInit = r.Dest
		fmt.Println("Before dest path: ", r.Dest)
		r.Dest = r.Dest + "/"
		fmt.Println("After dest path: ", r.Dest)
	}
	Tracer.codeGenDestFinal = r.Dest
	Tracer.codeGenTypeReq = r.Type

	if r.Type == "qr" {
		err := qrgen.QrGenerate(qrgen.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate QR code")
			return err
		}
	}

	if r.Type == "aztec" {
		err := aztec.Generate(aztec.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate Aztec code")
			return err
		}
	}

	if r.Type == "code39" {
		err := code39.Generate(code39.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate code39 code")
			return err
		}
	}

	if r.Type == "code128ID" {
		err := code128.Generate(code128.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate code128ID code")
			return err
		}
	}

	if r.Type == "code93" {
		err := code93.Generate(code93.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate code93 code")
			return err
		}
	}

	if r.Type == "datamatrix" {
		err := datamatrix.Generate(datamatrix.RequestValue(r))
		if err != nil {
			log.Println("Encountered an error while trying to generate datamatrix code")
			return err
		}
	}

	return nil
}

func main() {
	Tracer.allCMDArguments = os.Args
	request := flagParse()
	Tracer.flagParseResponseVal = request
	if request.Present == false {
		fmt.Println("No text passed")
		os.Exit(1)
	}
	codeGen(request)

	fmt.Println(Tracer)
}
