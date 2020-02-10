package go_print

import (
	"fmt"
)

func print_go() {
	var flagname string

	fmt.Println("name=", flagname)
	fmt.Printf("%+v\n", rootCmd.PersistentFlags())
}
