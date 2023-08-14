// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	AuthToken string
	RPM       int
	TPM       int
	RepoPath  string
)

func init() {
	viper.SetDefault("AuthToken", "")
	viper.SetDefault("RPM", 3)
	viper.SetDefault("TPM", 40000)
	viper.SetDefault("RepoPath", "")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	_ = viper.ReadInConfig()

	viper.SetEnvPrefix("CASDOC")
	viper.AutomaticEnv()

	RepoPath = viper.GetString("RepoPath")
	AuthToken = viper.GetString("AuthToken")
	RPM = viper.GetInt("RPM")
	TPM = viper.GetInt("TPM")

	if AuthToken == "" {
		fmt.Println("Not found AuthToken, please type it: ")
		var input string
		_, _ = fmt.Scanln(&input)
		AuthToken = input
		if AuthToken == "" {
			panic("AuthToken is empty")
		}
	}

	if RepoPath == "" {
		fmt.Println("Not found RepoPath, please type it: ")
		var input string
		_, _ = fmt.Scanln(&input)
		RepoPath = input
		if RepoPath == "" {
			panic("RepoPath is empty")
		}
	}

	fmt.Println("RepoPath: ", RepoPath)
	fmt.Println("RPM: ", RPM)
	fmt.Println("TPM: ", TPM)
}
