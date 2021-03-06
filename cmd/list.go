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

	"github.com/spf13/cobra"

	over "raw.tools/over/pkg"
)

// ListCmd represents the from command
var ListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List known projects",
	Long:    `List all known projects`,
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := Parse()
		if err != nil {
			return err
		}

		repo := over.NewRepository(cfg)
		for _, overlay := range repo.List() {
			fmt.Println(overlay)
			err := overlay.ParseConfig()
			if err != nil {
				return fmt.Errorf("unable to load overlay %s:\n%w", overlay, err)
			}
		}

		return nil
	},
}
