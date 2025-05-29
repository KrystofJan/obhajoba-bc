
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func filterLinesByRegex(inputFilePath, regexPattern, outputFilePath string) {
    inputFile, err := os.Open(inputFilePath)
    if err != nil {
	fmt.Printf("❌ Chyba při otevírání vstupního souboru '%s': %v\n", inputFilePath, err)
	return
    }
    defer func() {
	if err := inputFile.Close(); err != nil {
	    fmt.Printf("⚠️ Chyba při zavírání vstupního souboru: %v\n", err)
	}
    }()

    outputFile, err := os.Create(outputFilePath)
    if err != nil {
	fmt.Printf("❌ Chyba při vytváření výstupního souboru '%s': %v\n", outputFilePath, err)
	return
    }
    defer func() {
	if err := outputFile.Close(); err != nil {
	    fmt.Printf("⚠️ Chyba při zavírání výstupního souboru: %v\n", err)
	}
    }()

    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(outputFile)
    defer writer.Flush()

    pattern, err := regexp.Compile(regexPattern)
    if err != nil {
	fmt.Printf("❌ Chyba v regulárním výrazu: %v\n", err)
	return
    }

    lineNumber := 0
    for scanner.Scan() {
	line := scanner.Text()
	lineNumber++
	match := pattern.MatchString(line)
	if match {
	    _, err := writer.WriteString(line + "\n")
	    if err != nil {
		fmt.Printf("❌ Chyba při zápisu na řádku %d: %v\n", lineNumber, err)
		return
	    }
	}
    }

    if err := scanner.Err(); err != nil {
	fmt.Printf("❌ Chyba při čtení souboru: %v\n", err)
    }
}

func main() {
    if len(os.Args) != 4 {
	fmt.Println("❗ Použití: <vstupní_soubor> <výstupní_soubor> <regulární_výraz>")
	return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]
    regexPattern := os.Args[3]

    fmt.Printf("📂 Zpracovávám soubor '%s' s výrazem '%s'\n", inputFile, regexPattern)
    filterLinesByRegex(inputFile, regexPattern, outputFile)
    fmt.Println("✅ Hotovo. Výsledky jsou zapsány do výstupního souboru.")
}
