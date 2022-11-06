package model

import (
  "testing"
)

func nodeEquality(obj1 *MemoNode , obj2 *MemoNode) bool {
  if obj1.Title != obj2.Title {
    return false
  }
  if obj1.Text != obj2.Text {
    return false
  }
  if obj1.Attribute != obj2.Attribute {
    return false
  }
  if obj1.Level != obj2.Level {
    return false
  }

  return true
}

func deepEquality(obj1 *MemoNode, obj2 *MemoNode) bool {
  if len(obj1.Children) != len(obj2.Children) {
    return false
  }

  if !nodeEquality(obj1, obj2) {
    return false
  }

  for index := range obj1.Children {
    if !nodeEquality(obj1.Children[index], obj2.Children[index]) {
      return false
    }
  }

  return true
}

func TestParse(t *testing.T) {
  data := `# [TOP] sample
  ## [WIP] sample2`
  var expect *MemoNode = &MemoNode{
    Title:		 "sample",
    Text:			"# [TOP] sample",
    Parent:		nil,
    Children:	nil,
    Attribute: "TOP",
    Level:		 1,
  }
  var wip *MemoNode = &MemoNode{
    Title:		 "sample2",
    Text:			"## [WIP] sample2",
    Parent:		expect,
    Children:	nil,
    Attribute: "WIP",
    Level:		 2,
  }

  expect.Children = []*MemoNode{wip}

  ret, err := parse(data)
  if err != nil {
    t.Fatal(err)
  }
  if !deepEquality(expect, ret) {
    t.Fatalf("parse results missing expect %#v but got %#v\n", expect, ret)
  }
}

func TestGetNode(t *testing.T) {
}
