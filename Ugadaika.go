/*
игра, где надо угадать число от 0 до 100
на это есть 10 попыток1
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Ya zagadal chislo")
	fmt.Println("Poprobuy ugadat'")
	chislo := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)
	uspeh := false // альтернатива
	for popitki := 0; popitki <= 9; popitki++ {

		fmt.Println("(Ostalos", 10-popitki, "popitok)")
		fmt.Print("Vvedi chislo: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSuffix(input, "\n")
		vvod, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if vvod < chislo {
			fmt.Println("Malo")
		} else if vvod > chislo {
			fmt.Println("Mnogo")
		} else {
			uspeh = true // альтернатива
			println("Yeah!!!")
			break
		}

	}
	if !uspeh { // альтернатива
		fmt.Println("Soryan, zakonchilis popitki")
		fmt.Println("Bilo zagadano", chislo)
		/*if popitki == 9 {
			fmt.Println("Soryan, zakonchilis popitki")
			fmt.Println("Bilo zagadano", chislo)
		}*/
	}

}
q