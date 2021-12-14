package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var numberfornewline int = 0

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// checking for errors other than io.EOF
func checkErrNoEOF(e error) {
	if e != nil && e != io.EOF {
		log.Fatal(e)
	}
}

// scaning 8 lines starting from startLine
func scanChar(r io.Reader, startLine int) ([]string, error) {
	lineScanner := bufio.NewScanner(r)
	bigCharLines := []string{}
	curLine := 0
	linesAdded := 0
	for lineScanner.Scan() {
		curLine++ // coz the first line of txt is 1
		if curLine == startLine {
			// scan 8 lines
			for sc := 0; sc < 8; sc++ {
				bigCharLines = append(bigCharLines, lineScanner.Text())
				linesAdded++
				lineScanner.Scan() // advance to the next line
			}
		}
	}
	// fmt.Println(bigCharLines) // this will print stuff crumbled into a single line
	return bigCharLines, io.EOF
}

func printBigChar(chMap *map[byte][]string, inpBSlice []byte) {
	for l := 0; l < 8; l++ {
		chLine := ""
		for ch := 0; ch < len(inpBSlice); ch++ {
			chLine += string((*chMap)[inpBSlice[ch]][l])
		}

		// fmt.Print(chLine)
		// fmt.Println("")
		// fmt.Fprintln(w, chLine)
		// fmt.Fprintln(w, "")

		// if numberfornewline > 0 {
		// 	fmt.Print("\n") // not recognised in html, at least not outside of <pre>
		// }
		// res1 := strings.Split(os.Args[3], "=") // need to be modified
		
		arraychline := []string{chLine}
		// fmt.Print(arraychline)

		file, err := os.OpenFile("artwork.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(file)

		for _, data := range arraychline {
			_, _ = datawriter.WriteString(data + "\n")
		}

		datawriter.Flush()
		file.Close()

	}
}

func AsciiArt(inputStr string, banner string) {
	var inputStrSlices []string
	inputrune := []rune(inputStr)
	s1 := inputStr
	if len(inputStr) < 2 {
		inputStrSlices = strings.Fields(s1)
	} else if inputrune[len(inputStr)-2] == '\\' && inputrune[len(inputStr)-1] == 'n' {
		if last := len(s1) - 1; last >= 0 && s1[last] == 'n' {
			s1 = s1[:last] // remove the last n
		}
		s2 := s1
		if last := len(s2) - 1; last >= 0 && s2[last] == '\\' {
			s2 = s2[:last] // remove the 2nd last \
		}
		// convert s2 (string) to slice of string (with only itself in the slice)
		inputStrSlices = strings.Split(s2, " ")
		numberfornewline++ // increment the counter
	} else {
		inputStrSlices = strings.Split(inputStr, "\\n") // normal case
	}
	for _, inputSlice := range inputStrSlices {
		// process the str
		inputBSlice := []byte(inputSlice)
		// fmt.Println(inputBSlice)
		charMap := make(map[byte][]string)

		for inp := 0; inp < len(inputBSlice); inp++ {
			// find the corresponding line num
			startLine := (int(inputBSlice[inp])-32)*9 + 2
			// fmt.Println("startLine: ", startLine)
			// scan the reqired lines in from the txt file
			fread, err := os.Open(banner + ".txt")
			checkErr(err)
			defer fread.Close()

			scanner := bufio.NewScanner(fread)

			scanner.Split(bufio.ScanBytes)

			// scan 8 lines starting from startLine from the txt file
			bigChar, err := scanChar(fread, startLine)
			checkErrNoEOF(err)
			if len(bigChar) != 8 {
				fmt.Println("Wrong number of lines read")
			}

			charMap[inputBSlice[inp]] = bigChar
		}
		printBigChar(&charMap, inputBSlice)
	}
}
