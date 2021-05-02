/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	fuckdb_go "fuckdb/cmd/fuckdb/fuckdbGo"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "fuckdb go to generate golang struct with gorm and json tag",
	Run: func(cmd *cobra.Command, args []string) {
		res := fuckdb_go.FuckdbGo()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
