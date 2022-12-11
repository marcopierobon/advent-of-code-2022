package models

type Node struct {
	label    string
	children []*Node
	size     int
}
