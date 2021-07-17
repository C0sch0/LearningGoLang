
// Write a program which reads information from a file and represents it in a slice of structs.
//Assume that there is a text file which contains a series of names.
//Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
//
//Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
//Each field will be a string of size 20 (characters).
//
//Your program should prompt the user for the name of the text file.
//Your program will successively read each line of the text file and
//create a struct which contains the first and last names found in the file.
//Each struct created will be added to a slice, and after all lines have been read from the file,
//your program will have a slice containing one struct for each line in the file.
//After reading all lines from the file, your program should iterate through your slice of structs and
//print the first and last names found in each struct.


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
//Each field will be a string of size 20 (characters).

type Person struct {
	fname string
	lname string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main()  {
	var address string
	//Your program should prompt the user for the name of the text file.
	fmt.Println("Enter file address: ")
	fmt.Scan(&address)
	file, err := os.Open(address)
	check(err)
	print(err)
	//Your program will successively read each line of the text file and
	//create a struct which contains the first and last names found in the file.

	sliced_storage := make([]Person, 0, 10)
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Split(string(line), " ")

		//Each struct created will be added to a slice, and after all lines have been read from the file,
		//your program will have a slice containing one struct for each line in the file.

		pe := Person{fname: fields[0], lname: fields[1]}
		sliced_storage = append(sliced_storage, pe)
	}

	//After reading all lines from the file, your program should iterate through your slice of structs and
	//print the first and last names found in each struct.
	for _, pers := range sliced_storage {
		fmt.Printf("fname: [%-20s]   ;   lname: [%-20s]\n", pers.fname, pers.lname)
	}
}
