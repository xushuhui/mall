package test

import (
	"sort"
	"testing"
	"time"
)

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))
	return p
}
func TestFunc(t *testing.T) {
	ts, _ := time.Parse("2006-01-02", "2021-10-11")

	for i := 0; i < 24; i++ {
		s := ts
		ts = ts.Add(1 * time.Hour)
		e := ts
		t.Log(s, e)
	}
}

//func write(data [][]string) {
//	f, err := s.OpenFile("20210608black.csv", os.O_RDWR|os.O_CREATE, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//
//	f.WriteString("\xEF\xBB\xBF") // 写入一个UTF-8 BOM
//
//	w := csv.NewWriter(f) //创建一个新的写入文件流
//	w.WriteAll(data)
//	w.Flush()
//}
