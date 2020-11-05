package hash

import (
	"fmt"
	"testing"
)

func TestHashRing(t *testing.T) {
	hashRing := NewHashRing()
	for i := 1; i <= 30*12; i++ {
		hashRing.Add(fmt.Sprintf("node%d", i))
	}

	str1, str2, str3, str4 := "str1", "str2", "str3", "str4"

	node1, _ := hashRing.Get(str1)
	node2, _ := hashRing.Get(str2)
	node3, _ := hashRing.Get(str3)
	node4, _ := hashRing.Get(str4)
	fmt.Println("str1:", node1, ",str2", node2, ",str3", node3, ",str4", node4)

	for i := 1; i <= 30; i++ {
		hashRing.Add(fmt.Sprintf("node%d", i))
	}
	fmt.Println("============================")

	node1, _ = hashRing.Get(str1)
	node2, _ = hashRing.Get(str2)
	node3, _ = hashRing.Get(str3)
	node4, _ = hashRing.Get(str4)
	fmt.Println("str1:", node1, ",str2", node2, ",str3", node3, ",str4", node4)
}
