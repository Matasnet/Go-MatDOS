
# MatDOS

MatDOS to prosty emulator systemu operacyjnego DOS napisany w języku Go. Program umożliwia interaktywne korzystanie z podstawowych funkcji, takich jak tworzenie plików, folderów, zapisywanie treści do pliku czy odczytywanie zawartości katalogów.  

### Wersja programu obecnie to 0.3


## Funkcje

### 1. Utwórz plik tekstowy lub folder  
```shell
create [ścieżka/nazwa_pliku_lub_folderu]
```
### 2. Zapisz treść do pliku  
```shell
save [ścieżka/nazwa_pliku] [treść]
```

### 3. Dodaj treść do pliku  
```shell
append [ścieżka/nazwa_pliku] [treść]
```

### 4. Uruchom plik tekstowy  
```shell
run [ścieżka/nazwa_pliku]
```

### 5. Wyświetl zawartość katalogu  
```shell
ls 
```

### 6. Utwórz folder  
```shell
mkdir [ścieżka/nazwa_folderu]
```

### 7. Zmień bieżący katalog  
```shell
cd [ścieżka]
```

### 8. Wyjście z programu  
```shell
exit
```
### 9. Wyświetlenie instrukcji
```shell
help
```
## Jak korzystać

1.  Uruchom program w terminalu.
2.  Wpisz polecenia, korzystając z powyższej listy funkcji.
3.  Program odpowiednio obsłuży wprowadzone polecenia.

## Wymagania

-   Go (wersja 1.20 lub nowsza)
