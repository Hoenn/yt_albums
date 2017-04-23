package main

import (
	"bufio"
	"fmt"
	"os"
  "os/exec"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	artistName, albumName := getNames(scanner)
	//is this Correct prompt?
	fmt.Printf("\nArtist: %v Album: %v\n", artistName, albumName)
	//If no loop

  path := makeDir(artistName, albumName)
  os.Chdir(path)

  //youtube-dl -o path
  //youtube-dl --extract-audio --audio-format mp3 -i <SongName/PlaylistLink>

  fmt.Print("Playlist URL: ")
  scanner.Scan()
  url := scanner.Text()

  cmd := exec.Command("youtube-dl", url)
  s,e := cmd.Output()
  fmt.Println(s, e)

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
