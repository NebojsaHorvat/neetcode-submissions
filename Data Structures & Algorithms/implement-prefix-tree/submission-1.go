type PrefixTree struct {
	wordMap map [string]bool
	nodes [26]*PrefixTree
	endOfWord bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	this.CreateTree(word)
}

func (this *PrefixTree) CreateTree(word string) {
	if len(word) == 0 {
		this.endOfWord = true
		return
	}
	index := int(word[0])-97
	word = word[1:]
	if this.nodes[index] == nil{
		this.nodes[index] = &PrefixTree{}
		this.nodes[index].CreateTree(word)

	}else{
		this.nodes[index].CreateTree(word)
	}
}


func (this *PrefixTree) Search(word string) bool {
	if len(word) == 0{
		return this.endOfWord
	}
	index := int(word[0])-97
	word = word[1:]
	if this.nodes[index] == nil{
		return false
	}
	return this.nodes[index].Search(word)
}

func (this *PrefixTree) StartsWith(word string) bool {
	if len(word) == 0{
		return true
	}
	index := int(word[0])-97
	word = word[1:]
	if this.nodes[index] == nil{
		return false
	}
	return this.nodes[index].StartsWith(word)
}
