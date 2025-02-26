package main

func main() {
    path := "tryout.esp"
    outFileName, err := ReadWriteFile(path)
    check(err)

    Run(outFileName)
}
