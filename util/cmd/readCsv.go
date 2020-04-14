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
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// readCsvCmd represents the readCsv command
var readCsvCmd = &cobra.Command{
	Use:   "readCsv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readCsv called")
		err := doReadCsv()
		if err != nil {
			fmt.Printf("doReadCsv failed! err=%v", err)
		} else {
			fmt.Printf("doReadCsv succeed!")
		}
	},
}

// var (
// 	csvPath string
// 	outPath string
// )

func init() {
	csvCmd.AddCommand(readCsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCsvCmd.PersistentFlags().String("foo", "", "A help for foo")
	// csvCmd.PersistentFlags().StringVar(&csvPath, "path", "", "cvs file path")
	// csvCmd.PersistentFlags().StringVar(&outPath, "outpath", "", "output file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doReadCsv() error {

	//准备读取文件
	fmt.Println("path=", csvPath, " outpath=", outPath)
	fileName := csvPath
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
		fmt.Printf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
			fmt.Printf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}

	//针对小文件，也可以一次性读取所有的文件
	//注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
		fmt.Printf("can not readall, err is %+v", err)
	}
	for _, row := range content {
		fmt.Println(row)

	}

	//创建一个新文件
	newFileName := outPath
	//这样打开，每次都会清空文件内容
	//nfs, err := os.Create(newFileName)

	//这样可以追加写
	nfs, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
		fmt.Printf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)

	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	row := []string{"1", "2", "3", "4", "5,6"}
	err = w.Write(row)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
		fmt.Printf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()

	//一次写入多行
	var newContent [][]string
	newContent = append(newContent, []string{"1", "2", "3", "4", "5", "6"})
	newContent = append(newContent, []string{"11", "12", "13", "14", "15", "16"})
	newContent = append(newContent, []string{"21", "22", "23", "24", "25", "26"})
	w.WriteAll(newContent)
	return nil

}
