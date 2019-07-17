package model

// MemoModel is
type MemoModel struct {
}

// NewMemoModel is
func NewMemoModel() *MemoModel {
	return &MemoModel{}
}

// Parse is
func (m MemoModel) Parse(memo string) (*MemoNode, error) {
	return parse(memo)
}
