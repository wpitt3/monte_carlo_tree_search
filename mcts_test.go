package mcts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchTree() SearchTree {
	return NewSearchTree(RandomRollout{}, float32(1.414), FlatPolicy{})
}

func Test_rootnode(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	rootNode := createNode(board)

	assert.Equal(t, nil, rootNode.action)
	assert.Equal(t, 9, len(rootNode.unexpandedActions))
	assert.Equal(t, false, rootNode.isTerminal)
	assert.Equal(t, 0, rootNode.winner)
}

func Test_newNode(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	rootNode := createNode(board)
	newNode := newNode(rootNode, TicTacToeAction{0, 0})

	assert.Equal(t, TicTacToeAction{0, 0}, newNode.action)
	assert.Equal(t, 8, len(newNode.unexpandedActions))
	assert.Equal(t, false, newNode.isTerminal)
	assert.Equal(t, 0, newNode.winner)
	assert.Equal(t, -1, newNode.board.(TicTacToe).board[0][0])
	assert.Equal(t, 9, len(rootNode.unexpandedActions))
	assert.Equal(t, rootNode.children[0], newNode)
}

func Test_newNodeHasWinner(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	board = board.PerformMove(TicTacToeAction{0, 0})
	board = board.PerformMove(TicTacToeAction{0, 1})
	board = board.PerformMove(TicTacToeAction{1, 0})
	board = board.PerformMove(TicTacToeAction{0, 2})
	rootNode := createNode(board)
	newNode := newNode(rootNode, TicTacToeAction{2, 0})

	assert.Equal(t, TicTacToeAction{2, 0}, newNode.action)
	assert.Equal(t, -1, newNode.winner)
	assert.Equal(t, true, newNode.isTerminal)
}

func Test_selectLeafNode_selectFirstNode(t *testing.T) {
	rootNode := createNode(TicTacToe{})
	action1 := rootNode.unexpandedActions[0]

	node := searchTree().selectLeafNode(rootNode)
	assert.Equal(t, node.action, action1)
}

func Test_selectLeafNode_expandChildIfNoUnexpandedActions(t *testing.T) {
	rootNode := createNode(TicTacToe{})
	rootNode.unexpandedActions = rootNode.unexpandedActions[:1]
	searchTree := searchTree()

	childNode := searchTree.selectLeafNode(rootNode)
	grandChildNode := searchTree.selectLeafNode(rootNode)

	assert.NotEqual(t, childNode, grandChildNode)
	assert.Equal(t, childNode.children[0], grandChildNode)
	assert.Equal(t, 0, len(rootNode.unexpandedActions))
	assert.Equal(t, 7, len(childNode.unexpandedActions))
	assert.Equal(t, 7, len(grandChildNode.unexpandedActions))
}

func Test_selectLeafNode_doNotExpandTerminal(t *testing.T) {
	rootNode := createNode(TicTacToe{})
	rootNode.isTerminal = true
	searchTree := searchTree()

	childNode := searchTree.selectLeafNode(rootNode)

	assert.Equal(t, childNode, rootNode)
}

func Test_findBestChild(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	rootNode := createNode(board)
	nodeA := newNode(rootNode, TicTacToeAction{1, 0})
	nodeB := newNode(rootNode, TicTacToeAction{2, 0})
	searchTree := searchTree()
	rootNode.numer = 1.0
	rootNode.denom = 4.0
	nodeA.numer = 1.0
	nodeA.denom = 2.0
	nodeB.numer = 2.0
	nodeB.denom = 2.0

	// B is better move and should be explored more
	assert.Equal(t, nodeB, searchTree.findBestChild(rootNode))

	// B is better move still, but a should be explored more as b is overly explored
	rootNode.denom = 5.0
	nodeB.denom = 3.0
	assert.Equal(t, nodeA, searchTree.findBestChild(rootNode))
}

func Test_findBestChildByTurns_win(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	board = board.PerformMove(TicTacToeAction{0, 0})
	board = board.PerformMove(TicTacToeAction{0, 1})
	board = board.PerformMove(TicTacToeAction{1, 0})
	board = board.PerformMove(TicTacToeAction{0, 2})

	searchTree := searchTree()
	expectedAction := TicTacToeAction{2, 0}

	assert.Equal(t, expectedAction, searchTree.FindBestMoveByTurns(board, 50))
	assert.Equal(t, expectedAction, searchTree.FindBestMoveByTime(board, 10))
}

func Test_findBestChildByTurns_stopLoss(t *testing.T) {
	var board State[Action] = TicTacToe{[3][3]int{}}
	board = board.PerformMove(TicTacToeAction{0, 0})
	board = board.PerformMove(TicTacToeAction{1, 1})
	board = board.PerformMove(TicTacToeAction{1, 0})
	board = board.PerformMove(TicTacToeAction{2, 0})

	searchTree := searchTree()
	expectedAction := TicTacToeAction{0, 2}

	assert.Equal(t, expectedAction, searchTree.FindBestMoveByTurns(board, 50))
	assert.Equal(t, expectedAction, searchTree.FindBestMoveByTime(board, 10))
}
