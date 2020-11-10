package main

func main() {
	
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


type Node struct {
	Val int
	Children []*Node
}

// dfs
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}else if len(root.Children) == 0 {
		return 1
	}else{
		level := 0
		for _, item := range root.Children{
			level = max(level, maxDepth(item))
		}
		return level + 1
	}
}


func max(a, b int)int  {
	if a > b {
		return a
	}
	return b
}

// bfs
func maxDepthByBFS(root *Node)int  {
	if root == nil {
		return 0
	}
	var queue = []*Node{root}
	var level int
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length --
			for _, n := range queue[0].Children{
				queue = append(queue,n )
			}
			queue = queue[1:]
		}
		level++
	}
	return level
}