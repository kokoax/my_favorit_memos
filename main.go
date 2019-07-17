package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

import "github.com/kokoax/my_favorit_memos/lib/manager"

func main() {
	filepath := "/home/kokoax/memos.md"
	f, _ := os.Open(filepath)
	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	title := flag.Bool("t", false, "Show list of title")
	extract := flag.String("e", "", "Extract memo and start edit")
	flag.Parse()

	mm, err := manager.NewMemoManager(string(b))
	if err != nil {
		panic(err)
	}

	if *title == true {
		fmt.Println(mm.TitleList())
	}

	if *extract != "" {
		fmt.Println(mm.GetNode(*extract))
	}

	// fmt.Println(mm.Decode(mm.TopLevel))
}
