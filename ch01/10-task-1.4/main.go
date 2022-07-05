/*
Выводит текст каждой строки, которая появляется в стандартном вводе более одного раза,
а также количество её появлений. Программа читает стандартный ввод или список именованных файлов.

Упражнение 1.4
Измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) bool {
	haveDup := false
	input := bufio.NewScanner(f)
	line := ""
	for input.Scan() {
		line = input.Text()
		counts[line]++
		if !haveDup && counts[line] > 1 {
			haveDup = true
		}
	}
	// Примечание: игнорируем потенциальные ошибки из input.Err()
	return haveDup
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		var filesWithDup []string
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			haveDup := countLines(f, counts)
			if haveDup {
				filesWithDup = append(filesWithDup, filename)
			}
			f.Close()
		}
		if len(filesWithDup) > 0 {
			fmt.Println("Duplicates found in files:")
			for _, filename := range filesWithDup {
				fmt.Println(filename)
			}
		} else {
			fmt.Println("Duplicates not found")
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
