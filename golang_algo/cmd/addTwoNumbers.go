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

	"github.com/spf13/cobra"
)

// addTwoNumbersCmd represents the addTwoNumbers command
var addTwoNumbersCmd = &cobra.Command{
	Use:   "addTwoNumbers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addTwoNumbers called")
	},
}

func init() {
	listCmd.AddCommand(addTwoNumbersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTwoNumbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTwoNumbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ============================================================================
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// ============================================================================

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pos1 := l1
	pos2 := l2
	result := (*ListNode)(nil)
	tail := (*ListNode)(nil)

	more := 0
	for pos1 != nil || pos2 != nil {
		cur := &ListNode{}
		if pos1 != nil {
			cur.Val += pos1.Val
			pos1 = pos1.Next
		}
		if pos2 != nil {
			cur.Val += pos2.Val
			pos2 = pos2.Next
		}
		cur.Val += more
		more = cur.Val / 10
		cur.Val %= 10

		if result == nil {
			result = cur
			tail = cur
		} else {
			tail.Next = cur
			tail = tail.Next
		}
	}
	if more != 0 {
		cur := &ListNode{
			Val: more,
		}
		tail.Next = cur
		tail = tail.Next
	}

	return result
}
