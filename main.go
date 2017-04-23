package main

import(
  "bufio"
  "os"
  "fmt"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  artistName, albumName := getNames(scanner)
  //is this Correct prompt?
  fmt.Printf("\nArist: %v Album: %v\n", artistName, albumName)

  //If no loop


  //If yes then...
}

func getNames (scanner *bufio.Scanner) (string,string) {

  fmt.Print("Enter artist name: ")
  scanner.Scan()
  artistName := scanner.Text();

  fmt.Print("Enter album name: ")
  scanner.Scan()
  albumName := scanner.Text();

  return artistName, albumName

}
