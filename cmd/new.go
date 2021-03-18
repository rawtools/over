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

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new overlay",
	// Long:  `List all known projects`,
	Run: func(cmd *cobra.Command, args []string) {
		// user := user.Current()
		// for _, project := range user.Projects() {
		// 	fmt.Println(project.Name)
		// }
	},
}
