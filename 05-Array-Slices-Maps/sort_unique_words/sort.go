package sortuniquewords

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SortUnique() {
	words := make(map[string]int)

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		val := scan.Text()
		words[val]++
	}

	fmt.Printf("unique words:%d\n%+v\n", len(words), words)

	type kv struct {
		key string
		val int
	}

	var sortedSlice []kv

	for k, v := range words {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	sort.Slice(sortedSlice, func(a int, b int) bool {
		return sortedSlice[a].val > sortedSlice[b].val
	})

	for _, item := range sortedSlice[:5] {
		fmt.Printf("%s appears %d times.\n", item.key, item.val)
	}
}
