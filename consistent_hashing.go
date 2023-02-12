package consistenthashing

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

var totalSlots int = 50
var activeNodes []StorageNode
var nodeKeys []int

func consistentHash(key string) int {
	h := sha256.New()
	h.Write([]byte(key))
	sum := 0
	for _, b := range h.Sum(nil) {
		sum += int(b)
	}
	return sum % totalSlots
}

func set(key, val string) {
	h := consistentHash(key)
	index := findNextGreaterPos(nodeKeys, h)
	node := activeNodes[index]
	node.set(key, val)
}

func get(key string) string {
	h := consistentHash(key)
	index := findNextGreaterPos(nodeKeys, h)
	node := activeNodes[index]
	val, err := node.get(key)
	if err != nil {
		fmt.Println("Error occured during get: ", err)
	}
	return val
}

func addNode(node StorageNode) {
	hk := consistentHash(node.Host)
	index := findNextGreaterPos(nodeKeys, hk)
	// todo: perform data migration
	nodeKeys = insert(nodeKeys, hk, index)
	activeNodes = insert(activeNodes, node, index)
}

func insert[T any](keys []T, key T, index int) []T {
	if len(keys) == index {
		keys = append(keys, key)
	} else if index == 0 {
		keys = append([]T{key}, keys...)
	} else {
		temp := append(keys[:index], key)
		keys = append(temp, keys[index:]...)
	}
	return keys
}

func reHashNode(node StorageNode) {
	allKeys := node.getAllKeys()
	for _, k := range allKeys {
		val, err := node.get(k)
		if err != nil {
			fmt.Println("Error during reHashNode: ", err)
		}
		set(k, val)
	}
}

func findNextGreaterPos(arr []int, val int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] >= val })
}
