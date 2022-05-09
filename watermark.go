package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unsafe"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/font"
	logs "github.com/pdfcpu/pdfcpu/pkg/log"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

const (
	// space
	rowFontSpace = 3.4507383110687027
	colSpace     = 41.5792
)

var Ret unsafe.Pointer = nil

func calcText(text string, w, h float64) string {
	text = text + "      "
	newL := math.Sqrt(w*w + h*h)
	newRow := int(newL / colSpace)
	newCol := int(newL / rowFontSpace / float64(len(text)))

	outStr := make([]string, 0)
	for i := 1; i <= newRow; i++ {
		innerStr := make([]string, 0)
		for j := 1; j <= newCol; j++ {
			innerStr = append(innerStr, text)
		}
		outStr = append(outStr, strings.Join(innerStr, " "))
	}
	return strings.Join(outStr, "\n\n\n")
}

//export AddWaterMark
func AddWaterMark(in *C.char, text *C.char) *C.char {
	goIn := string(C.GoString(in))
	goText := string(C.GoString(text))
	f, err := os.Open(goIn)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		defer func() { recover() }()
		if err = f.Close(); err != nil {
			return
		}
	}()

	config := pdfcpu.NewDefaultConfiguration()
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalln("os.get_current_path error:", err.Error())
	}
	font.UserFontDir = dir
	err = api.InstallFonts([]string{"wryh.ttf"})
	if err != nil {
		log.Fatalln("api.InstallFonts error: ", err.Error())
	}

	if err = font.LoadUserFonts(); err != nil {
		log.Fatalln("load fonts error:", err.Error())
	}

	err = api.InstallFonts([]string{"wryh.ttf"})
	if err != nil {
		log.Fatalln("api.InstallFonts error:", err.Error())
	}

	pagesSize, err := api.PageDims(f, config)
	if err != nil {
		log.Fatalln("api.PageDims  error", err.Error())
	}

	var pdfSize0 pdfcpu.Dim
	if len(pagesSize) > 0 {
		pdfSize0 = pagesSize[0]
	}

	wm, err := api.TextWatermark("",
		"sc: 1.4 abs, op:.2, rot:30, pos:br, fillc: 0.5 0.5 0.5",
		true, false, 4,
	)

	if err != nil {
		log.Fatalln("api.TextWatermark error:", err.Error())
	}

	wm.FontName = "MicrosoftYaHei"
	wm.FontSize = 8
	wm.TextString = calcText(goText, pdfSize0.Width, pdfSize0.Height)
	buf := new(bytes.Buffer)
	err = api.AddWatermarks(f, buf, nil, wm, config)
	defer func() { recover() }()
	if err != nil {
		fmt.Println(err)
	}
	result := C.CString(base64.StdEncoding.EncodeToString(buf.Bytes()))
	Ret = unsafe.Pointer(result) // 定义一个全局变量，来保存 堆内存的地址,释放内存的时候会用到
	return result
}

//export ReleaseMemory
func ReleaseMemory() {
	if Ret != nil {
		C.free(Ret)
		Ret = nil
	}
}

func testWaterMark(in string, text string) []byte {
	logs.SetDefaultOptimizeLogger()
	goIn := in
	goText := text
	f, err := os.Open(goIn)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		defer func() {
			recover()
		}()
		if err = f.Close(); err != nil {
			return
		}
	}()

	//pdfcpu.ConfigPath = "."
	config := pdfcpu.NewDefaultConfiguration()

	font.UserFontDir = "."
	err = api.InstallFonts([]string{"wryh.ttf"})
	if err != nil {
		log.Fatalln("api.InstallFonts error: ", err.Error())
	}

	if err = font.LoadUserFonts(); err != nil {
		log.Fatalln("load fonts error:", err.Error())
	}

	pagesSize, err := api.PageDims(f, config)
	var pdfSize0 pdfcpu.Dim
	if len(pagesSize) > 0 {
		pdfSize0 = pagesSize[0]
	}
	//one of tl,tc,tr,l,c,r,bl,bc,br.
	wm, err := api.TextWatermark(
		"",
		fmt.Sprintf("sc: 1.4 abs, op:.2, rot:30, pos:br, offset:0 0, fillc: 0.5 0.5 0.5"),
		true,
		false,
		4,
	)
	if err != nil {
		log.Fatalln("api.TextWatermark: ", err.Error())
	}

	wm.FontName = "MicrosoftYaHei"
	wm.FontSize = 10
	wm.TextString = calcText(goText, pdfSize0.Width, pdfSize0.Height)
	buf := new(bytes.Buffer)
	err = api.AddWatermarks(f, buf, nil, wm, config)
	defer func() {
		if er := recover(); er != any(nil) {
			log.Println(er)
		}
	}()

	if err != nil {
		fmt.Println(err.Error())
	}
	return buf.Bytes()
}

func main() {

}
