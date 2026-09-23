/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    dict := make(map[*Node]*Node)

	for cur := head; cur != nil; cur = cur.Next {
		dict[cur] = &Node{Val: cur.Val}
	}

	for cur := head; cur != nil; cur = cur.Next {
		dict[cur].Next = dict[cur.Next]
		dict[cur].Random = dict[cur.Random]
	}
	return dict[head]
}
