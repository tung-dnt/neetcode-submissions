/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func getOrCreate(dict map[*Node]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}
	if cp, ok := dict[node]; ok {
		return cp
	}
	cp := &Node{Val: node.Val}
	dict[node] = cp
	return cp
}

func copyRandomList(head *Node) *Node {
	dict := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		cp := getOrCreate(dict, cur)
		cp.Next = getOrCreate(dict, cur.Next)
		cp.Random = getOrCreate(dict, cur.Random)
	}
	return dict[head]
}