/*
Выводит аргументы командной строки

Упражнение 1.1
Изменить программу так, чтобы она выводила также os.Args[0].

Упражнение 1.2
Изменить программу так, чтобы она выводила индекс и значение каждого аргумента
по одному аргументу в строке.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
