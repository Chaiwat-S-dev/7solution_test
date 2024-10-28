package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type SummaryResponse struct {
	Beef map[string]int `json:"beef"`
}

func FetchAPI() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func CountMeats(text string) map[string]int {
	re := regexp.MustCompile(`\b[\w-]+`)
	words := re.FindAllString(strings.ToLower(text), -1)
	meatCount := map[string]int{
		"t-bone":   0,
		"fatback":  0,
		"pastrami": 0,
		"pork":     0,
		"meatloaf": 0,
		"jowl":     0,
		"enim":     0,
		"bresaola": 0,
	}

	for _, word := range words {
		if _, ok := meatCount[word]; ok {
			meatCount[word]++
		}
	}
	return meatCount
}

func BeefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	text, err := FetchAPI()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	meatCount := CountMeats(text)
	response := SummaryResponse{Beef: meatCount}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/beef/summary", BeefSummaryHandler)

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
