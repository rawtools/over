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

	"github.com/sami2020pro/gmoji"
	"github.com/spf13/cobra"

	over "raw.tools/over/pkg"
	"raw.tools/over/pkg/styles"
)

// ApplyCmd represents the from command
var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply an overlay",
	Long:  `Apply an overlay to a target directory`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name, target := args[0], args[1]
		cfg, err := Parse()
		if err != nil {
			return err
		}
		msg := styles.White.Sprintf("Applying overlay %s", styles.WhiteItalic.Sprint(name))
		fmt.Printf("\U0001f4e6 %s\n", msg)

		repo := over.NewRepository(cfg)
		overlay := repo.Get(name)
		if overlay == nil {
			return fmt.Errorf("Unable to find overlay '%s'", name)
		}
		err = overlay.Apply(target)
		if err != nil {
			return fmt.Errorf("Unable to apply overlay %s to %s:\n%w", overlay.Name, target, err)
		}
		fmt.Println(
			gmoji.Check,
			styles.White.Sprint("Overlay"),
			styles.WhiteItalic.Sprint(name),
			styles.White.Sprint("applied with success"),
		)
		return nil
	},
}
