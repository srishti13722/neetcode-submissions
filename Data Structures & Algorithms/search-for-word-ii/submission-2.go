type Node struct{
	links [26]*Node
	flag bool
	word string
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

func(n *Node)SetWord(word string){
	n.word = word
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
	curr.SetWord(word)
}

func find(i, j int, board [][]byte, curr *Node, res *[]string){
    ch := board[i][j]

    if !curr.ContainsKey(ch) {
        return
    }

    curr = curr.Get(ch)

    if curr.word != "" {
        *res = append(*res, curr.word)
        curr.word = ""
    }

    board[i][j] = '#'
		
		dx := []int{-1, 0, 0 , 1}
		dy := []int{0, -1, 1 , 0}

		for k := 0 ; k < 4; k++{
			r := i + dx[k]
			c := j + dy[k]

			if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && board[r][c] != '#'{
				find(r,c,board, curr, res)
			}
		}

		board[i][j] = ch
	}


func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
    for _, word := range words{
		trie.Insert(word)
	}

	n := len(board)
	m := len(board[0])

	curr := trie.root

	var res []string

	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			find(i, j, board, curr, &res )
		}
	}

	return res
}
