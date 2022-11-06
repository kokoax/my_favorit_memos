package model

import (
	"errors"
	"regexp"
)

// MemoTree is
type MemoTree struct {
}

// generateTree is
func generateTree(eachLine []string, parent *MemoNode, index int) (*MemoNode, error) {
	var child *MemoNode = &MemoNode{}
  var err error = nil

  if regexp.MustCompile(`^#`).Match([]byte(eachLine[index])) {
    child, err = getNode(eachLine[index], parent)
    if err != nil {
      return nil, err
    }

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

  return parent, nil
}

// getNode is
func getNode(line string, parent *MemoNode) (*MemoNode, error) {
  switch {
  case regexp.MustCompile(`^###`).Match([]byte(line)):
    title, attribute, err := getLevel3(line)
    if err != nil {
      return nil, err
    }

    return NewMemoNode(title, line, parent, nil, attribute, 3), nil
  case regexp.MustCompile(`^##`).Match([]byte(line)):
    title, attribute, err := getLevel2(line)
    if err != nil {
      return nil, err
    }

    return NewMemoNode(title, line, parent, nil, attribute, 2), nil
  case regexp.MustCompile(`^#`).Match([]byte(line)):
    title, attribute, err := getLevel1(line)
    if err != nil {
      return nil, err
    }

    return NewMemoNode(title, line, parent, nil, attribute, 1), nil
  }

  return &MemoNode{}, nil
}

// parse is
func parse(memo string) (*MemoNode, error) {
  eachLine := regexp.MustCompile(`\n`).Split(memo, -1)
  var parent *MemoNode
  if regexp.MustCompile(`^#\s+`).Match([]byte(eachLine[0])) {
    node, err := getNode(eachLine[0], nil)
    if err != nil {
      return nil, err
    }

    parent, err = generateTree(eachLine, node, 1)
    if err != nil {
      return nil, err
    }
  } else {
    return nil, errors.New("Missing Format: First must memo title")
  }

  return parent, nil
}
