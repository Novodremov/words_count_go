package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

func countWords(text string, caseSensitive bool, deleteSpecialSymbols bool) map[string]int {
    specialSymbols := ".,!?:;()[]{}<>\"'`~/\\|@#$%^&*-_+"
    wordCount := make(map[string]int)

    words := strings.Fields(text)
    for _, word := range words {
        if !caseSensitive {
            word = strings.ToLower(word)
        }
        if deleteSpecialSymbols {
            word = strings.Trim(word, specialSymbols)
        }
        if word != "" {
            wordCount[word]++
        }
    }
    return wordCount
}

func readText(source *os.File) string {
    scanner := bufio.NewScanner(source)
    var text string
    for scanner.Scan() {
        text += scanner.Text() + " "
    }
    return text
}

func writeResult(wordCount map[string]int, outputFile string) {
    if outputFile == "" {
        for word, count := range wordCount {
            fmt.Printf("%s: %d\n", word, count)
        }
    } else {
        file, _ := os.Create(outputFile)
        defer file.Close()
        writer := bufio.NewWriter(file)
        for word, count := range wordCount {
            fmt.Fprintf(writer, "%s: %d\n", word, count)
        }
        writer.Flush()
    }
}

func promptInput(scanner *bufio.Scanner, prompt string) string {
    fmt.Print(prompt)
    scanner.Scan()
    return scanner.Text()
}

func main() {
    fileInPtr := flag.String("file_in", "", "Путь к файлу с текстом")
    fileOutPtr := flag.String("file_out", "", "Путь к файлу для сохранения результата")
    casePtr := flag.String("case", "", "Учитывать регистр (y/n)")
    lettersPtr := flag.String("letters", "", "Удалять спецсимволы по краям слова (y/n)")
    helpPtr := flag.Bool("help", false, "Показать справку")

    flag.Parse()

    if *helpPtr {
        fmt.Println("Использование:")
        fmt.Println("  -file_in <путь>   Указать путь к файлу с текстом")
        fmt.Println("  -file_out <путь>  Указать путь для сохранения результата")
        fmt.Println("  -case y|n         Учитывать регистр")
        fmt.Println("  -letters y|n      Удалять спецсимволы по краям слова")
        fmt.Println("  -help             Показать справку")
        return
    }

    scanner := bufio.NewScanner(os.Stdin)

    if *fileInPtr == "" {
        *fileInPtr = promptInput(scanner, "Введите путь к файлу с текстом (оставьте пустым для ввода вручную): ")
    }
    if *casePtr != "y" && *casePtr != "n" {
        *casePtr = promptInput(scanner, "Учитывать регистр? (y/n): ")
    }
    if *lettersPtr != "y" && *lettersPtr != "n" {
        *lettersPtr = promptInput(scanner, "Удалять спецсимволы по краям слова? (y/n): ")
    }
    if *fileOutPtr == "" {
        *fileOutPtr = promptInput(scanner, "Введите путь к файлу для сохранения результата (оставьте пустым для вывода на экран): ")
    }

    caseSensitive := *casePtr == "y"
    deleteSpecialSymbols := *lettersPtr == "y"

    var text string
    if *fileInPtr != "" {
        file, _ := os.Open(*fileInPtr)
        defer file.Close()
        text = readText(file)
    } else {
        text = readText(os.Stdin)
    }

    wordCount := countWords(text, caseSensitive, deleteSpecialSymbols)
    writeResult(wordCount, *fileOutPtr)
}
