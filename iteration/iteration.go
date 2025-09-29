package iteration

import "strings"

//Iterate the character based on the times arguments and return the updated string
func Iterate(character string,times int) string{
	var iteration strings.Builder
	for range times{
		iteration.WriteString(character)
	}

	return iteration.String()
}
