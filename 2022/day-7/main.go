package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/milanmayr/advent-of-code/2022/utils"
)

// THIS CHALLENGE IS INCOMPLETE

const explanation = "AKA your code is broken."

func main() {
	print("The sum of the total sizes of those directories that are at most 100,000 large is: ")
	println(totalSizeOfDirectoriesUnderOneHundredThousand())
}

// Represents a directory
type Directory struct {
	Name            string
	ParentDirectory *Directory
	Subdirectories  map[string]*Directory
	Files           map[string]int
	Size            int
}

func totalSizeOfDirectoriesUnderOneHundredThousand() int {
	input := utils.GetInput("input")
	filesystem := createFilesystem(input[1:])
	sum := 0
	directorySizes(&filesystem, &sum)

	return sum
}

// Checks if line is a command
func isCommand(line string) bool {
	isCommand, err := regexp.MatchString("^\\$.*", line)
	if err != nil {
		panic(err)
	}
	return isCommand
	// return line[0:1] == "$"
}

// Checks if line is a directory
func isDirectory(line string) bool {
	isDirectory, err := regexp.MatchString("^dir.*", line)
	if err != nil {
		panic(err)
	}
	return isDirectory
	// return line[0:3] == "dir"
}

// Checks if line is a file
func isFile(line string) bool {
	isfile, err := regexp.MatchString("^[0-9].*$", line)
	if err != nil {
		panic(err)
	}
	return isfile
}

func createEmptyDir(name string, parentDirectory *Directory) *Directory {
	directory := Directory{
		Name:            name,
		Files:           make(map[string]int, 0),
		Subdirectories:  make(map[string]*Directory, 0),
		ParentDirectory: parentDirectory,
	}
	return &directory
}

func createFilesystem(input []string) (filesystem Directory) {
	filesystem = Directory{
		Name:           "/",
		Files:          make(map[string]int, 0),
		Subdirectories: make(map[string]*Directory, 0),
	}
	currentDirectory := &filesystem

	for _, line := range input {
		if isCommand(line) { // handle commands
			switch line[2:4] {
			case "cd":
				switch line[5:] {
				case "..":
					currentDirectory = currentDirectory.ParentDirectory
				default:
					directoryName := strings.Split(line, " ")[2]
					currentDirectory = currentDirectory.Subdirectories[directoryName]
				}
			case "ls":
				// show dir contents
				continue
			}
		} else if isDirectory(line) { // handle directories
			// Add directory to current directory
			directoryName := strings.Split(line, " ")[1]
			currentDirectory.Subdirectories[directoryName] = createEmptyDir(directoryName, currentDirectory)
		} else if isFile(line) { // handle files
			fileSize, err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				panic("Error parsing file size." + explanation)
			}

			fileName := strings.Split(line, " ")[1]

			currentDirectory.Files[fileName] = fileSize
		} else {
			panic("Line not recognized as command, directory, or file!" + explanation)
		}

	}

	return filesystem
}

func directorySizes(dir *Directory, sum *int) {
	for _, file := range dir.Files {
		dir.Size += file
	}
	if len(dir.Subdirectories) > 0 {
		for _, subdir := range dir.Subdirectories {
			directorySizes(subdir, sum)
			dir.Size += subdir.Size
		}
	}
	if dir.Size <= 100000 {
		*sum += dir.Size
	}
}
