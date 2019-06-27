//package main
package binary_tree
import _"fmt"

type Node struct {
	val    int
	left    *Node
	right   *Node
}
// 定义完全二叉树
type BinaryTree struct {
	root *Node
}

//有序插入
func (tree *BinaryTree) Insert(v int) {
	if tree == nil{
		tree.root = &Node{val:v}
		return
	}
	cur := tree.root
	for {
		switch  {
		case v < cur.val:
			if cur.left==nil{
				cur.left = &Node{val:v}
				return
			}
			cur = cur.left
		case v > cur.val:
			if cur.right==nil{
				cur.right = &Node{val:v}
				return
			}
			cur = cur.right
		default:
			return
		}
	}

}

//前序遍历 parent->left->right
func (tree *BinaryTree) PreOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}
	f(node.val)
	if node.left!=nil{
		tree.PreOrderTraverse(node.left, f)
	}
	if node.right!=nil{
		tree.PreOrderTraverse(node.right, f)
	}

}

//中序遍历 left->parent->right
func (tree *BinaryTree) InOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}
	if node.left!=nil{
		tree.InOrderTraverse(node.left, f)
	}
	f(node.val)
	if node.right!=nil{
		tree.InOrderTraverse(node.right, f)
	}

}

//后序遍历  left->right->parent
func (tree *BinaryTree) PostOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}
	if node.left!=nil{
		tree.PostOrderTraverse(node.left, f)
	}
	if node.right!=nil{
		tree.PostOrderTraverse(node.right, f)
	}
	f(node.val)
}

// BFS
// 使用队列实现 BFS 广度(横向)优先遍历
func (tree *BinaryTree) BreadthFirstTraverse(f func(int)) {
	if tree.root==nil{
		return
	}
	q := make([]*Node,0)
	q = append(q,tree.root)
	for len(q)>0{
		node := q[0]
		q = q[1:]
		f(node.val)
		if node.left!=nil{
			q = append(q,node.left)
		}
		if node.right!=nil{
			q = append(q,node.right)
		}
	}
}

// 搜索
// 类似有序数组二分查找元素
func (tree *BinaryTree) Search(target int) *Node {
	if tree.root==nil{
		return nil
	}
	cur := tree.root
	for cur !=nil{
		switch  {
		case cur.val>target:
			cur = cur.left
		case cur.val<target:
			cur = cur.right
		default:
			return cur
		}
	}
	return nil
}

// 搜索 递归版本
func (tree *BinaryTree) CurSearch(node *Node, target int) *Node {
	if node == nil{
		return nil
	}
	switch {
	case node.val>target:
		return tree.CurSearch(node.left,target)
	case node.val<target:
		return tree.CurSearch(node.right,target)
	default:
		return node
	}
}

// 计算树的深度
// Node 的深度是 root 到 Node 的唯一路径长
func (tree *BinaryTree) Depth(node *Node) int {
	if node==nil{
		return 0
	}
	lDepth := tree.Depth(node.left)
	rDepth := tree.Depth(node.right)
	if lDepth > rDepth{
		return lDepth+1
	}
	return rDepth+1
}

// Node 节点数
func (tree *BinaryTree) NodeCount(node *Node) int {
	if node==nil{
		return 0
	}
	return 1+tree.NodeCount(node.left)+tree.NodeCount(node.right)
}

// Leaf 节点数
func (tree *BinaryTree) LeafCount(node *Node) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1
	}
	return tree.LeafCount(node.left) + tree.LeafCount(node.right)
}
//复制
func (tree *BinaryTree) Copy(node *Node) *Node{
	if node==nil{
		return nil
	}
	newNode := &Node{val:node.val}
	newNode.left = tree.Copy(node.left)
	newNode.right = tree.Copy(node.right)
	return newNode
}

// 镜像复制
func (tree *BinaryTree) Mirror(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{val: node.val}
	newNode.left = node.right
	newNode.right = node.left
	return newNode
}

// 判等
// 所有子节点的位置和值均相等的两棵树
func (tree *BinaryTree) Same(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.val == node2.val && tree.Same(node1.left, node2.left) && tree.Same(node1.right, node2.right)
}

// 遍历所有路径
// 找寻所有叶子节点的过程
func (tree *BinaryTree) AllPath(node *Node, stack []int, f func(int)){
	if node==nil{
		return
	}
	stack = append(stack,node.val)
	if node.left==nil && node.right==nil{
		f(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		return
	}
	tree.AllPath(node.left,stack,f)
	tree.AllPath(node.right,stack,f)
	stack = stack[:len(stack)-1]
}

// 删除整棵树
func (tree *BinaryTree) Free() {
	tree.root = nil // GC 会处理
}



func main(){
	arr := []int{4, 8, 2, 5, 7, 9, 1, 3, 10}
	var tree BinaryTree
	tree.root = &Node{val: 6}
	//tree.Insert(10)
	for _,i := range arr{
		tree.Insert(i)
	}
	tree.InOrderTraverse(tree.root,f)
}
