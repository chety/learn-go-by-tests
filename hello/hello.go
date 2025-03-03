package hello

import "fmt"

type Language string
type Greeting string

const (
	English = "English"
	Spanish = "Spanish"
	French  = "French"
)

const (
	EnglishGreeting = "Hello,"
	SpanishGreeting = "Hola,"
	FrenchGreeting  = "Bonjour,"
)

var languageGreetingStore = map[Language]Greeting{
	English: EnglishGreeting,
	Spanish: SpanishGreeting,
	French:  FrenchGreeting,
}

func getGreeting(language Language) Greeting {
	if value, exists := languageGreetingStore[language]; exists {
		return value
	}
	return EnglishGreeting
}

func Hello(name string, language Language) string {
	if len(name) == 0 {
		name = "World!"
	}
	return fmt.Sprintf("%s %s", getGreeting(Language(language)), name)
}
