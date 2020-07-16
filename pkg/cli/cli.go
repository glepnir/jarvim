// Package cli provides ...
package common

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/spf13/cobra"
)

func NewCmd(use, description string, f func() error) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Long:  description,
		Short: description,
		RunE: func(cmd *cobra.Command, args []string) error {
			return f()
		},
	}
}

func MultiSelectTemplate(questionname, message string, options []string, pagesize int) ([]string, error) {
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
		os.Exit(1)
	}
	return answers, nil
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
		os.Exit(1)
	}
	return answer
}
