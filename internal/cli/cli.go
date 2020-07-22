// Copyright 2020 The jarvim Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/glepnir/jarvim/internal/logic"
	"github.com/spf13/cobra"
)

var (
	cli_version = "0.1.6"
	version     bool
	genConfig   bool
)

var rootCmd = &cobra.Command{
	Use:   "jarvim",
	Short: "jarvim is a cli tool to generate a module vim configruation which like a pro",
	Long:  "jarvim is a cli tool to generate a module vim configruation which like a pro",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		if genConfig {
			return logic.RunLogic()
		}
		return cmd.Help()
	},
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&version, "Version", "v", false, "show current version of CLI")
	rootCmd.Flags().BoolVarP(&genConfig, "Generate vim config", "g", false, "generate new configuration")
}

// Execute do the rootmcmd.Execute() function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func printVersion() error {
	fmt.Println("Version: ", cli_version)
	return nil
}
