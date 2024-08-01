package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type sitevars struct {
	Success bool
	Ans     string
	Opt1    string
	Opt2    string
	Opt3    string
	Opt4    string
}

func word(w http.ResponseWriter, r *http.Request) {
	list, _ := getlist()
	s := sitevars{}
	rand.Seed(time.Now().UnixNano())
	s.Opt1 = list[rand.Intn(len(list))]
	// fmt.Println(s)
	list = remove(list, s.Opt1)
	s.Opt2 = list[rand.Intn(len(list))]
	list = remove(list, s.Opt2)
	s.Opt3 = list[rand.Intn(len(list))]
	list = remove(list, s.Opt3)
	s.Opt4 = list[rand.Intn(len(list))]
	list = []string{s.Opt1, s.Opt2, s.Opt3, s.Opt4}
	s.Ans = list[rand.Intn(len(list))]
	// dir, _ := os.Getwd()
	// s.Ans = "\\consonants\\" + s.Ans
	s.Success = false
	tmpl := template.Must(template.ParseFiles("word.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, s)
		return
	}

	i := input{
		Input: r.FormValue("input"),
		Ans:   r.FormValue("img"),
	}
	fmt.Println(i)
	if i.Input == i.Ans {
		s.Success = true
	}

	tmpl.Execute(w, s)
}

func letters(w http.ResponseWriter, r *http.Request) {
	list, _ := getlist()
	s := sitevars{}
	rand.Seed(time.Now().UnixNano())
	s.Opt1 = list[rand.Intn(len(list))]
	// fmt.Println(s)
	list = remove(list, s.Opt1)
	s.Opt2 = list[rand.Intn(len(list))]
	list = remove(list, s.Opt2)
	s.Opt3 = list[rand.Intn(len(list))]
	list = remove(list, s.Opt3)
	s.Opt4 = list[rand.Intn(len(list))]
	list = []string{s.Opt1, s.Opt2, s.Opt3, s.Opt4}
	s.Ans = list[rand.Intn(len(list))]
	// dir, _ := os.Getwd()
	// s.Ans = "\\consonants\\" + s.Ans
	s.Success = false
	tmpl := template.Must(template.ParseFiles("letters.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, s)
		return
	}

	i := input{
		Input: r.FormValue("input"),
		Ans:   r.FormValue("img"),
	}
	fmt.Println(i)
	if i.Input == i.Ans {
		s.Success = true
	}

	tmpl.Execute(w, s)
}
