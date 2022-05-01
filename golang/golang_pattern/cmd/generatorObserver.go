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
	// "fmt"

	// "github.com/spf13/cobra"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// generatorObserverCmd represents the generatorObserver command
var generatorObserverCmd = &cobra.Command{
	Use:   "generatorObserver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generatorObserver called")
		n := eventSubject{
			observers: sync.Map{},
		}

		var obs1 = eventObserver{id: 1, time: time.Now()}
		var obs2 = eventObserver{id: 2, time: time.Now()}
		n.AddListener(&obs1)
		n.AddListener(&obs2)

		go func() {
			select {
			case <-time.After(time.Millisecond * 10): // 定时任务
				n.RemoveListener(&obs1)
			}
		}()

		for x := range fib(10000000) {
			// color.Green("%d", x)
			n.Notify(Event{data: x})
		}
	},
}

func init() {
	rootCmd.AddCommand(generatorObserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generatorObserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorObserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ============================================================================
//  https://www.youtube.com/watch?v=Rb8s0zw1SNM
// ============================================================================

// generator
func fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

// event & observer
type (
	Event struct {
		data int
	}

	Observer interface {
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	eventObserver struct {
		id   int
		time time.Time
	}

	eventSubject struct {
		observers sync.Map
	}
)

func (e *eventObserver) NotifyCallback(event Event) {
	// fmt.Printf("observer: %d Received: %d after %v\n", e.id, event.data, time.Since(e.time))
	color.Red("observer: %d Received: %d after %v\n", e.id, event.data, time.Since(e.time))
}

func (s *eventSubject) AddListener(obs Observer) {
	s.observers.Store(obs, struct{}{})
}

func (s *eventSubject) RemoveListener(obs Observer) {
	s.observers.Delete(obs)
}

func (s *eventSubject) Notify(event Event) {
	s.observers.Range(func(key interface{}, val interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}
