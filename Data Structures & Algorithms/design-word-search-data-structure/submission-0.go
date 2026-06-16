type Node struct{
	links [26]*Node
	flag bool
}

func(n *Node)ContainsKey(ch byte)bool{
	return n.links[ch -'a'] != nil
}

func(n *Node)Put(ch byte , node *Node){
	n.links[ch - 'a'] = node
}

func(n *Node)Get(ch byte)*Node{
	return n.links[ch -'a']
}

func(n *Node)IsEnd()bool{
	return n.flag == true
}

func(n *Node)SetEnd(){
	n.flag = true
}

type WordDictionary struct {
    root *Node
}

func Constructor() WordDictionary {
    return WordDictionary{
		root : &Node{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root

	for i :=0 ; i < len(word); i++{
		if !curr.ContainsKey(word[i]){
			curr.Put(word[i], &Node{})
		}
		curr = curr.Get(word[i])
	}

	curr.SetEnd()
}

func searchUtil(curr *Node , word string) bool {
	if curr == nil {
        return false
    }
	for i := 0 ; i < len(word); i++{
		if word[i] == '.'{
			for j := byte('a') ; j <= byte('z') ; j++{
				if searchUtil(curr.Get(j), word[i+1:]){
					return true
				}
			}

			return false
		}else if !curr.ContainsKey(word[i]){
			return false
		}

		curr = curr.Get(word[i])
	}

	return curr.IsEnd()
}

func (this *WordDictionary) Search(word string) bool {
	return searchUtil(this.root, word)
}
