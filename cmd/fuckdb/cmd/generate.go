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
	"os"

	"github.com/spf13/cobra"
)

var mysqlInfo = `{
  "db": {
    "host": "localhost",
    "port": 3306,
    "password": "password",
    "user": "root",
    "table": "table name",
    "database": "db_example",
    "packageName": "packageName",
    "structName": "structName",
    "structSorted": false,
    "jsonAnnotation": true,
    "dbAnnotation": false,
    "xmlAnnotation": false,
    "gormAnnotation": true
      }
}`
var fileName = "fuckdb.json"

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "use `fuckdb generate` to generate fuckdb.json",
	Long: `use fuckdb generate to generate fuckdb.json:
{
  "db": {
    "host": "localhost",
    "port": 3306,
    "password": "password",
    "user": "root",
    "table": "table name",
    "database": "db_example",
    "packageName": "packageName",
    "structName": "structName",
	"structSorted": false,
    "jsonAnnotation": true,
    "dbAnnotation": false,
    "xmlAnnotation": false,
    "gormAnnotation": true
      }
}`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("## fuckdb.json generated ##")
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write([]byte(mysqlInfo))
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
