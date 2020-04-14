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
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)

// mongoQueryCmd represents the mongoQuery command
var mongoQueryCmd = &cobra.Command{
	Use:   "mongoQuery",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mongoQuery called")
		err := doMongoQuery()
		if err != nil {
			fmt.Printf("doMongoQuery failed! err=%v", err)
		} else {
			fmt.Printf("doMongoQuery succeed!")
		}
	},
}

func init() {
	mongoCmd.AddCommand(mongoQueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mongoQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mongoQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ============================================================================
// SEE: https://www.cnblogs.com/zcqkk/p/11234227.html
// ============================================================================

// type MetricItem struct {
// 	time       int64
// 	platform   int64
// 	experiment int64
// 	strategy   int64
// 	pv1        string
// 	pv2        string
// 	uv1        string
// 	uv2        string
// 	ctr        string
// }

type MetricItem map[string]interface{}

func doMongoQuery() error {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), opts)
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

	// 获取数据总数
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	color.Green("collection.CountDocuments:", count)

	// 查询单条数据
	var one MetricItem
	err = collection.FindOne(context.Background(),
		bson.M{"time": 201905141935, "strategy": 220004}).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	color.Green("collection.FindOne: ", one)

	// 查询多条数据(方式一)
	// cur, err := collection.Find(context.Background(), bson.D{"strategy": 220004})
	cur, err := collection.Find(context.Background(), bson.M{"strategy": 220004})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*MetricItem
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	for _, one := range all {
		log.Println(one)
	}

	// 查询多条数据(方式二)
	cur, err = collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		var b MetricItem
		if err = cur.Decode(&b); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find cur.Next:", b)
	}
	cur.Close(context.Background())

	// 模糊查询
	cur, err = collection.Find(context.Background(), bson.M{"name": primitive.Regex{Pattern: "深入"}})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		var b MetricItem
		if err = cur.Decode(&b); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{深入}: ", b)
	}
	cur.Close(context.Background())

	// 二级结构体查询
	cur, err = collection.Find(context.Background(), bson.M{"author.country": "china"})
	// cur, err = collection.Find(context.Background(), bson.D{bson.E{"author.country", countryChina}})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		var b MetricItem
		if err = cur.Decode(&b); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find author.country=", "china", ":", b)
	}
	cur.Close(context.Background())

	// 修改一条数据
	// b1 := books[0].(*Book)
	// b1.Weight = 2
	// update := bson.M{"$set": b1}
	// updateResult, err := collection.UpdateOne(context.Background(), bson.M{"name": b1.Name}, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("collection.UpdateOne:", updateResult)

	// 修改一条数据，如果不存在则插入
	// new := &Book{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "球状闪电",
	// 	Category: categorySciFi,
	// 	Author: AuthorInfo{
	// 		Name:    "刘慈欣",
	// 		Country: countryChina,
	// 	},
	// }
	// update = bson.M{"$set": new}
	// updateOpts := options.Update().SetUpsert(true)
	// updateResult, err = collection.UpdateOne(context.Background(), bson.M{"_id": new.Id}, update, updateOpts)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("collection.UpdateOne:", updateResult)

	// 删除一条数据
	newid := primitive.NewObjectID()
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": newid})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)

	return nil
}
