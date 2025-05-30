package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
)

func filterLinesByRegex(inputFilePath, regexPattern, outputFilePath string) {
    inputFile, err := os.Open(inputFilePath)
    if err != nil { // +1
	log.Fatalf("❌ Chyba při otevírání vstupního souboru '%s': %v", inputFilePath, err)
    }
    defer func() {
	if err := inputFile.Close(); err != nil { // +1
	    log.Printf("⚠️ Chyba při zavírání vstupního souboru: %v", err)
	}
    }()
    outputFile, err := os.Create(outputFilePath)
    if err != nil { // +1
	log.Fatalf("❌ Chyba při vytváření výstupního souboru '%s': %v", outputFilePath, err)
    }
    defer func() {
	if err := outputFile.Close(); err != nil { // +1
	    log.Printf("⚠️ Chyba při zavírání výstupního souboru: %v", err)
	}
    }()
    writer := bufio.NewWriter(outputFile)
    defer writer.Flush()

    pattern, err := regexp.Compile(regexPattern)
    if err != nil { // +1
	log.Fatalf("❌ Chyba v regulárním výrazu: %v", err)
    }
    scanner := bufio.NewScanner(inputFile)
    for lineNumber := 1; scanner.Scan(); lineNumber++ { // +1
	line := scanner.Text()
	if pattern.MatchString(line) { // +1
	    if _, err := writer.WriteString(line + "\n"); err != nil {
		log.Fatalf("❌ Chyba při zápisu na řádku %d: %v", lineNumber, err)
	    }
	}
	}
	    if err := scanner.Err(); err != nil { // +1
	    log.Fatalf("❌ Chyba při čtení souboru: %v", err)
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
