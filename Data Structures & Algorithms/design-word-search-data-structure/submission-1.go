type WordDictionary struct {
    root *Node
}

type Node struct{
	nodes [26]*Node
	endOfTheWord bool
}

func Constructor() WordDictionary {
    return WordDictionary{
		root:&Node{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    this.root.AddWord(word)
}

func (this *Node) AddWord (word string){
	if len(word) == 0 {
		this.endOfTheWord = true
		return
	}
	char := word[0]
	word = word[1:]
	index := int(char)-97

	if this.nodes[index] == nil{
		this.nodes[index] = &Node{}
	}
	this.nodes[index].AddWord(word)
}

func (this *WordDictionary) Search(word string) bool {
    return this.root.Search(word)
}

func (this *Node) Search (word string) bool{
	// if len(word) == 2 && word[0] == '.' && word[1] == '.'{
	// 	return true
	// }else if len(word) == 1 && word[0] == '.'{
	// 	return true
	// }else 
	if len(word) == 0 {
		return this.endOfTheWord
	}

	char := word[0]
	word = word[1:]
	index := int(char)-97
	if char != '.'{
		if this.nodes[index] == nil{
			return false
		}
		return this.nodes[index].Search(word)
	}else {
		for _, node := range this.nodes{
			if node == nil{
				continue
			}
			found := node.Search(word)
			if found {
				return true
			}
		}
		return false
	}
}
