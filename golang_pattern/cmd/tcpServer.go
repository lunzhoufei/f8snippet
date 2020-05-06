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
	"bytes"
	"fmt"
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

// tcpServerCmd represents the tcpServer command
var tcpServerCmd = &cobra.Command{
	Use:   "tcpServer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tcpServer called")
		e := startTcpServer()
		if e != nil {
			fmt.Printf("startTcpServer failed! error=%v", e)
		}
	},
}

var tcpServerParam struct {
	Port uint32
}

func init() {
	netCmd.AddCommand(tcpServerCmd)
	tcpServerCmd.Flags().Uint32Var(&tcpServerParam.Port, "toggle", 19139, "Help message for toggle")
}

func handleConn(in <-chan *net.TCPConn, out chan<- *net.TCPConn) {
	for conn := range in {
		// proxyConn(conn)
		buf := &bytes.Buffer{}
		for {
			data := make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil {
				panic(err)
			}
			// \r\n
			if data[0] == 13 && data[1] == 10 {
				break
			}
			buf.Write(data[:n])
		}
		data := buf.Bytes()
		fmt.Printf("received: %v", data)

		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}

		_, err := conn.Write(data)
		if err != nil {
			fmt.Printf("conn.Write failed! error=%v", err)
		}

		out <- conn
	}
}

func closeConn(in <-chan *net.TCPConn) {
	for conn := range in {
		conn.Close()
	}
}

func startTcpServer() error {

	localAddr := "localhost:" + strconv.FormatUint(uint64(tcpServerParam.Port), 10)
	fmt.Printf("Listening: %v", localAddr)

	addr, err := net.ResolveTCPAddr("tcp", localAddr)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	pending, complete := make(chan *net.TCPConn), make(chan *net.TCPConn)

	for i := 0; i < 5; i++ {
		go handleConn(pending, complete)
	}
	go closeConn(complete)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		pending <- conn
	}

	return nil
}
