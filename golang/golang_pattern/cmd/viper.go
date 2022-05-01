/*
Copyright © 2020 lunzhoufei

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viperCmd represents the viper command
var viperCmd = &cobra.Command{
	Use:   "viper",
	Short: "开源配置管理库: https://github.com/spf13/viper",
	Long:  "开源配置管理库: https://github.com/spf13/viper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("viper called")
		doLoadPrintConfig()
	},
}

type MysqlConfig struct {
	Url      string
	Username string
	Password string
}

type Config struct {
	Port  int
	Mysql MysqlConfig
}

var viperPara struct {
	configName string
	configPath string
}

func init() {
	opensourceCmd.AddCommand(viperCmd)
	viperCmd.PersistentFlags().StringVar(&viperPara.configName, "configname",
		"feliz", "config file name")
	viperCmd.PersistentFlags().StringVar(&viperPara.configPath, "configpath",
		"./cnf", "config file path")
}

func doLoadPrintConfig() {
	var config Config
	viper.SetConfigName(viperPara.configName) // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(viperPara.configPath) // 第一个搜索路径
	// viper.SetConfigName("feliz") // 设置配置文件名 (不带后缀)
	// viper.AddConfigPath("cnf")   // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config) // 将配置信息绑定到结构体上
	fmt.Printf("%+v\n", config)
}

func doPrintConfig() {
	return
}
