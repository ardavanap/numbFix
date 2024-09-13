package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global variables
const outputExtension string = ".txt"

var rawNumber string
var numbersSlice []string
var indexes []string
var slicedInput []string
var generatedNumbSlice []string
var lenOfInput int = len(rawNumber)
var generatedNumb string

func main() {

	greetingAndInput()
	var filename string = rawNumber + outputExtension
	numbToSlice(rawNumber)
	var lenOfX int = len(indexes)
	numbGenerator(lenOfX)
	result(filename)
	fmt.Println("your wordlist file: ", filename)
}

func greetingAndInput() {
	fmt.Println("_                 _______  ______   _______ _________         ")
	fmt.Println("( (    /||\\     /|(       )(  ___ \\ (  ____ \\__   __/|\\     /|")
	fmt.Println("|  \\  ( || )   ( || () () || (   ) )| (    \\/   ) (   ( \\   / )")
	fmt.Println("|   \\ | || |   | || || || || (__/ / | (__       | |    \\ (_) / ")
	fmt.Println("| (\\ \\) || |   | || |(_)| ||  __ (  |  __)      | |     ) _ (  ")
	fmt.Println("| | \\   || |   | || |   | || (  \\ \\ | (         | |    / ( ) \\ ")
	fmt.Println("| )  \\  || (___) || )   ( || )___) )| )      ___) (___( /   \\ )")
	fmt.Println("|/    )_)(_______)|/     \\||/ \\___/ |/       \\_______/|/     \\|")

	fmt.Println("\n\n-put 'x' for any unknown digit. e.g: 5639xx14xx21 . ")
	fmt.Print("\n	-Insert the number: ")
	fmt.Scan(&rawNumber)
}

func numbToSlice(input string) {

	var i int = 0
	for index, digit := range input {

		digit := strconv.QuoteRune(digit)          // converting every index tu 'rune'
		index := strconv.Itoa(index)               // and converting any digit to string
		numbersSlice = append(numbersSlice, digit) //adding every digit to 'numbersSlice'

		if digit == "'x'" { // adding index of any 'x' into 'indexes' slice

			indexes = append(indexes, index)

		}

		i = i + 1
	}

}

func numbGenerator(lenOfX int) {

	// generates every possible number for every x,

	for i := 0; len(strconv.Itoa(i)) <= lenOfX; i++ { // as long as length of 'i' variable is shorter or equal to length of all 'x' :
		lenOfI := len(strconv.Itoa(i))
		if lenOfI < lenOfX {
			generatedNumb = strings.Repeat("0", (lenOfX-lenOfI)) + strconv.Itoa(i) //zero padding

		} else {
			generatedNumb = strconv.Itoa(i)

		}
		generatedNumbSlice = append(generatedNumbSlice, generatedNumb)
	}
}

func result(filename string) {

	file, err := os.Create(filename) //making a file. name of the file is the user's input
	errorCheck(err)
	lenOfGeneratedNumbSlice := len(generatedNumbSlice)
	slicedInput = strings.Split(rawNumber, "")

	for i := 0; i < lenOfGeneratedNumbSlice; i++ { // this 'for' loop uses every generated number
		indexesCopy := indexes
		slicedGeneratedNumb := strings.Split(generatedNumbSlice[i], "")

		for j := 0; j < len(indexes); j++ { // this 'for' loop changes the 'x's with generated numbers.
			z, _ := strconv.Atoi(indexesCopy[j])

			slicedInput[z] = slicedGeneratedNumb[j]
		}

		fmt.Println(strings.Join(slicedInput, ""))
		file.WriteString(strings.Join(slicedInput, "") + "\n")

	}
	file.Close()
}
func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
