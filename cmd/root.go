package cmd

import (
	// "fmt"
	// "github.com/go-playground/locales/root"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	// "os"
)

var rootCmd = &cobra.Command{
	Use:   "calc cli",
	Short: "calculator cli in golnag",
	Long:  `Gopher CLI application in written in Go`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
