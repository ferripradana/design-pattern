package node

type INode interface {
	Print(indentation string)
	Clone() INode
}
