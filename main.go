package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var currentDir string // Zmienna globalna przechowująca aktualny katalog

func main() {
	currentDir, _ = os.Getwd() // Początkowy katalog to katalog roboczy aplikacji
	fmt.Println("Witaj w MatDOS!")
	version := 0.3
	fmt.Printf("Wersja %v\n", version)
	fmt.Println("Wpisz 'help' aby uzyskać listę dostępnych poleceń.")

	for {
		fmt.Printf("MatDOS [%s] >> ", currentDir)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Błąd odczytu wejścia:", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch {
		case input == "help":
			fmt.Println("Dostępne polecenia:")
			fmt.Println(" - help: Wyświetl tę wiadomość pomocy")
			fmt.Println(" - exit: Zakończ program")
			fmt.Println(" - create [ścieżka/nazwa_pliku]: Utwórz nowy plik tekstowy lub folder")
			fmt.Println(" - save [ścieżka/nazwa_pliku] [treść]: Zapisz treść do pliku")
			fmt.Println(" - append [ścieżka/nazwa_pliku] [treść]: Dodaj treść do pliku")
			fmt.Println(" - run [ścieżka/nazwa_pliku]: Uruchom plik tekstowy")
			fmt.Println(" - ls [ścieżka]: Wyświetl zawartość katalogu")
			fmt.Println(" - mkdir [ścieżka/nazwa_folderu]: Utwórz folder")
			fmt.Println(" - cd [ścieżka]: Zmień bieżący katalog")
		case input == "exit":
			fmt.Println("Do widzenia!")
			os.Exit(0)
		default:
			processCommand(input)
		}
	}
}

func processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	switch command {
	case "create":
		createFileOrFolder(parts)
	case "save":
		saveToFile(parts)
	case "append":
		appendToFile(parts)
	case "run":
		runFile(parts)
	case "ls":
		listFiles(parts)
	case "mkdir":
		createFolder(parts)
	case "cd":
		changeDirectory(parts)
	default:
		fmt.Println("Nieznane polecenie. Wpisz 'help' aby uzyskać listę poleceń.")
	}
}

func createFileOrFolder(parts []string) {
    if len(parts) < 2 {
        fmt.Println("Podaj nazwę pliku lub folderu.")
        return
    }

    path := getFullPath(parts[1])

    if strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\") {
        createFolder([]string{path})
    } else {
        createFile(parts)
    }
}


func createFolder(parts []string) {
	var folderPath string
	if len(parts) < 2 {
		// Jeśli nie podano ścieżki, użyj bieżącego katalogu
		folderPath = currentDir
	} else {
		folderPath = getFullPath(parts[1])
	}

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println("Błąd tworzenia folderu:", err)
		return
	}

	fmt.Printf("Utworzono folder '%s'.\n", folderPath)
}

func createFile(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Podaj nazwę pliku.")
		return
	}

	fileName := parts[1]
	filePath := getFullPath(fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Błąd tworzenia pliku:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Utworzono plik '%s'.\n", filePath)
}

func saveToFile(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Podaj nazwę pliku i treść do zapisania.")
		return
	}

	fileName := parts[1]
	filePath := getFullPath(fileName)

	content := strings.Join(parts[2:], " ")

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Błąd zapisu do pliku:", err)
		return
	}
	fmt.Printf("Zapisano treść do pliku '%s'.\n", filePath)
}

func appendToFile(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Podaj nazwę pliku i treść do dodania.")
		return
	}

	fileName := parts[1]
	filePath := getFullPath(fileName)

	content := strings.Join(parts[2:], " ") + "\n"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Błąd otwarcia pliku do dodawania:", err)
		return
	}
	defer file.Close()
	
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Błąd dodawania treści do pliku:", err)
		return
	}

	fmt.Printf("Dodano treść do pliku '%s'.\n", filePath)
}

func runFile(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Podaj nazwę pliku do uruchomienia.")
		return
	}

	fileName := parts[1]
	filePath := getFullPath(fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Błąd odczytu pliku:", err)
		return
	}

	fmt.Println("Treść pliku:")
	fmt.Println(string(content))
}

func listFiles(parts []string) {
	var dirPath string
	if len(parts) > 1 {
		dirPath = getFullPath(parts[1])
	} else {
		dirPath = currentDir
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Błąd odczytu katalogu:", err)
		return
	}

	fmt.Printf("Zawartość katalogu '%s':\n", dirPath)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func changeDirectory(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Podaj ścieżkę do katalogu.")
		return
	}

	newDir := getFullPath(parts[1])
	err := os.Chdir(newDir)
	if err != nil {
		fmt.Println("Błąd zmiany katalogu:", err)
		return
	}

	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Błąd pobierania aktualnego katalogu:", err)
		return
	}

	fmt.Printf("Zmieniono bieżący katalog na '%s'.\n", currentDir)
}

func getFullPath(relativePath string) string {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Błąd uzyskiwania pełnej ścieżki:", err)
		return relativePath
	}
	return absPath
}
