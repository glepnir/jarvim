// Package cli provides ...
package cli

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func ConfirmTemplate(message string) bool {
	confirm := false
	prompt := &survey.Confirm{
		Message: message,
	}
	err := survey.AskOne(prompt, &confirm)
	if err == terminal.InterruptErr {
		fmt.Println("interrupted")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
	return confirm
}

func MultiSelectTemplate(questionname, message string, options []string, pagesize int) []string {
	answers := []string{}
	var question = []*survey.Question{
		{
			Name: questionname,
			Prompt: &survey.MultiSelect{
				Message: message,
				Options: options,
			},
		},
	}
	err := survey.Ask(question, &answers, survey.WithIcons(func(icons *survey.IconSet) {
		icons.UnmarkedOption.Text = "○"
		icons.MarkedOption.Text = "◉"
	}), survey.WithPageSize(pagesize))
	if err == terminal.InterruptErr {
		fmt.Println("interrupted")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
	return answers
}

func SingleSelectTemplate(message string, options []string) string {
	var answer string
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	err := survey.AskOne(prompt, &answer)
	if err == terminal.InterruptErr {
		fmt.Println("interrupted")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
	return answer
}
