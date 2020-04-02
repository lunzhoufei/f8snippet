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
	"reflect"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// factoryCmd represents the factory command
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("factory called")
		color.Red("=> test1")
		test1()
		color.Red("=> test2")
		test2()

	},
}

func test1() {
	env1 := "production"
	env2 := "development"

	db1 := DatabaseFactory(env1)
	db2 := DatabaseFactory(env2)

	db1.(Database).PutData("test", "this is mongodb")
	fmt.Println(db1.(Database).GetData("test"))

	db2.(Database).PutData("test", "this is sqlite")
	fmt.Println(db2.(Database).GetData("test"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())
}

func SetupConstructors(env string) (Database, FileSystem) {
	fs := AbstractFactory("filesystem")
	db := AbstractFactory("database")
	return db(env).(Database), fs(env).(FileSystem)
}

func test2() {
	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstructors(env1)
	db2, fs2 := SetupConstructors(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fs1.CreateFile("../example/testntfs.txt")
	fmt.Println(fs1.FindFile("../example/testntfs.txt"))

	fs2.CreateFile("../example/testext4.txt")
	fmt.Println(fs2.FindFile("../example/testext4.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(fs1).Name())
	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(fs2).Name())
	fmt.Println(reflect.TypeOf(&fs2).Elem())
}

func init() {
	rootCmd.AddCommand(factoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ============================================================================
//  https://www.youtube.com/watch?v=-1xgg7yUlUc
// ============================================================================

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Factory func(string) interface{}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext ext4) CreateFile(path string) {
	file := file{content: "EXT4 file", name: path}
	ext.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}
	return ntfs.files[path]
}

func (ext ext4) FindFile(path string) file {
	if _, ok := ext.files[path]; !ok {
		return file{}
	}
	return ext.files[path]
}

// func FilesystemFactory(env string) FileSystem {
func FilesystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

// func DatabaseFactory(env string) Database {
func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

func AbstractFactory(fact string) func(string) interface{} {
	// func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil
	}
}
