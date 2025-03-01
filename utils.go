package main

import "strings"

var reservedWords = map[string]string{
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


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func changeLine(line string) string {
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

