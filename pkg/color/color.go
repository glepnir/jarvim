package color

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// Print Error Message use red color
func PrintError(message string) {
	fmt.Println(Red + message + Reset)
}

// Print Warn Message use yellow color
func PrintWarn(message string) {
	fmt.Println(Yellow + message + Reset)
}

// Print Success Message use green color
func PrintSuccess(message string) {
	fmt.Println(Green + message + Reset)
}
