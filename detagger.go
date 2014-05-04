package main

import "fmt"
import "io/ioutil"
import "regexp"
import "strings"
import "flag"

var re = regexp.MustCompile("\\s*(style|class)=(\\'|\").*(\\'|\")\\s*")

func stripFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
	}
	out := re.ReplaceAllString(string(data), "")
	err = ioutil.WriteFile(filename, []byte(out), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %s", err)
	}
}

func getHTMLFiles(directory, extension string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %s", err)
	}

	out := []string{}
	for _, file := range files {
		if strings.Contains(file.Name(), "."+extension) {
			out = append(out, directory+"/"+file.Name())
		}
	}
	return out
}

func main() {
	directoryPtr := flag.String("dir", ".", "directory to convert")
	fileExtension := flag.String("ext", "html", "file extension to convert")
	flag.Parse()

	files := getHTMLFiles(*directoryPtr, *fileExtension)
	fmt.Println(files)
	for _, file := range files {
		stripFile(file)
	}
}
