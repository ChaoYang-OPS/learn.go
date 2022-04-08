package main

type fatRateSuggestion struct {
}

func (c fatRateSuggestion)GetSuggestion(person *Person) string {
	return "todo"
}
