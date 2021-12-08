package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// var numberfornewline int

// func checkErr(e error) {
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// }

// // checking for errors other than io.EOF
// func checkErrNoEOF(e error) {
// 	if e != nil && e != io.EOF {
// 		log.Fatal(e)
// 	}
// }

// // scaning 8 lines starting from startLine
// func scanChar(r io.Reader, startLine int) ([]string, error) {
// 	lineScanner := bufio.NewScanner(r)
// 	bigCharLines := []string{}
// 	curLine := 0
// 	linesAdded := 0
// 	for lineScanner.Scan() {
// 		curLine++ // coz the first line of txt is 1
// 		if curLine == startLine {
// 			// scan 8 lines
// 			for sc := 0; sc < 8; sc++ {
// 				bigCharLines = append(bigCharLines, lineScanner.Text())
// 				linesAdded++
// 				lineScanner.Scan() // advance to the next line
// 			}
// 		}
// 	}
// 	// fmt.Println(bigCharLines) // this will print stuff crumbled tgt in a single line
// 	return bigCharLines, io.EOF
// }

// func printBigChar(chMap *map[byte][]string, inpBSlice []byte) []string {
// 	chLine := ""
// 	chLine2 := ""
// 	chLine3 := ""
// 	chLine4 := ""
// 	chLine5 := ""
// 	chLine6 := ""
// 	chLine7 := ""
// 	chLine8 := ""

// 	chLine += string((*chMap)[inpBSlice[0]][0])
// 	chLine2 += string((*chMap)[inpBSlice[0]][1])
// 	chLine3 += string((*chMap)[inpBSlice[0]][2])
// 	chLine4 += string((*chMap)[inpBSlice[0]][3])
// 	chLine5 += string((*chMap)[inpBSlice[0]][4])
// 	chLine6 += string((*chMap)[inpBSlice[0]][5])
// 	chLine7 += string((*chMap)[inpBSlice[0]][6])
// 	chLine8 += string((*chMap)[inpBSlice[0]][7])

// 	result := []string{chLine, chLine2, chLine3, chLine4, chLine5, chLine6, chLine7, chLine8}

// 	// fmt.Println(chLine)
// 	// fmt.Println(chLine2)
// 	// fmt.Println(chLine3)
// 	// fmt.Println(chLine4)
// 	// fmt.Println(chLine5)
// 	// fmt.Println(chLine6)
// 	// fmt.Println(chLine7)
// 	// fmt.Println(chLine8)

// 	return result
// }

// func AsciiArt(s string, bannertype string) []string {
// 	// if len(os.Args) != 2 {
// 	// 	fmt.Println("Arg amount is not right")
// 	// }
// 	var result2 []string

// 	// read input str
// 	var inputStrSlices []string

// 	// fmt.Println("input: ", inputStr)
// 	inputrune := []rune(s)
// 	s1 := s
// 	// case when \n at the end of the input
// 	if len(s) == 4 {
// 		if inputrune[len(s)-2] == '\\' && inputrune[len(s)-1] == 'n' {
// 			if last := len(s1) - 1; last >= 0 && s1[last] == 'n' {
// 				s1 = s1[:last] // remove the last n
// 			}
// 			s2 := s1
// 			if last := len(s2) - 1; last >= 0 && s2[last] == '\\' {
// 				s2 = s2[:last] // remove the 2nd last \
// 			}
// 			// convert s2 (string) to slice of string (with only itself in the slice)
// 			inputStrSlices = strings.Split(s2, " ")
// 			numberfornewline++ // increment the counter
// 		} else {
// 			inputStrSlices = strings.Split(s, "\\n") // normal case
// 		}
// 	} else {
// 		inputStrSlices = strings.Split(s, "\\n")
// 	}
// 	for _, inputSlice := range inputStrSlices {
// 		// process the str
// 		inputBSlice := []byte(inputSlice)
// 		// fmt.Println(inputBSlice)
// 		charMap := make(map[byte][]string)

// 		for inp := 0; inp < len(inputBSlice); inp++ {
// 			// find the corresponding line num
// 			startLine := (int(inputBSlice[inp])-32)*9 + 2
// 			// fmt.Println("startLine: ", startLine)
// 			// scan the reqired lines in from the txt file
// 			fread, err := os.Open(bannertype + ".txt")
// 			checkErr(err)
// 			defer fread.Close()

// 			scanner := bufio.NewScanner(fread)

// 			scanner.Split(bufio.ScanBytes)

// 			// scan 8 lines starting from startLine from the txt file
// 			bigChar, err := scanChar(fread, startLine)
// 			checkErrNoEOF(err)
// 			if len(bigChar) != 8 {
// 				fmt.Println("Wrong number of lines read")
// 			}

// 			charMap[inputBSlice[inp]] = bigChar

// 		}

// 		// }

// 		result2 = printBigChar(&charMap, inputBSlice)

// 	}
// 	return result2
// }
