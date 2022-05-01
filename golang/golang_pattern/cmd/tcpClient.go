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
	// "bytes"
	"fmt"
	"net"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// tcpClientCmd represents the tcpClient command
var tcpClientCmd = &cobra.Command{
	Use:   "tcpClient",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tcpClient called")
		e := startTcpClient()
		if e != nil {
			fmt.Printf("startTcpClient failed! error=%v", e)
		}
	},
}

var tcpClientParam struct {
	Port uint32
	Ip   string
}

func init() {
	netCmd.AddCommand(tcpServerCmd)
	tcpClientCmd.Flags().Uint32Var(&tcpClientParam.Port, "server-port", 19139, "Help message for toggle")
	tcpClientCmd.Flags().StringVar(&tcpClientParam.Ip, "server-ip", "localhost", "Help message for toggle")
}

func startTcpClient() error {

	remoteAddr := tcpClientParam.Ip + ":" + strconv.FormatUint(uint64(tcpClientParam.Port), 10)
	rAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		panic(err)
	}

	// localAddr is automatically chosen
	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		panic(err)
	}
	defer rConn.Close()

	rConn.Write([]byte("lunzhoufei"))
	rConn.Write([]byte("123"))
	rConn.Write([]byte("\r\n"))

	data := make([]byte, 1024)
	rConn.Read(data)
	color.Green("%s", string(data))

	return nil
}
