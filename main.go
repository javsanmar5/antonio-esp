package main

import "os"

func main() {
    
    inFileName := os.Args[1]
    outFileName, err := ReadWriteFile(inFileName)
    check(err)

    Run(outFileName)
    DeleteFile(outFileName)
}
