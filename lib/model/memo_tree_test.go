package model

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := `# [TOP] sample
## [WIP] sample2
`
	var expect *MemoNode = &MemoNode{
		Title:     "sample",
		Text:      "# [TOP] sample",
		Parent:    nil,
		Children:  nil,
		Attribute: "TOP",
		Level:     1,
	}
	var wip *MemoNode = &MemoNode{
		Title:     "sample2",
		Text:      "## [WIP] sample2",
		Parent:    expect,
		Children:  nil,
		Attribute: "WIP",
		Level:     2,
	}

	expect.Children = []*MemoNode{wip}

	ret, err := parse(data)
	if err != nil {
		t.Fatal(err)
	}
	if ret != expect {
		t.Fatalf("parse results missing expect %v but got %v\n", expect, ret)
	}
}

func TestGetNode(t *testing.T) {
}
