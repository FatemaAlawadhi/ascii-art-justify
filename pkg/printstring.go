package pkg

import (
	"fmt"
	"strings"
)

func PrintString(Banner []string, TextArray []string, color string, ColoredLetters string, Align string) []string {
	var word []rune
	var LineBanner int
	var Line string
	var Text []string
	NewLine := 0
	for str := 0; str < len(TextArray); str++ {
		word = []rune(TextArray[str])
		if len(word) > 0 {
			for Linenum := 1; Linenum <= 8; Linenum++ {
				var JustifyArray []string
				for chr := 0; chr < len(word); chr++ {
					LineBanner = ((int(word[chr]) - 32) * 9) + (1 * Linenum)
					if LineBanner < 0 || LineBanner > 856 {
						fmt.Println("Please write in English")
						return Text
					}
					if color != "" && color != "No Color" && ColoredLetters != "empty" && strings.TrimSpace(ColoredLetters) != "" {
						Banner[LineBanner] = Color(ColoredLetters, color, Banner[LineBanner], word[chr])
					}
					Line = Line + Banner[LineBanner]
					if word[chr] == ' ' && Align == "justify" && chr != 0 {
						JustifyArray = append(JustifyArray, Line)
						Line = ""
					} else if chr == (len(word)-1) && Align == "justify" {
						JustifyArray = append(JustifyArray, Line)
					}
				}

				if color != "No Color" && ColoredLetters == "empty" {
					Line = Color("", color, Line, ' ')
				}
				if Align != "No Alignment" {
					if len(JustifyArray) == 1 && Align == "justify" {
						Align = "left"
					}
					Line = Justify(Align, Line, JustifyArray)
					if Line == "Error" {
						fmt.Println("Text width is more than Terminal width")
						return Text
					}
				}
				Text = append(Text, Line)
				Text = append(Text, "\n")
				Line = ""
			}
		} else if len(word) == 0 {
			NewLine++
			if str == (len(TextArray)-1) && NewLine == len(TextArray) {
				return Text
			}
			Text = append(Text, "\n")
		}
	}
	return Text
}
