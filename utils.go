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
    spaceSplitted := strings.Split(line, " ")
    var o []string

    for _, part := range spaceSplitted {
        colonSplitted := strings.Split(part, ":")
        for i, word := range colonSplitted {
            if antoWord, ok := reservedWords[word]; ok {
                colonSplitted[i] = antoWord
            }
        }
        o = append(o, strings.Join(colonSplitted, ":"))
    }
    return strings.Join(o, " ") 
}

