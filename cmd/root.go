/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/dinesh882002/rack/pkg/fileio"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
  fileType string
  isAllFiles *bool
  rootCmd = &cobra.Command{
    Use:   "rack",
    Short: "A command-line utility to keep your files organized",
    Long: `Rack is a command-line utility to keep your files organized.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    Args: cobra.MinimumNArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error { 
      return run(fileType, args[0], args[1])
    },
  }
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rack.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
  rootCmd.Flags().StringVarP(&fileType, "type", "t", "", "Specifies the type of files or file extensions")
  isAllFiles = rootCmd.Flags().BoolP("allfiles", "a", false, "Organize all files into its own directories")
  rootCmd.MarkFlagsOneRequired("type", "allfiles")
  rootCmd.MarkFlagsMutuallyExclusive("type", "allfiles")
  // rootCmd.MarkFlagRequired("type")
}

func run(fileType, sourceDir string, target string) error {
  // Check if the source directory exists
  if err := fileio.IsDirExists(sourceDir); err != nil {
    return err
  }

  if *isAllFiles {
    return fileio.OrganizeFiles(sourceDir, target)
  }

  return fileio.CopyFiles(sourceDir, target, fileType)
}
