package main

import (
	"bufio"
	"fmt"

	"os"
	"os/exec"
  "path/filepath"
  "io/ioutil"

	"strings"

  id3 "github.com/mikkyang/id3-go"
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

	path := makeDirs(user.artist, user.album)
	ytdlpath:= path + "/%(title)s.%(ext)s"

	args := []string{
		"--extract-audio",
		"--audio-format", "mp3",
		"-i",
		"-o", ytdlpath,
		url}

	cmd := exec.Command("youtube-dl", args...)
	pipe, _ := cmd.StdoutPipe()
	cmd.Start()
	ytdl := bufio.NewScanner(pipe)
	for ytdl.Scan() {
		fmt.Println(string(ytdl.Text()))
	}

  fmt.Println("*********Download Complete*********")

  //Update ID3 tags//
  mainDir,_ := os.Getwd()
  os.Chdir(path)

  files, _ := ioutil.ReadDir("./")
  for _, f := range files {
    mp3File, _ := id3.Open("./"+f.Name())
    mp3File.SetArtist(user.artist)
    mp3File.SetAlbum(user.album)
    fmt.Println("ID3 tags set")
    fmt.Println(mp3File.Artist())
    mp3File.Close()
  }

  os.Chdir(mainDir)

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

func makeDirs(parent, child string) string {
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
