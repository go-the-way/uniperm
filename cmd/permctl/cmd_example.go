// Copyright 2024 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"

	"fmt"

	"github.com/spf13/cobra"
)

var (
	exampleCmd = &cobra.Command{
		Use:     "example",
		Aliases: []string{"eg", "e"},
		Short:   "print the example",
		Long:    "print the example with format json",
		Run:     exampleRun,
	}
	//go:embed files/example.json
	exampleJsonFs string
)

func exampleRun(_ *cobra.Command, _ []string) {
	fmt.Println(exampleJsonFs)
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
