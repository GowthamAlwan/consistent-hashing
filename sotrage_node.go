package consistenthashing

import (
	"errors"
	"fmt"
)

type StorageNode struct {
	Name rune
	Host string
}

var store map[string]string
var ErrKeyNotFound = errors.New("key not found")

func (s *StorageNode) set(key, val string) {
	fmt.Printf("set called on Node: %+v for Key: %s\n", s, key)
	store[key] = val
}

func (s *StorageNode) get(key string) (string, error) {
	fmt.Printf("set called on Node: %+v for Key: %s\n", s, key)
	val, ok := store[key]
	if !ok {
		fmt.Println(ErrKeyNotFound)
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (s *StorageNode) getAllKeys() []string {
	allKeys := []string{}
	for k := range store {
		allKeys = append(allKeys, k)
	}
	return allKeys
}

func (s *StorageNode) String() string {
	return fmt.Sprintf("{Name: %c, Host: %s}", s.Name, s.Host)
}
