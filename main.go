package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	f := flag.Bool("leap_year", false, "leap year")
	flag.Parse()

	daylist := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if *f == true {
		daylist[1] = 29
	}
	fmt.Println(daylist)

	if err := os.Mkdir("DailyReports", 0666); err != nil {
		fmt.Println(err)
	}

	os.Chdir("DailyReports")

	for i := 1; i <= 12; i++ {
		subDir := strconv.Itoa(i)
		if err := os.MkdirAll(subDir, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for j := 1; j <= daylist[i-1]; j++ {
			str := "# 今日(" + strconv.Itoa(i) + "/" + strconv.Itoa(j) + ")の進捗\n"
			str = str + "## やったこと\n"
			str = str + "## わかったこと\n"
			str = str + "## わからなかったこと\n"
			str = str + "## 今後について\n"
			str = str + "## 感想\n"
			by := []byte(str)
			err := ioutil.WriteFile(subDir+"/"+strconv.Itoa(j)+".md", by, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
