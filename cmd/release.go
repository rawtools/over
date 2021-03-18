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
	"github.com/spf13/cobra"
)

// ReleaseCmd represents the from command
var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release a file from the overlay",
	Long:  `Release will remove a file from the overlay, and inline its content if the overlay is applied somewhere`,
	Run: func(cmd *cobra.Command, args []string) {
		// user := user.Current()
		// for _, project := range user.Projects() {
		// 	fmt.Println(project.Name)
		// }
	},
}
