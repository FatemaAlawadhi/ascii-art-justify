package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func CreateFile(FileName string ,FinalText []string) {
	Text := strings.Join(FinalText, "")
	file, err := os.Create(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// Write the text to the file
	_, err = writer.WriteString(Text)
	_, err = writer.WriteString("\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Flush the writer to ensure that all data has been written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}