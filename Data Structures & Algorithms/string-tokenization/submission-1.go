type Node struct {
	children map[rune]*Node
	id       int
	isEnd    bool
}

func NewNode(id int) *Node {
	return &Node{children: make(map[rune]*Node), id: id}
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: NewNode(0)}
}

func (t *Trie) Insert(word string, id int) {
	node := t.root

	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewNode(id)
		}

		node = node.children[char]
	}

	node.isEnd = true
}

func tokenize(text string, dictionary map[string]int) []string {
	output := []string{}
	trie := InitTrie()

	for k, i := range dictionary {
		trie.Insert(k, i)
	}

	pos := 0
	for pos < len(text) {
		node := trie.root
		longestMatchLen := 0
		matchedId := ""

		for i := pos; i < len(text); i++ {
			c := rune(text[i])

			nextNode, exists := node.children[c]

			if !exists {
				break
			}
			node = nextNode

			if node.isEnd {
				longestMatchLen = (i - pos) + 1
				matchedId = strconv.Itoa(node.id)
			}
		}

		if longestMatchLen > 0 {
			output = append(output, matchedId)
			pos += longestMatchLen
		} else {
			output = append(output, string(text[pos]))
			pos++
		}
	}

	return output
}
