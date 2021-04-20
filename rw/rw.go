// Package rw provides Read/Write operations on text files.
package rw

import (
	"fmt"
	"log"
	"os"
)

func ReadFromFile(f string) string {
	xb, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(xb)
}

// Write to a text file. If a file with a given name doesn't exist
// then it is created.
func WriteToFile(fileName string, text []byte) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(text)
}

// Function Copy copies text from one file
// to another file.
// Target file is created if it doesn't exist
// If deleteSourceFile is set to true then the source file
// is deleted after copying.
func Copy(from, to string, deleteSourceFile bool) {
	target, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Reading from source file
	xb, err := os.ReadFile(from)
	if err != nil {
		log.Fatal(err)
	}
	// Writing to target file
	_, err = target.Write(xb)
	if err != nil {
		log.Fatal(err)
	}

	if deleteSourceFile {
		err = os.Remove(from)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Update adds given text to end of the file.
// It assumes the file already exists.
func Update(fileName string, text []byte) {
	target, err := os.OpenFile(fileName, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	xb, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	toWrite := []byte(string(xb) + string(text))

	target.Write(toWrite)
}

// Function Convert converts the contents of one file to the specified form
// and writes the result to an automatically generated text file.
// Format options: binary, hex, base10
func Convert(file, format string) {
	var toWrite string
	var targetFile string

	xb, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	sourceText := string(xb)

	switch format {
	case "binary":
		toWrite = convertToBinary(sourceText)
		targetFile = file + "_" + "binary.txt"
	case "hex":
		toWrite = convertTo(sourceText, "%x")
		targetFile = file + "_" + "hex.txt"
	case "base10":
		toWrite = convertToBase10(sourceText)
		targetFile = file + "_" + "base10.txt"
	}
	newFile, err := os.Create(targetFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = newFile.Write([]byte(toWrite))
	if err != nil {
		log.Fatal(err)
	}
}

//******* Helpers *******//

func convertTo(s, verb string) string {
	return fmt.Sprintf(verb+"\n", s)
}

func convertToBase10(s string) string {
	var output string
	for i := 0; i < len(s); i++ {
		output += fmt.Sprintf("%d", s[i]) + " "
	}
	return output
}

func convertToBinary(s string) string {
	var output string
	for _, c := range s {
		output = fmt.Sprintf("%s%b ", output, c)
	}
	return output
}
