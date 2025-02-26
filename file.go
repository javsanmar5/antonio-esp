package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadWriteFile(inFileName string) (string, error) {
	inFile, err := os.Open(inFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return "", err
	}
	defer inFile.Close()

    outFileName := strings.Split(inFileName, ".")[0] + "_temp.py"
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return "", err
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

    return outFileName, err
}

func DeleteFile(outFileName string) {
    err := os.Remove(outFileName)
    check(err)
}
