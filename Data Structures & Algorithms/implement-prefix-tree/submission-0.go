type Node struct{
	links [26]*Node
	flag bool
}

func(n *Node)ContainsKey(ch byte)bool{
	return n.links[ch-'a'] != nil
}

func(n *Node)PutKey(ch byte, node *Node){
	n.links[ch-'a'] = node
}

func(n *Node)Get(ch byte) *Node{
	return n.links[ch-'a']
}

func(n *Node)SetEnd(){
	n.flag = true
}

func(n *Node)IsEnd()bool{
	return n.flag == true
}

func NewNode()*Node{
	return &Node{}
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{
		root : NewNode(),
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root

	for i := 0 ; i < len(word); i++{
		if !curr.ContainsKey(word[i]){
			node := NewNode()
			curr.PutKey(word[i], node)			
		}
		curr = curr.Get(word[i])
	}

	curr.SetEnd()
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root
	n := len(word)
	i := 0

	for i < n{
		if !curr.ContainsKey(word[i]){
			return false
		}

		curr = curr.Get(word[i])
		i++
	}

	return curr.IsEnd()
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root
	n := len(prefix)
	i := 0

	for i < n{
		if !curr.ContainsKey(prefix[i]){
			return false
		}

		curr = curr.Get(prefix[i])
		i++
	}

	return true
}
