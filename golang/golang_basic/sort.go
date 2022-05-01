package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 3 types sort api

// ============================================================================
// 1. builtin api
// ============================================================================

// sort.Ints([]int)
// sort.Floats([]float)
// sort.Strings([]string)

func type1sort() {
	s := []int{4, 3, 2, 1}
	sort.Ints(s)
	fmt.Println(s) // 1,2,3,4
}

// ============================================================================
// 2. cusmize compare
// ============================================================================

// sort.Slice([]interface{}, less func(int, int) bool)
// sort.SliceStable([]interface{}, less func(int, int) bool)

type Type2 struct {
	name string
	age  int
}

func type2sort() {
	s := []Type2{Type2{"fei", 1}, Type2{"fe2", 3}, Type2{"k", 18}}

	// stable
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].age > s[j]
	})

	//
	sort.Slice(s, func(i, j int) bool {
		return s[i].age > s[j]
	})

	fmt.Println(s) // 1,2,3,4
}

// ============================================================================
// 3. cusmize compare
// ============================================================================

type Interface interface {
	Swap(i, j int)
	Len() int
	Less(i, j int) bool
}

// SortSongByPlayDays 根据播放天数对歌曲进行排序
type SortedSongList []*annualpb.SongListInfo

func (p SortedSongList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p SortedSongList) Len() int           { return len(p) }
func (p SortedSongList) Less(i, j int) bool { return p[i].GetPlayDays() > p[j].GetPlayDays() }

func type3sort() {
	rawSlice := make([]*annualpb.SongListInfo)
	sort.Sort(SortedSongList(rawSlice))
}
