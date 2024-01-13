package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func Banner(FileType string) []string {
	var Banner []string
	if FileType == "standard" {
		FileType = "Banner/standard.txt"
	} else if FileType == "shadow" {
		FileType = "Banner/shadow.txt"
	} else if FileType == "thinkertoy" {
		FileType = "Banner/thinkertoy.txt"
	}else if FileType == "No Banner"{
		FileType = "Banner/standard.txt"
	} else {
		fmt.Println("Banner can be standard, shadow and thinkertoy")
		fmt.Println("If the Banner is not specified, it will be printed as standard")
		FileType = "Banner/standard.txt"
	}
	
	file, err := os.Open(FileType)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	defer file.Close()
	// Create a scanner for the file
	scanner := bufio.NewScanner(file)
	// Read each line of the file
	for scanner.Scan() {
		Banner = append(Banner, scanner.Text())
	}

	return Banner
}