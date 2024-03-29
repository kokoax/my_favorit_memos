package model

import (
  "testing"
)

func TestGetLevel3WORK(t *testing.T) {
  rtitle, rattribute, err := getLevel3("### [WORK] 2019/07/01")

  if err == nil {
    t.Fatalf("Execute faild getLevel3\n")
  }

  if rtitle != "2019/07/01" {
    t.Fatalf("getLevel3 title missing, expect 2019/07/01 but got '%s'\n", rtitle)
  }

  if rattribute != "WORK" {
    t.Fatalf("getLevel3 attribute missing expect 'WORK' but got '%s'\n", rattribute)
  }
}

func TestGetLevel3REFS(t *testing.T) {
  rtitle, rattribute, err := getLevel3("### [REFS] some title")

  if err == nil {
    t.Fatalf("Execute faild getLevel3\n")
  }

  if rtitle != "some title" {
    t.Fatalf("getLevel3 title missing, expect ''(empty) but got '%s'\n", rtitle)
  }

  if rattribute != "REFS" {
    t.Fatalf("getLevel3 attribute missing expect 'REFS' but got '%s'\n", rattribute)
  }
}

func TestGetLevel2REVIEW(t *testing.T) {
  rtitle, rattribute, err := getLevel2("## [REVIEW] sample title")

  if err == nil {
    t.Fatalf("Execute faild getLevel2\n")
  }

 if rtitle != "sample title" {
    t.Fatalf("getLevel2 title missing, expect 'sample title' but got '%s'\n", rtitle)
  }

  if rattribute != "REVIEW" {
    t.Fatalf("getLevel2 attribute missing expect 'REVIEW' but got '%s' \n", rattribute)
  }
}

func TestGetLevel1TOP(t *testing.T) {
  rtitle, rattribute, err := getLevel1("# [TOP] sample title")

  if err == nil {
    t.Fatalf("Execute faild getLevel1\n")
  }

  if rtitle != "sample title" {
    t.Fatalf("getLevel1 title missing, expect 'sample title' but got '%s'\n", rtitle)
  }

  if rattribute != "TOP" {
    t.Fatalf("getLevel1 attribute missing expect 'TOP' but got '%s'\n", rattribute)
  }
}
