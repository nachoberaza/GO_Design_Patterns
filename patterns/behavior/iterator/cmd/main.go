package main

import (
	"fmt"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/iterator/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/iterator/internal/iterators"
)

func main() {
	//   1
	//  / \
	// 2   3

	// in-order:  213
	// preorder:  123
	// postorder: 231

	root := domain.NewNode(
		1,
		domain.NewTerminalNode(2),
		domain.NewTerminalNode(3),
	)

	inOrderIterator := iterators.NewInOrderIterator(root)
	for inOrderIterator.MoveNext() {
		fmt.Printf("%d,", inOrderIterator.GetCurrent().Value)
	}
	fmt.Println("\b")
}
