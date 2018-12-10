// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charliemaiors/playlist-zipper/archiver"
	"github.com/charliemaiors/playlist-zipper/playlist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "playlist-zipper",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Got error reading working directory %v", err)
			os.Exit(2)
		}
		fmt.Printf("Current dir is %s\n", curDir)
		produceArchives(curDir)
	},
}

var playlistType, archiveType string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(viper.AutomaticEnv)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&playlistType, "playlist-ext", "zpl", "Set playlist Type")
	rootCmd.PersistentFlags().StringVar(&archiveType, "archive-types", "zip", "Set archive Type")
}

func produceArchives(currentDir string) {
	playlistHanlder, err := playlist.NewReader(playlistType)
	if err != nil {
		fmt.Printf("Error defining playlist handler %v\n", err)
		os.Exit(3)
	}

	archiveHandler, err := archiver.NewArchiver(archiveType)
	if err != nil {
		fmt.Printf("Error defining archive handler %v\n", err)
		os.Exit(4)
	}

	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Printf("Error reading executable directory %v\n", err)
		os.Exit(5)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == "."+playlistType {
			fmt.Printf("File is %s", file.Name())
			playlistFiles, err := playlistHanlder.ReadPlaylist(file.Name())
			fmt.Printf("Files are %v", playlistFiles)
			if err != nil {
				fmt.Printf("Error reading playlist %s: %v", file.Name(), err)
				return
			}
			if err = archiveHandler.Archive(filepath.Base(file.Name()+".zip"), playlistFiles); err != nil {
				fmt.Printf("Error producing archive for playlist %s: %v", file.Name(), err)
				return
			}
			fmt.Printf("Done, created archive for %s", file.Name())
		}
	}
}
