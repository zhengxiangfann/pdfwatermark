package main

import "C"
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/font"
	logx "github.com/pdfcpu/pdfcpu/pkg/log"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

const (
	// space
	rowFontSpace = 3.4507383110687027
	colSpace     = 41.5792
)

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
		outStr = append(outStr, strings.Join(innerStr, ""))
	}

	return strings.Join(outStr, "\n\n\n")
}

//export AddWaterMark
func AddWaterMark(in *C.char, text *C.char) *C.char {
	//func AddWaterMark(in string, text string) []byte {
	goIn := string(C.GoString(in))
	goText := string(C.GoString(text))

	f, err := os.Open(goIn)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err != nil {
			log.Fatalln("os.Open error:", err.Error())
		}
		if err = f.Close(); err != nil {

			return
		}
	}()

	config := pdfcpu.NewDefaultConfiguration()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("os.Getwd() error:", err.Error())
	}
	log.Println("os.Getwd()", dir)
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
	var pdfSize0 pdfcpu.Dim
	if len(pagesSize) > 0 {
		pdfSize0 = pagesSize[0]
	}
	//tmpStr, bbWidth, bbHeight := calcText(goText, pdfSize0.Width, pdfSize0.Height)
	//dx := ll.X + wm.bb.Width()/2 + float64(wm.Dx) + sin*(wm.bb.Height()/2+dy) - cos*wm.bb.Width()/2
	//dy = ll.Y + wm.bb.Height()/2 + float64(wm.Dy) - cos*(wm.bb.Height()/2+dy) - sin*wm.bb.Width()/2
	//pointX := pdfSize0.Width - bbWidth + bbWidth/2 + bbHeight/2*sin30 - cos30*bbWidth/2
	//pointY := bbHeight/2 - cos30*bbHeight/2 - sin30*bbWidth/2
	//offsetX := int(pdfSize0.Width / 2 * sin30)
	//offsetY := int(pdfSize0.Width / 2 * cos30)
	//one of tl,tc,tr,l,c,r,bl,bc,br.
	wm, err := api.TextWatermark(calcText(goText, pdfSize0.Width, pdfSize0.Height),
		"sc: 1.4 abs, op:.2, rot:30, pos:br, fillc: 0.5 0.5 0.5",
		true, false, 4,
	)
	if err != nil {
		log.Fatalln("api.TextWatermark error:", err.Error())
	}
	wm.FontName = "MicrosoftYaHei"
	wm.FontSize = 8
	buf := new(bytes.Buffer)
	err = api.AddWatermarks(f, buf, nil, wm, config)
	if err != nil {
		fmt.Println(err)
	}
	ret := C.CString(base64.StdEncoding.EncodeToString(buf.Bytes()))
	//C.free(unsafe.Pointer(ret))
	return ret
}

func testWaterMark(in string, text string) []byte {
	logx.SetDefaultDebugLogger()
	goIn := in
	goText := text
	f, err := os.Open(goIn)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalln(err.Error())
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
	//dx := ll.X + wm.bb.Width()/2 + float64(wm.Dx) + sin*(wm.bb.Height()/2+dy) - cos*wm.bb.Width()/2
	//dy = ll.Y + wm.bb.Height()/2 + float64(wm.Dy) - cos*(wm.bb.Height()/2+dy) - sin*wm.bb.Width()/2
	//pointX := pdfSize0.Width - bbWidth + bbWidth/2 + bbHeight/2*sin30 - cos30*bbWidth/2
	//pointY := bbHeight/2 - cos30*bbHeight/2 - sin30*bbWidth/2
	//offsetX := int(pdfSize0.Width / 2 * sin30)
	//offsetY := int(pdfSize0.Width / 2 * cos30)
	//offsetX, offsetY = 0, 0

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
	wm.FontSize = 8
	wm.TextString = calcText(goText, pdfSize0.Width, pdfSize0.Height)
	buf := new(bytes.Buffer)
	err = api.AddWatermarks(f, buf, nil, wm, config)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}

func main() {

}
