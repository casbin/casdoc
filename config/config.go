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
