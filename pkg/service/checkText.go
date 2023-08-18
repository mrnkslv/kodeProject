package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	note "github.com/mrnkslv/kodeProject"
)

func SpellCheck(text string) (string, error) {
	apiUrl := "https://speller.yandex.net/services/spellservice.json/checkTexts"

	input := note.TextsInput{
		Texts: []string{text},
		Lang:  "ru",
	}

	formValues, err := getFormValues(input)
	if err != nil {
		return "", err
	}

	response, err := http.PostForm(apiUrl, formValues)
	if err != nil {
		return "", err
	}

	var results []note.InputResult
	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", errors.New("no results returned")
	}
	correctedText := text
	for _, result := range results[0] {
		correctedText = strings.Replace(correctedText, result.Word, result.Suggestions[0], -1)
	}
	return correctedText, nil
}

func getFormValues(input note.TextsInput) (url.Values, error) {
	formValues := url.Values{}
	formValues.Set("text", strings.Join(input.Texts, "\n"))
	formValues.Set("lang", input.Lang)

	return formValues, nil
}
