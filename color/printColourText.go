package color

import "fmt"

// Foreground color, background color, action 具体值参阅readme
func BuildColorText(text string, fg, bg, act int8) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1b, act, bg, fg, text, 0x1b)
}
