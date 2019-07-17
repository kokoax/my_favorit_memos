package manager

import (
	"strings"
)

import "github.com/kokoax/my_favorit_memos/lib/model"

// MemoManager is
type MemoManager struct {
	model    model.MemoModel
	TopLevel *model.MemoNode
}

// MemoManagerInterface is
type MemoManagerInterface interface {
	NewMemoManager(memo string) (*MemoManager, error)
	Level1() *model.MemoNode
	Level2() *model.MemoNode
	Level3() *model.MemoNode
	Decode(parent *model.MemoNode) string
}

// NewMemoManager is
func NewMemoManager(memo string) (*MemoManager, error) {
	topLevel, err := model.MemoModel{}.Parse(memo)
	if err != nil {
		return nil, err
	}

	return &MemoManager{
		model:    model.MemoModel{},
		TopLevel: topLevel,
	}, nil
}

// Level1 is
func (m MemoManager) Level1() *model.MemoNode {
	return m.TopLevel
}

// Level2 is
func (m MemoManager) Level2() []*model.MemoNode {
	return m.TopLevel.Children
}

// Level3 is
func (m MemoManager) Level3() []*model.MemoNode {
	var ret []*model.MemoNode
	for index := range m.TopLevel.Children {
		ret = append(ret, m.TopLevel.Children[index])
	}

	return ret
}

// GetNode is
func (m MemoManager) GetNode(title string) string {
	for index := range m.TopLevel.Children {
		if m.TopLevel.Children[index].Title == title {
			return m.Decode(m.TopLevel.Children[index])
		}
	}
	return ""
}

// TitleList is
func (m MemoManager) TitleList() string {
	titleList := []string{}
	for index := range m.TopLevel.Children {
		titleList = append(titleList, m.TopLevel.Children[index].Title)
	}
	return strings.Join(titleList, "\n")
}

// Decode is
func (m MemoManager) Decode(parent *model.MemoNode) string {
	text := parent.Text
	for index := range parent.Children {
		text += "\n" + m.Decode(parent.Children[index])
	}

	return text
}
