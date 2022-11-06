package manager

import (
  "fmt"
  "strings"
)

import "github.com/kokoax/my_favorit_memos/lib/model"

// MemoManager is
type MemoManager struct {
  model	model.MemoModel
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
    model:	model.MemoModel{},
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

// ExtractFromIndex is
func (m MemoManager) ExtractByIndex(index int) string {
  return m.Decode(m.TopLevel.Children[index])
}

// TitleList is
func (m MemoManager) TitleList() string {
  titleList := []string{}
  for index := range m.TopLevel.Children {
    titleList = append(titleList, fmt.Sprintf("[%s] %s", m.TopLevel.Children[index].Attribute, m.TopLevel.Children[index].Title))
  }
  return strings.Join(titleList, "\n")
}

// TitleListWithIndex is
func (m MemoManager) TitleListWithIndex() string {
  titleList := []string{}
  for index := range m.TopLevel.Children {
    titleList = append(titleList, fmt.Sprintf("%3d [%s] %s", index, m.TopLevel.Children[index].Attribute, m.TopLevel.Children[index].Title))
  }
  return strings.Join(titleList, "\n")
}

func (m MemoManager) GetLevel3ByIndex(level2Index int) (string, error) {
  titleList := []string{}
  level3 := m.TopLevel.Children[level2Index].Children
  if len(level3) < level2Index {
    return "", fmt.Errorf(fmt.Sprintf("Dose not exist index %d", level2Index))
  }
  for index := range level3 {
    titleList = append(titleList, fmt.Sprintf("%3d [%s] %s", index, level3[index].Attribute, level3[index].Title))
  }
  return strings.Join(titleList, "\n"), nil
}

// Decode is
func (m MemoManager) Decode(parent *model.MemoNode) string {
  text := parent.Text
  for index := range parent.Children {
    text += "\n" + m.Decode(parent.Children[index])
  }

  return text
}

func (m MemoManager) GetLevel() (model.LevelYaml, error) {
  return model.GetLevelYaml()

}
