package main

const spanish = "Spanish"
const french = "French"
const forBetterHelloEnglish = "Hello, "
const forBetterHelloSpanish = "Hola, "
const forBetterHelloFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "Asgardia"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefixOutput string) {
	switch language {
	case french:
		prefixOutput = forBetterHelloFrench
	case spanish:
		prefixOutput = forBetterHelloSpanish
	default:
		prefixOutput = forBetterHelloEnglish
	}

	return
}

func main() {
	//	fmt.Println(Hello("Asgardia"))
}
