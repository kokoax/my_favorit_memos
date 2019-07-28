package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

import "github.com/kokoax/my_favorit_memos/lib/manager"

func main() {
	filepath := "/Users/y-tokoi/memos/memo.md"
	f, _ := os.Open(filepath)
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	title := flag.Bool("t", false, "Show list of title")
	titleWithIndex := flag.Bool("ti", false, "Show list of title")
	extract := flag.String("e", "", "Extract memo and start edit")
	flag.Parse()

	mm, err := manager.NewMemoManager(string(b))
	if err != nil {
		panic(err)
	}

	if *title == true {
		fmt.Println(mm.TitleList())
	}

	if *titleWithIndex == true {
		fmt.Println(mm.TitleListWithIndex())
	}

	if *extract != "" {
		fmt.Println(mm.ExtractByIndex(0))
	}

	// fmt.Println(mm.Decode(mm.TopLevel))
}
