package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadWriteFile(path string) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inFile.Close()

	outFileName := "tryout_temp.py"
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer outFile.Close()

	reader := bufio.NewReaderSize(inFile, 10*1024)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		check(err)

		_, err = outFile.WriteString(changeLine(line))
		check(err)
	}
}


func changeLine(line string) string {
    reservedWords := map[string]string{
        "QueVa":        "False",
        "Nanai":        "None",
        "Ji":           "True",
        "asopla":       "and",
        "perroSanche":  "break",
        "tiktok":       "class",
        "biblio":       "continue",
        "carajote":     "def",
        "agenda2030":   "del",
        "porElCuloQue": "elif",
        "que":          "else",
        "incompetente": "except",
        "JAJAJAJA":     "finally",
        "unaDuda":      "for",
        "queVuelva":    "from",
        "porElCulo":    "if",
        "deLaMili":     "import",
        "porfi":        "in",
        "eIgua":        "is",
        "mujer":        "not",
        "piola":        "or",
        "niCaso":       "pass",
        "queGuapoEre":  "return",
        "rojo":         "try",
        "vivaVox":      "while",
    }
    splited := strings.Split(line, " ")
    var o []string

    for _, word := range splited {
        if antoWord, ok := reservedWords[word]; ok {
            o = append(o, antoWord)
        } else {
            o = append(o, word)
        }
    }
    return strings.Join(o, " ") 
}
 
