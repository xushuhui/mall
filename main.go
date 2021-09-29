package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mall_go/internal/model"
	"mall_go/internal/router"
	"mall_go/pkg/utils"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	app := router.InitRouter()
	log.Fatal(app.Listen(":3000"))

}
func test() {
	//准备读取文件

	fs, err := os.Open("./test/20210609sms.csv")
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件

	m := make(map[string]int, 0)
	tm := make(map[string][]string, 0)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		m[row[1]]++
		tm[row[1]] = append(tm[row[1]], row[8])

	}

	s := sortMapByValue(m)
	var data = make([][]string, 0)

	data = append(data, []string{"手机号", "次数", "userId", "发送时间"})

	for _, v := range s {
		if v.Value > 5 {
			time.Sleep(5 * time.Second)
			models, _ := model.GetUserByName(v.Key)
			timestr := strings.Join(tm[v.Key], ",")
			p := []string{v.Key, utils.IntToString(v.Value), utils.Int64ToString(models.ID), timestr}

			data = append(data, p)
		}
	}

	write("./test/20210609black.csv", data)

}

type Pair struct {
	Key   string
	Value int
}

func write(fileName string, data [][]string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
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
