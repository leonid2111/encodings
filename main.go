package main
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var su = "Широкая электрификация южных губерний даст мощный толчок подъёму сельского хозяйства."

func main() {
	fmt.Println(su)
	var bw bytes.Buffer
	wInUTF8 := transform.NewWriter(&bw, charmap.Windows1251.NewEncoder())
	wInUTF8.Write([]byte(su))
	wInUTF8.Close()

	fmt.Printf("%#v\n", bw)
	sw := bw.String()
	fmt.Printf("%s\n", sw) 


	// --- Decoding: convert encS from Windows1251 to UTF8
	rInUTF8 := transform.NewReader(strings.NewReader(sw), charmap.Windows1251.NewDecoder())
	bu, _ := ioutil.ReadAll(rInUTF8)
	su2 := string(bu)
	fmt.Println(su2)
	
}
