package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	reg, _ := regexp.Compile(`\/\*[^*]*\*+(?:[^\/*][^*]*\*+)*\/`)

	lst := readdir("/home/dreddsa/go/src/github.com/dreddsa5dies/codewars.com/")

	for _, val := range lst {
		if strings.HasSuffix(val, ".go") {
			file, err := ioutil.ReadFile(val)
			if err != nil {
				log.Fatalln(err)
			}
			text := string(file)

			comments := ""
			if reg.MatchString(text) {
				comments = reg.FindString(text)
			}

			comments = strings.TrimLeft(comments, "/*")
			comments = strings.TrimRight(comments, "*/")

			tmp, _ := os.OpenFile(strings.TrimSuffix(val, "kata.go")+"README.md", os.O_WRONLY|os.O_CREATE, 0600)
			defer tmp.Close()

			if _, err := tmp.Write([]byte(comments)); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func readdir(dir string) []string {
	tmp := []string{}

	// open directory
	dh, err := os.Open(dir)
	if err != nil {
		log.Fatalln(err)
	}
	defer dh.Close()

	// read files
	for {
		fis, err := dh.Readdir(10)
		if err == io.EOF {
			break
		}

		for _, fi := range fis {
			tmp = append(tmp, filepath.Join(dir, fi.Name()))

			// recursive read directories
			if fi.IsDir() {
				for _, val := range readdir(dir + "/" + fi.Name()) {
					tmp = append(tmp, val)
				}
			}
		}
	}

	return tmp
}
