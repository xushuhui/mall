package test

import (
	"encoding/csv"
	"fmt"
	"github.com/xushuhui/goal/utils"
	"io"
	"log"
	"os"
	"sort"
	"testing"
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

	//准备读取文件

	fs, err := os.Open("sms.csv")
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件

	m := make(map[string]int, 0)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		m[row[1]]++

	}
	s := sortMapByValue(m)
	var data = make([][]string, 0)
	p1 := []string{"手机号", "次数", "userId"}
	data = append(data, p1)

	for _, v := range s {
		if v.Value > 3 {
			p := []string{v.Key, utils.IntToString(v.Value)}
			data = append(data, p)
		}
	}

	write(data)

}

func write(data [][]string) {
	f, err := os.OpenFile("20210608black.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入一个UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}
