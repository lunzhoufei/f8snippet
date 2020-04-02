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
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// decoratorCmd represents the decorator command
var decoratorCmd = &cobra.Command{
	Use:   "decorator",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decorator called")
		// fmt.Plintln(Pi(1000))
		// fmt.Plintln(Pi(50000))
		fpcache := wrapcache(Pi, &sync.Map{})
		fplog := wraplogger(fpcache, log.New(os.Stdout, "pi", 1))
		fplog(100000)
		fplog(2000)
		fplog(100000)

		fdcache := wrapcache(Divide, &sync.Map{})
		fdlog := wraplogger(fdcache, log.New(os.Stdout, "divide", 1))
		fdlog(100000)
		fdlog(2000)
		fdlog(100000)

	},
}

func init() {
	rootCmd.AddCommand(decoratorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decoratorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decoratorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ============================================================================
//  https://www.youtube.com/watch?v=F365lY5ECGY&list=PLJbE2Yu2zumAKLbWO3E2vKXDlQ8LT_R28
// ============================================================================

type piFunc func(int) float64

func wraplogger(fun piFunc, logger *log.Logger) piFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

func wrapcache(fun piFunc, cache *sync.Map) piFunc {
	return func(n int) float64 {
		fn := func(n int) float64 {
			key := fmt.Sprintf("n=%d", n)
			val, ok := cache.Load(key)
			if ok {
				return val.(float64)
			}
			result := fun(n)
			cache.Store(key, result)
			return result
		}
		return fn(n)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}
	result := 0.0
	for k := 0; k <= n; k++ {
		result += <-ch
	}
	return result
}

func Divide(n int) float64 {
	return float64(n / 32)
}
