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
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// csv2mongoCmd represents the csv2mongo command
var csv2mongoCmd = &cobra.Command{
	Use:   "csv2mongo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("csv2mongo called")
		err := doCsv2mongo()
		if err != nil {
			fmt.Printf("doReadCsv failed! err=%v", err)
		} else {
			fmt.Printf("doReadCsv succeed!")
		}
	},
}

func init() {
	csvCmd.AddCommand(csv2mongoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csv2mongoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csv2mongoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doCsv2mongo() error {

	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		color.Red("mongo.NewClient failed! error=%v", err)
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		color.Red("check connect failed!")
		return err
	} else {
		color.Green("check connect ok!")
	}

	// collection
	collection := client.Database("feliz").Collection("abtest_metric")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

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
		line, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
			fmt.Printf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		// insert mongo
		row := strings.Split(line[0], "\t")

		if len(row) != 6 {
			color.Red("invalid row=%v, len=%d", row, len(row))
			break
		}
		metricTime, _ := strconv.ParseInt(row[0], 10, 64)
		metricPlatform, _ := strconv.ParseInt(row[1], 10, 64)
		metricExperiment, _ := strconv.ParseInt(row[2], 10, 64)
		metricStrategy, _ := strconv.ParseInt(row[3], 10, 64)

		item := bson.M{"time": metricTime,
			"platform":   metricPlatform,
			"experiment": metricExperiment,
			"strategy":   metricStrategy,
		}
		metrics := strings.Split(row[4], "&")
		for _, v := range metrics {
			metric := strings.Split(v, "=")
			if len(metric) != 2 {
				continue
			}
			item[metric[0]] = metric[1]
		}

		res, err := collection.InsertOne(ctx, item)
		if err == nil {
			color.Green("insert mongo succeed! id=%v", res.InsertedID)
		}
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
