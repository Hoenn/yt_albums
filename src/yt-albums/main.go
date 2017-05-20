package main

import (
	"bufio"
	"fmt"

	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"strings"

	id3 "github.com/mikkyang/id3-go"
)

type UserInput struct {
	artist, album, url string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var userInputs = []UserInput{}
	moreInput := true
	for moreInput {
		artistName, albumName, url := getInput(scanner)
		user := UserInput{artistName, albumName, url}
		userInputs = append(userInputs, user)
		moreInput = confirmInput("Add another artist/album (Y/N): ", scanner)
	}

	//Save main directory to return to
	for _, user := range userInputs {

		path := makeDirs(user.artist, user.album)
		ytdlpath := path + "/%(title)s.%(ext)s"

		args := []string{
			"--extract-audio",
			"--audio-format", "mp3",
			"-i",
			"-o", ytdlpath,
			user.url}

		cmd := exec.Command("youtube-dl", args...)
		pipe, _ := cmd.StdoutPipe()
		cmd.Start()
		ytdl := bufio.NewScanner(pipe)
		for ytdl.Scan() {
			fmt.Println(string(ytdl.Text()))
		}

		fmt.Println("*********Download Complete*********")

		updateID3Tags(path, user) //send additional args struct
	}

	fmt.Println("Press enter to run again, Control-C to quit")
	scanner.Scan()
	main()
}

func updateID3Tags(path string, user UserInput) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		mp3File, _ := id3.Open(path + "/" + f.Name())
		mp3File.SetArtist(user.artist)
		mp3File.SetAlbum(user.album)
		mp3File.Close()
	}
	fmt.Println("*********ID3 Tags Set*********")
}

func confirmInput(msg string, scanner *bufio.Scanner) bool {
	fmt.Print(msg)
	scanner.Scan()
	response := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if response == "y" {
		return true
	} else {
		return false
	}

}

func makeDirs(parent, child string) string {
	parentPath := filepath.Join(".", parent)
	os.MkdirAll(parentPath, os.ModePerm)
	childPath := filepath.Join(".", parentPath, child)
	os.MkdirAll(childPath, os.ModePerm)
	return childPath
}

func getInput(scanner *bufio.Scanner) (string, string, string) {

	fmt.Print("Enter artist name: ")
	scanner.Scan()
	artistName := scanner.Text()

	fmt.Print("Enter album name: ")
	scanner.Scan()
	albumName := scanner.Text()

	fmt.Print("Playlist URL: ")
	scanner.Scan()
	url := scanner.Text()

	return artistName, albumName, url

}
