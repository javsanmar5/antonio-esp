package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadWriteFile(path string) (string, error) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return "", err
	}
	defer inFile.Close()

	outFileName := "tryout_temp.py"
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
 
