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
	"math/rand"
	"sort"
	"strconv"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// mapCmd represents the map command
var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("map called")
		doTestSyncMap()
	},
}

func init() {
	concurrencyCmd.AddCommand(mapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type GidInfo struct {
	id        int64
	name      string
	generator string
}

func doTestSyncMap() {
	gidlist := make([]int64, 1234)
	for i := 0; i < 1234; i++ {
		gidlist[i] = int64(i + 1)
	}

	// compose result in batch
	result := &sync.Map{}
	const batchCapacity = 99
	var batchSize int
	if len(gidlist)%batchCapacity == 0 {
		batchSize = len(gidlist) / batchCapacity
	} else {
		batchSize = len(gidlist)/batchCapacity + 1
	}

	wg := &sync.WaitGroup{}
	wg.Add(batchSize)

	composeGidInfo := func(generator int64, gids []int64) {
		defer wg.Done()
		for _, v := range gids {
			itemResult := &GidInfo{
				id:        v,
				generator: strconv.FormatInt(generator, 10),
				name:      RandStringBytes(12),
			}
			result.Store(v, itemResult)
		}
	}

	for k := 0; k < batchSize; k++ {
		if (k+1)*batchCapacity <= len(gidlist) {
			go composeGidInfo(int64(k), gidlist[k*batchCapacity:(k+1)*batchCapacity])
		} else {
			go composeGidInfo(int64(k), gidlist[k*batchCapacity:])
		}
	}
	wg.Wait()

	// Show result

	keys := make([]int64, 0, len(gidlist))
	vals := make([]*GidInfo, 0, len(gidlist))

	color.Green("======================= random order =======================")
	result.Range(func(key, value interface{}) bool {
		color.Green("[%04d] => %+v", key.(int64), *value.(*GidInfo))
		keys = append(keys, key.(int64))
		vals = append(vals, value.(*GidInfo))
		return true
	})

	color.Green("====================== key in order ========================")
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		item, _ := result.Load(k)
		color.Green("[%04d] => %+v", k, *item.(*GidInfo))
	}

	color.Green("===================== name in order ========================")
	// TODO
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].name < vals[j].name
	})
	for _, k := range keys {
		item, _ := result.Load(k)
		color.Green("[%04d] => %+v", k, *item.(*GidInfo))
	}

	color.Green("================= generatoor in order ======================")
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].generator < vals[j].generator
	})
	for _, k := range keys {
		item, _ := result.Load(k)
		color.Green("[%04d] => %+v", k, *item.(*GidInfo))
	}
}
