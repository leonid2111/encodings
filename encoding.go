package encodings

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

// EncodeWin1251 converts UTF-8 string to Windows1251 string
func EncodeWin1251(su string) string {
	var bw bytes.Buffer
	wInUTF8 := transform.NewWriter(&bw, charmap.Windows1251.NewEncoder())
	wInUTF8.Write([]byte(su))
	wInUTF8.Close()
	return bw.String()
}

// DecodeWin1251 converts Windows1251 encoded string to UTF-8
func DecodeWin1251(sw string) string {
	rInUTF8 := transform.NewReader(strings.NewReader(sw), charmap.Windows1251.NewDecoder())
	bu, _ := ioutil.ReadAll(rInUTF8)
	return string(bu)
}
