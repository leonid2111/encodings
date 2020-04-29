package encodings_test

import (
	"fmt"
	"github.com/leonid2111/encodings"
	"testing"
)

var su = "Широкая электрификация южных губерний даст мощный толчок подъёму сельского хозяйства."

func TestEncodings(t *testing.T) {
	fmt.Printf("Original UTF-8: len=%d string: %s\n", len(su), su)
	sw := encodings.EncodeWin1251(su)
	fmt.Printf("Encoded W1251:  len=%d string: %s\n", len(sw), sw) 
	su2 := encodings.DecodeWin1251(sw)
	fmt.Printf("Decoded back:  len=%d string: %s\n", len(su2), su2)
}
