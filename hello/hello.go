package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const thai = "Thai"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const thaiPrefix = "萨瓦迪卡, "

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case thai:
		prefix = thaiPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world",""))
}