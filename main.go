package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
)

import "github.com/kokoax/my_favorit_memos/lib/manager"

func exec() error {
  filepath := "/Users/ca01072/memos/memo.md"
  f, _ := os.Open(filepath)
  defer f.Close()

  b, err := ioutil.ReadAll(f)
  if err != nil {
    return err
  }

  title := flag.Bool("t", false, "Show list of title")
  titleWithIndex := flag.Bool("ti", false, "Show list of title")
  extract := flag.String("e", "", "Extract memo and start edit")
  getLevel3ByIndex := flag.Int("g3", -1, "Show level3 by level2 index")
  level := flag.Bool("l", false, "Show list of level")
  flag.Parse()

  mm, err := manager.NewMemoManager(string(b))
  if err != nil {
    return err
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

  if *getLevel3ByIndex != -1 {
    ret, err := mm.GetLevel3ByIndex(*getLevel3ByIndex)
    if err != nil {
      return err
    }
    fmt.Println(ret)
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

  return nil
}

func main() {
  err := exec()
  if err != nil {
    fmt.Printf("%#v\n", err)
  }
}
