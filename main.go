package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	artistName, albumName := getNames(scanner)
	//is this Correct prompt?
	fmt.Printf("\nArtist: %v Album: %v\n", artistName, albumName)

	//If no loop

	artistPath := filepath.Join(".", artistName)
	os.MkdirAll(artistPath, os.ModePerm)
	albumPath := filepath.Join(".", artistName, albumName)
	os.MkdirAll(albumPath, os.ModePerm)
}

func getNames(scanner *bufio.Scanner) (string, string) {

	fmt.Print("Enter artist name: ")
	scanner.Scan()
	artistName := scanner.Text()

	fmt.Print("Enter album name: ")
	scanner.Scan()
	albumName := scanner.Text()

	return artistName, albumName

}
