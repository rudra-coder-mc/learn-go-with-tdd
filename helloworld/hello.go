package helloworld

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(input string, lang string) string {
	if input == "" {
		return englishHelloPrefix + "world"
	}
	if lang == "" {
		return englishHelloPrefix + input
	}
	if lang == spanish {
		return spanishHelloPrefix + input
	}
	if lang == french {
		return frenchHelloPrefix + input
	}
	return englishHelloPrefix + input
}
