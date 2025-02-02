package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Yosorable/gomic/initial"
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/route"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "project-test",
		Short: "Run web server",
		Long:  `A golang web server with frontend`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initial.SetLogrusAndGinFromConfigLogLevel()
			initial.RefreshTmpDB()

			r, err := route.CreateRoute()
			if err != nil {
				return err
			}

			return r.Run(global.CONFIG.Host + ":" + strconv.Itoa(global.CONFIG.Port))
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Run rootCmd error: %s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 8080)
	viper.SetDefault("data", "/data")
	viper.SetDefault("media", "/media")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		dir, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Cannot read config file: %v, using default settings\n", err)
	}

	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		fmt.Fprintln(os.Stderr, "Unmarshal config error:", err)
		os.Exit(1)
	} else {
		logrus.Infof("Config: %#v", global.CONFIG)
	}
}
