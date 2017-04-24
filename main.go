package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type UserInput struct {
  artist, album string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	artistName, albumName := getNames(scanner)
  user := UserInput{artistName, albumName}
	confirmInput(scanner, user)

  url := getUrlInput(scanner)

	path := makeDir(user.artist, user.album)
	path = path + "/%(title)s.%(ext)s"

	args := []string{
		"--extract-audio",
		"--audio-format", "mp3",
		"-i",
		"-o", path,
		url}

	cmd := exec.Command("youtube-dl", args...)
	pipe, _ := cmd.StdoutPipe()
	cmd.Start()
	ytdl := bufio.NewScanner(pipe)
	for ytdl.Scan() {
		fmt.Println(string(ytdl.Text()))
	}
  fmt.Println("*********Download Complete*********")
	fmt.Println("Press Enter for new album, Control-C to quit")
	scanner.Scan()
	main()
}

func confirmInput(scanner *bufio.Scanner, u UserInput) {

	fmt.Printf("\nArtist: %v Album: %v\n", u.artist, u.album)
	fmt.Print("Is this information correct? Y to procced, N for re-entry: ")
	scanner.Scan()
	response := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if response == "y" {
    fmt.Println()
		return
	} else if response == "n" {
    fmt.Println()
		main()
	} else {
		fmt.Println("Unknown response")
		confirmInput(scanner, u)
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

func getUrlInput(scanner *bufio.Scanner) string {
	fmt.Print("Playlist URL: ")
	scanner.Scan()
	return scanner.Text()
}
