package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type input struct {
	text  string
	Input string
	Ans   string
}

type results struct {
	Success bool
	Data    string
	Date    string
}

func main() {
	// http.HandleFunc("/", mainPage)
	http.HandleFunc("/", word)
	http.HandleFunc("/letters", letters)

	http.Handle("/consonants/", http.StripPrefix("/consonants/", http.FileServer(http.Dir("consonants"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func chunk(s []string) [][]string {
	n := len(s) / 3
	var c [][]string
	for i := 0; i < n; i++ {
		c = append(c, []string{s[(i * 3)], s[(i*3)+1], s[(i*3)+2]})
	}
	return c
}

func getlist() ([]string, error) {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/consonants")
	if err != nil {
		return []string{"error"}, err
	}
	defer file.Close()
	names, err := file.Readdirnames(0)
	if err != nil {
		return []string{"error"}, err
	}
	// fmt.Println(names)
	for i, j := range names {
		names[i] = strings.ReplaceAll(j, ".png", "")
	}
	return names, nil
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
