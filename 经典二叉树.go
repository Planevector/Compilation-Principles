package main

import "fmt"

// TreeNode 结构体定义了一个二叉搜索树的节点，包含一个整数值 Val 和两个指向子节点的指针 Left 和 Right
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree 表示二叉树
type Tree struct {
	// 根节点
	Root *TreeNode
}

// Insert 函数用于向二叉搜索树中插入一个新节点
func (bt *Tree) Insert(val int) {
	// 如果二叉搜索树的根节点为空，则创建一个新的根节点并插入
	if bt.Root == nil {
		bt.Root = &TreeNode{Val: val}
	} else {
		// 否则，将新节点插入到二叉搜索树中
		bt.Root.insert(val)
	}
}

// insert 函数用于插入一个新节点到二叉搜索树中
func (node *TreeNode) insert(val int) {
	// 如果新节点的值小于当前节点的值
	if val < node.Val {
		// 如果当前节点的左子节点为空，则创建一个新的左子节点并插入
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			// 否则，将新节点插入到左子树中
			node.Left.insert(val)
		}
	} else {
		// 如果新节点的值大于等于当前节点的值
		if node.Right == nil {
			// 如果当前节点的右子节点为空，则创建一个新的右子节点并插入
			node.Right = &TreeNode{Val: val}
		} else {
			// 否则，将新节点插入到右子树中
			node.Right.insert(val)
		}
	}
}

// 前序遍历（Preorder Traversal）：遍历顺序为【根节点 -> 左子树 -> 右子树】。
// 如果需要复制整个二叉树，那么前序遍历可能是最高效的，因为先复制根节点，然后递归复制左右子树，整个过程更加直观。
// 前序遍历二叉树
// PreorderTraversal 前序遍历二叉树的入口函数
func (t *Tree) PreorderTraversal() []int {
	// 如果二叉树的根节点为空，直接返回空数组
	if t.Root == nil {
		return nil
	}
	// 调用根节点的前序遍历函数
	return t.Root.preorderTraversal()
}

// preorderTraversal 前序遍历二叉树的递归函数
func (t *TreeNode) preorderTraversal() []int {
	// 如果当前节点为空，直接返回空数组
	if t == nil {
		return nil
	}
	// 定义一个数组来保存遍历结果
	result := []int{}
	// 将当前节点的值添加到遍历结果中
	result = append(result, t.Val)
	// 递归遍历当前节点的右子树，并将结果添加到遍历结果中
	result = append(result, t.Right.preorderTraversal()...)
	// 递归遍历当前节点的左子树，并将结果添加到遍历结果中
	result = append(result, t.Left.preorderTraversal()...)
	return result
}

// 中序遍历（Inorder Traversal）：遍历顺序为【左子树 -> 根节点 -> 右子树】。
// 如果需要寻找二叉搜索树中的最小值或最大值，中序遍历是最高效的，因为二叉搜索树的中序遍历结果是有序的。

// 中序遍历二叉树
// InorderTraversal 中序遍历二叉树的入口函数
func (t *Tree) InorderTraversal() []int {
	// 如果二叉树的根节点为空，直接返回空数组
	if t.Root == nil {
		return nil
	}
	// 调用根节点的中序遍历函数
	return t.Root.inorderTraversal()
}

// inorderTraversal 中序遍历二叉树的递归函数
func (t *TreeNode) inorderTraversal() []int {
	// 如果当前节点为空，直接返回空数组
	if t == nil {
		return nil
	}

	// 定义一个数组来保存遍历结果
	result := []int{}

	// 递归遍历当前节点的左子树，并将结果添加到遍历结果中
	result = append(result, t.Left.inorderTraversal()...)

	// 将当前节点的值添加到遍历结果中
	result = append(result, t.Val)

	// 递归遍历当前节点的右子树，并将结果添加到遍历结果中
	result = append(result, t.Right.inorderTraversal()...)

	return result
}

// 后序遍历（Postorder Traversal）：遍历顺序为【左子树 -> 右子树 -> 根节点】。
// 如果需要释放整个二叉树的内存，后序遍历可能是最高效的，因为先释放左右子树的内存，最后再释放根节点的内存，可以避免重复操作。
// 后序遍历二叉树
// PostorderTraversal 后序遍历二叉树的入口函数
func (t *Tree) PostorderTraversal() []int {
	// 如果二叉树的根节点为空，直接返回空数组
	if t.Root == nil {
		return nil
	}
	// 调用根节点的后序遍历函数
	return t.Root.postorderTraversal()
}

// postorderTraversal 后序遍历二叉树的递归函数
func (t *TreeNode) postorderTraversal() []int {
	// 如果当前节点为空，直接返回空数组
	if t == nil {
		return nil
	}

	// 定义一个数组来保存遍历结果
	result := []int{}

	// 递归遍历当前节点的左子树，并将结果添加到遍历结果中
	result = append(result, t.Left.postorderTraversal()...)

	// 递归遍历当前节点的右子树，并将结果添加到遍历结果中
	result = append(result, t.Right.postorderTraversal()...)

	// 将当前节点的值添加到遍历结果中
	result = append(result, t.Val)

	return result
}

func main() {
	bt := &Tree{}
	bt.Insert(5)
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(8)
	bt.Insert(3)
	bt.Insert(9)

	// 前序遍历二叉树 //迭代方式
	fmt.Println("前序遍历：", bt.PreorderTraversal())

	// 中序遍历二叉树 //迭代方式
	fmt.Println("中序遍历：", bt.InorderTraversal())

	// 后序遍历以及以下 //迭代方式
	fmt.Println("后序遍历：", bt.PostorderTraversal())

}
