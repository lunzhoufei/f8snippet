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

	"sync"
	"time"

	"github.com/spf13/cobra"
)

// condCmd represents the cond command
var condCmd = &cobra.Command{
	Use:   "cond",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cond called")
		doTestCond()
	},
}

func init() {
	concurrencyCmd.AddCommand(condCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// condCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// condCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Golang的sync包中的Cond实现了一种条件变量，可以使用在多个Reader等待共享资源ready的场景（如果只有一读一写，一个锁或者channel就搞定了）。
// Cond的汇合点：多个goroutines等待、1个goroutine通知事件发生。

func doTestCond() {
	var locker = new(sync.Mutex)
	var cond = sync.NewCond(locker)

	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知，阻塞当前goroutine
			fmt.Println(x)
		}(i)
	}
	fmt.Println("睡眠1秒，使所有goroutine进入 Wait 阻塞状态")
	time.Sleep(time.Second * 1)

	fmt.Println("1秒后下发一个通知给已经获取锁的goroutine")
	time.Sleep(time.Second * 1)
	cond.Signal()

	fmt.Println("3秒后下发3个通知给已经获取锁的goroutine")
	time.Sleep(time.Second * 3)
	cond.Signal()
	cond.Signal()
	cond.Signal()

	fmt.Println("3秒后下发广播给所有等待的goroutine")
	time.Sleep(time.Second * 3)
	fmt.Println("Broadcast...")
	cond.Broadcast()

	fmt.Println("睡眠1秒，等待所有goroutine执行完毕")
	time.Sleep(time.Second * 1)
}
