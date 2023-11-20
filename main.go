package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var words = []string{"GOLANG", "PYTHON", "JAVASCRIPT", "JAVA"}

var chosenWord string
var guessedLetters []string
var attemptsLeft = 10
var tmpl *template.Template

func init() {
	rand.Seed(42)
	chosenWord = pickRandomWord(words)
}

func main() {
	var err error
	tmpl, err = template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
		return
	}

	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		resetGame()
		tmpl.ExecuteTemplate(w, "accueil", nil)
	})

	http.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "rules", nil)
	})

	http.HandleFunc("/hangman", hangmanHandler)
	http.HandleFunc("/guess", guessHandler)
	http.HandleFunc("/result", resultHandler)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("Serveur lancé sur : http://localhost:8080/accueil")
	http.ListenAndServe(":8080", nil)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	difficulty := r.URL.Query().Get("difficulty")

	wordGuessed := isWordGuessed(chosenWord, guessedLetters)
	lost := attemptsLeft <= 0 && !wordGuessed

	if wordGuessed || lost {
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	data := struct {
		ChosenWord     string
		GuessedLetters []string
		AttemptsLeft   int
		Difficulty     string
		Won            bool
		Lost           bool
	}{
		ChosenWord:     maskWord(chosenWord, guessedLetters),
		GuessedLetters: guessedLetters,
		AttemptsLeft:   attemptsLeft,
		Difficulty:     difficulty,
		Won:            wordGuessed,
		Lost:           lost,
	}

	err := tmpl.ExecuteTemplate(w, "game", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	wordGuessed := isWordGuessed(chosenWord, guessedLetters)
	lost := attemptsLeft <= 0 && !wordGuessed

	data := struct {
		Won        bool
		Lost       bool
		ChosenWord string
	}{
		Won:        wordGuessed,
		Lost:       lost,
		ChosenWord: chosenWord,
	}

	err := tmpl.ExecuteTemplate(w, "result", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	letter := strings.ToUpper(r.FormValue("letter"))
	if letter == "" {
		http.Error(w, "Pas de lettre reçue", http.StatusBadRequest)
		return
	}

	if isLetterInWord(letter, chosenWord) && !isLetterAlreadyGuessed(letter, guessedLetters) {
		guessedLetters = append(guessedLetters, letter)
	} else if !isLetterAlreadyGuessed(letter, guessedLetters) {
		guessedLetters = append(guessedLetters, letter)
		attemptsLeft--
	}

	if attemptsLeft <= 0 || isWordGuessed(chosenWord, guessedLetters) {
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/hangman", http.StatusSeeOther)
}

func pickRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

func maskWord(word string, guessedLetters []string) string {
	var masked strings.Builder
	for _, char := range word {
		if containsString(guessedLetters, string(char)) {
			masked.WriteRune(char)
		} else {
			masked.WriteRune('_')
		}
		masked.WriteRune(' ')
	}
	return masked.String()
}

func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func isLetterInWord(letter string, word string) bool {
	return strings.Contains(word, letter)
}

func isLetterAlreadyGuessed(letter string, guessedLetters []string) bool {
	return containsString(guessedLetters, letter)
}

func isWordGuessed(word string, guessedLetters []string) bool {
	for _, char := range word {
		if !containsString(guessedLetters, string(char)) {
			return false
		}
	}
	return true
}

func resetGame() {
	chosenWord = pickRandomWord(words)
	guessedLetters = []string{}
	attemptsLeft = 10
}
