package model

// MemoNode is
type MemoNode struct {
	Title     string
	Text      string
	Parent    *MemoNode
	Children  []*MemoNode
	Attribute string
	Level     int
}

// NewMemoNode is
func NewMemoNode(title string, text string, parent *MemoNode, children []*MemoNode, attribute string, level int) *MemoNode {
	return &MemoNode{
		Title:     title,
		Text:      text,
		Parent:    parent,
		Children:  children,
		Attribute: attribute,
		Level:     level,
	}
}

// AddChild is
func addChild(node *MemoNode, child *MemoNode) {
	node.Children = append(node.Children, child)
}

// AddText is
func addText(node *MemoNode, text string) {
	node.Text = node.Text + "\n" + text
}
