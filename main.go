package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
  "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	artistName, albumName := getNames(scanner)
  confirmInput(scanner, artistName, albumName)

	fmt.Print("Playlist URL: ")
	scanner.Scan()
	url := scanner.Text()

	path := makeDir(artistName, albumName)
  path = path + "%(title)s.%(ext)s"

  args := []string{
   "--extract-audio",
   "--audio-format", "mp3",
   "-i",
   "-o", path,
   url }

	cmd := exec.Command("youtube-dl", args...)
	pipe, _ := cmd.StdoutPipe()

	cmd.Start()
	ytdl := bufio.NewScanner(pipe)
	for ytdl.Scan() {
		fmt.Println(string(ytdl.Text()))
	}

  fmt.Println("Press Enter for new album, Control-C to quit")
  scanner.Scan()
  main()
}

func confirmInput(scanner *bufio.Scanner, artistName, albumName string) {

	fmt.Printf("\nArtist: %v Album: %v\n", artistName, albumName)
  fmt.Println("Is this information correct? Y to procced, N for re-entry")
  scanner.Scan()
  response := strings.TrimSpace(scanner.Text())
  if response == "Y"{
    return;

  } else if response == "N"{
    main();
  } else {
    fmt.Println("Unknown response")
    confirmInput(scanner, artistName, albumName)
  }

}

func makeDir(parent, child string) string {
	parentPath := filepath.Join(".", parent)
	os.MkdirAll(parentPath, os.ModePerm)
	childPath := filepath.Join(".", parentPath, child)
	os.MkdirAll(childPath, os.ModePerm)
	return childPath
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
