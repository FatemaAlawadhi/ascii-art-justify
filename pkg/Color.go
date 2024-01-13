package pkg

func Color(ColoredLetters string, color string, LetterLine string, char rune) string {
	FColor := []rune(color)
	for i := 0; i < len(FColor); i++ {
		if FColor[i] >= 'A' && FColor[i] <= 'Z' {
			FColor[i] = FColor[i] + 32
		}
	}
	color = string(FColor)
	if color == "black" {
		color = "\u001b[30m" 
	} else if color == "red" {
		color = "\u001b[31m"
	} else if color =="green" {
		color = "\u001b[32m"
	} else if color == "yellow" {
		color = "\u001b[33m"
	} else if color == "blue" {
		color = "\u001b[34m"
	} else if color == "magenta" {
		color = "\u001b[35m"
	} else if color == "cyan" {
		color = "\u001b[36m"
	}else if color == "white" {
		color =	"\u001b[37m"
	}
	Reset := "\u001b[0m"

	if ColoredLetters == "" && char == ' ' {
		LetterLine = color + LetterLine + Reset
	}else {
		ColoredLetter := []rune(ColoredLetters)
		for i:=0; i<len(ColoredLetter); i++ {
			if char == ColoredLetter[i] {
				LetterLine = color + LetterLine + Reset
			}
		}
	}
	
	return LetterLine
}