type PrefixTree struct {
	wordMap map [string]bool
	nodes [26]*PrefixTree
	isEmpty bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		wordMap: make(map [string]bool),
	}
}

func (this *PrefixTree) Insert(word string) {
	this.wordMap[word] = true
	this.CreateTree(word)
}

func (this *PrefixTree) CreateTree(word string) {
	if len(word) == 0 {
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
	return this.wordMap[word]
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
