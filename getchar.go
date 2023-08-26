package getchar

import (
	"github.com/pkg/term"
)

// Returns either an ascii code, or an arrow ("up", "down", "right", "left").
func getChar() (ascii int, arrow string, err error) {
	t, _ := term.Open("/dev/tty")
	err = term.RawMode(t)
	if err != nil {
		return
	}
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".
		// Since there are no ASCII codes for arrow keys, returning strings
		if bytes[2] == 65 {
			arrow = "up"
		} else if bytes[2] == 66 {
			arrow = "down"
		} else if bytes[2] == 67 {
			arrow = "right"
		} else if bytes[2] == 68 {
			arrow = "left"
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	}
	err = t.Restore()
	if err != nil {
		return
	}
	t.Close()
	return
}
