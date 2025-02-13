package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func wordsCounting(text string) map[string]int {
    specialSymbols := ".,!?:;()[]{}<>\"'`~/\\|@#$%^&*-_+"
    wordCount := make(map[string]int)
    words := strings.Fields(strings.ToLower(text))  // Регистр слов не учитываем

    for _, word := range words {
        word = strings.Trim(word, specialSymbols)
        if word != "" {
            wordCount[word]++
        }
    }

    return wordCount
}

func main() {
    fmt.Println("Введите текст (для завершения ввода нажмите Ctrl+Z (Windows) или Ctrl+D (Linux/Mac)):")
    scanner := bufio.NewScanner(os.Stdin)
    var textToCount string

    for scanner.Scan() {
        textToCount += scanner.Text() + " "
    }

    wordCount := wordsCounting(textToCount)

    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
