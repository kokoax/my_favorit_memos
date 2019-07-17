package model

import (
	"errors"
	"regexp"
)

// MemoTree is
type MemoTree struct {
}

// generateTree is
func generateTree(eachLine []string, parent *MemoNode, index int) *MemoNode {
	var child *MemoNode = &MemoNode{}

	if regexp.MustCompile(`^#`).Match([]byte(eachLine[index])) {
		child = getNode(eachLine[index], parent)

		// レベルが上がると親が変わるので親まで遡る
		if parent.Level >= child.Level {
			for parent.Level >= child.Level {
				parent = parent.Parent
			}
		}
		addChild(parent, child)
	} else {
		addText(parent, eachLine[index])
		child = parent
	}

	if index < len(eachLine)-1 {
		generateTree(eachLine, child, index+1)
	}

	return parent
}

// getNode is
func getNode(line string, parent *MemoNode) *MemoNode {
	switch {
	case regexp.MustCompile(`^###`).Match([]byte(line)):
		title, attribute := getLevel3(line)
		return NewMemoNode(title, line, parent, nil, attribute, 3)
	case regexp.MustCompile(`^##`).Match([]byte(line)):
		title, attribute := getLevel2(line)
		return NewMemoNode(title, line, parent, nil, attribute, 2)
	case regexp.MustCompile(`^#`).Match([]byte(line)):
		title, attribute := getLevel1(line)
		return NewMemoNode(title, line, parent, nil, attribute, 1)
	}

	return &MemoNode{}
}

// parse is
func parse(memo string) (*MemoNode, error) {
	eachLine := regexp.MustCompile(`\n`).Split(memo, -1)
	var parent *MemoNode
	if regexp.MustCompile(`^#\s+`).Match([]byte(eachLine[0])) {
		parent = generateTree(eachLine, getNode(eachLine[0], nil), 1)
	} else {
		return nil, errors.New("Missing Format: First must memo title")
	}

	return parent, nil
}
