package main

import (
	"github.com/milanmayr/advent-of-code/2022/utils"
)

// THIS CHALLENGE IS INCOMPLETE

func main() {
	print("The sum of the total sizes of those directories that are at most 100,000 large is: ")
	println(totalSizeOfDirectoriesUnderOneHundredThousand())
}

type Directory struct {
	Name string
	ParentDirectory  *Directory
	Subdirectories []*Directory
	Files []*File
}

type File struct {
	Name string
	Size int
}

func newDirectory(name string) *Directory {
	return &Directory{
		Name: name,
		ParentDirectory:  nil,
		Subdirectories: make([]*Directory, 0),
		Files: []*File{},
	}
}

func totalSizeOfDirectoriesUnderOneHundredThousand() (result int) {
	input := utils.GetInput("input")

	// filesystem := newDirectory("/")

	currentDirectory := "/"
	previousDirectory := ""
	for _, line := range input {
		// handle commands
		if isCommand(line) {
			switch line[2:3] {
			case "cd":
				switch line[5:7] {
				case "..":
					currentDirectory = previousDirectory
				default:
					currentDirectory = line[5:]
				}
			case "ls": 
				// show dir contents
				continue
			}
		}

		if isDirectory(line) {

		}

		print(currentDirectory)
	}

	return result
}

func isCommand(text string) (bool) {
	return text[0:1] == "$"
}

func isDirectory(text string) (bool) {
	return text[0:3] == "dir"
}