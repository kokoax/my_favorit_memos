package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
)

import "github.com/kokoax/my_favorit_memos/lib/manager"

func main() {
  filepath := "/Users/ca01072/memos/memo.md"
  f, _ := os.Open(filepath)
  defer f.Close()

  b, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }

  title := flag.Bool("t", false, "Show list of title")
  titleWithIndex := flag.Bool("ti", false, "Show list of title")
  extract := flag.String("e", "", "Extract memo and start edit")
  level := flag.Bool("l", false, "Show list of level")
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

  if *level == true {
    level, _ := mm.GetLevel()

    fmt.Println("Level1")
    for index := range level.Level1 {
      fmt.Println("\t", level.Level1[index])
    }
    fmt.Println("Level2")
    for index := range level.Level2 {
      fmt.Println("\t", level.Level2[index])
    }
    fmt.Println("Level3")
    for index := range level.Level3 {
      fmt.Println("\t", level.Level3[index])
    }
  }
}
