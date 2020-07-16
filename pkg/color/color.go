package color

import "fmt"

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// PrintError Message use red color
func PrintError(message string) {
	fmt.Println(Red + message + Reset)
}

// PrintWarn Message use yellow color
func PrintWarn(message string) {
	fmt.Println(Yellow + message + Reset)
}

// PrintSuccess Message use green color
func PrintSuccess(message string) {
	fmt.Println(Green + message + Reset)
}
