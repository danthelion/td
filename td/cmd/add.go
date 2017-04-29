// Copyright © 2017 Daniel Palma danivgy@gmail.com
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
	"github.com/danthelion/todo"
	"fmt"
	"log"
	"github.com/spf13/viper"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry to your to-do list.",
	Run:   addRun,
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1, 2, 3")

}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
