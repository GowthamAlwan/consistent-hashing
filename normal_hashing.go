package consistenthashing

import (
	"fmt"
)

var nodes []StorageNode = []StorageNode{
	{Name: 'A', Host: "10.131.213.12"},
	{Name: 'B', Host: "10.131.213.11"},
	{Name: 'C', Host: "10.131.213.10"},
	{Name: 'D', Host: "10.131.213.09"},
	{Name: 'E', Host: "10.131.213.08"},
}

func normalHash(key string) int {
	sum := 0
	for _, b := range []byte(key) {
		sum += int(b)
	}
	return sum % len(nodes)
}

func setVal(key, val string) {
	hash := normalHash(key)
	node := nodes[hash]
	node.set(key, val)
}

func getVal(key string) string {
	hash := normalHash(key)
	nodes := nodes[hash]
	val, err := nodes.get(key)
	if err != nil {
		fmt.Println("Error occured during getVal:", err)
	}
	return val
}
