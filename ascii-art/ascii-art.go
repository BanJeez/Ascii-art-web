package asciiArt

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	//"flag"
)

var (
	numberfornewline int = 0
	banner           string
)

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
		fmt.Print(chLine)
		fmt.Println("")
	}
	if numberfornewline > 0 {
		fmt.Print("\n")
	}
}

func AsciiConverer(text string, ban string) {
	banner = ban

	// Flags
	// outputString := flag.String("output", "", "a string")
	// flag.Parse()
	// read input str
	var inputStrSlices []string
	inputStr := text
	// fmt.Println("input: ", inputStr)
	inputrune := []rune(inputStr)
	s1 := inputStr
	if inputrune[len(inputStr)-2] == '\\' && inputrune[len(inputStr)-1] == 'n' {
		if last := len(s1) - 1; last >= 0 && s1[last] == 'n' {
			s1 = s1[:last]
		}
		s2 := s1
		if last := len(s2) - 1; last >= 0 && s2[last] == '\\' {
			s2 = s2[:last]
		}
		inputStrSlices = strings.Split(s2, "%&*(&()YKUYfa3432) 45sdt4") // just something unique
		numberfornewline++
	} else {
		inputStrSlices = strings.Split(inputStr, "\\n")
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
		// }
		printBigChar(&charMap, inputBSlice)
	}
}
