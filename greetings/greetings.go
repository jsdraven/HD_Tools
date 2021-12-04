package greetings

import "fmt"

// Hello returns a greeting for the given name.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	// Return a greeting that embeds the name in a message.
	return fmt.Sprintf("Hello, %s!", name)

}
