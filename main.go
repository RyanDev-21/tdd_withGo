package main

const( french = "French"
 spanish = "Spanish"
 FrenchHelloPrefix = "Bonjour , "
 EnglishHelloPrefix = "Hello , "
 SpanishHelloPrefix = "Hola , "
)
func Hello(name ,lang string) string{
	if name == ""{
		name = "World"	
	}

	return grettingPrefix(lang)+name
}

func grettingPrefix(lang string) string{
	prefix := EnglishHelloPrefix
	switch lang {
		case spanish:
			prefix = SpanishHelloPrefix
		case french:
			prefix = FrenchHelloPrefix
		

	}
	return prefix 
}

