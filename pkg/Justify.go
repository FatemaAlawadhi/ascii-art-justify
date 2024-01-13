package pkg

import (
	"os"
	"strings"
	"golang.org/x/term"
)

func Justify(Align string, Text string, JustifyArray []string) string {
	TermWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	Char := []rune(Text)
	if Align != "justify" {
		if len(Char) > TermWidth {
			return "Error"
		}
	}
	var FinalText string
	var Space int
	if Align == "center" {
		Space = (TermWidth - len(Char)) / 2
		FinalText = strings.Repeat(" ", Space) + Text + strings.Repeat(" ", Space)
	} else if Align == "right" {
		Space = (TermWidth - len(Char))
		FinalText = strings.Repeat(" ", Space) + Text
	} else if Align == "justify" {
		LenJustify := 0
		for i := 0; i < len(JustifyArray); i++ {
			LenJustify = LenJustify + len(JustifyArray[i])
		}
		if LenJustify > TermWidth {
			return "Error"
		}
		Space = (TermWidth - LenJustify)
		WordSpaces := len(JustifyArray) - 1
		AddSpaces := Space / WordSpaces
		NewSpaces := strings.Repeat(" ", AddSpaces)
		for i := 0; i < len(JustifyArray)-1; i++ {
			FinalText = FinalText + JustifyArray[i] + NewSpaces
		}
		FinalText = FinalText + JustifyArray[(len(JustifyArray)-1)]
	} else if Align == "left" {
		FinalText = Text
	}

	return FinalText
}


