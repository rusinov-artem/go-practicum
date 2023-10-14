package daily

type WordsGenerator struct {
	alphabet []int	
	state []int
}

func (this *WordsGenerator) Next() {
	if len(this.alphabet)==0 {
		return
	}
	i := len(this.state)-1
	for g := true; g && i>=0; {
		this.state[i]++	
		if this.state[i] >= len(this.state) {
			this.state[i] = 0
			i--
		} else {
			g = false
		}
	}
}

func (this *WordsGenerator) Get() []int {
	if len(this.alphabet) == 0 {
		return []int{}
	}
	res := make([]int, len(this.alphabet))
	for i :=0 ; i < len(this.alphabet) ; i++ {
		res[i] = this.alphabet[this.state[i]]
	}
	return res
}

func NewWordsGenerator(inp []int) *WordsGenerator {
	state := make([]int, len(inp))
	return &WordsGenerator{
		alphabet : inp,
		state : state,
	}
	
}


type PermGenerator struct {
	set []int
	state []int
	subset [][]int
}

func (this *PermGenerator) Next() bool {
	i := len(this.state)-2
	for ; i >=0; i-- {
		if this.state[i] < len(this.subset[i])-1 {
			this.state[i]++
			this.generateSubset(i+1)
			break
		} else {
			this.state[i] = 0
		}
	}
	if i==-1 && this.state[0] ==0 {
		this.generateSubset(0)
		return false
	}
	return true
}

func (this *PermGenerator) generateSubset(depth int) {

	if len(this.set)< 2 {
		return 
	}

	if depth == 0 {
		this.subset[0] = this.set
		depth++
	}

	for i := depth ; i < len(this.set) ; i++ {
		rem := make([]int,0, len(this.subset[i-1])-1)
		for j := 0; j<len(this.subset[i-1]); j++ {
			if this.state[i-1] != j {
				rem = append(rem, this.subset[i-1][j])
			}
		} 
		this.subset[i] = rem
	}
}

func (this *PermGenerator) Get() []int {
	if len(this.set) == 0 {
		return []int{}
	}

	if len(this.set) == 1 {
		return []int{this.set[0]}
	}

	res := make([]int, len(this.state))
	for i := 0 ; i < len(this.state) ; i ++ {
		res[i] = this.subset[i][this.state[i]]		
	}
	return res
}

func NewPermGenerator(inp []int) *PermGenerator {
	state := make([]int, len(inp))
	p := &PermGenerator{
		set : inp,
		state : state,
		subset : make([][]int, len(inp)),
	}
	p.generateSubset(0)
	return p
}


func perm(inp []int) [][]int {
	if len(inp) == 0 {
		return [][]int{}
	}

	if len(inp) == 1 {
		return [][]int{{inp[0]}}
	}

	if len(inp) ==2 {
		return [][]int{
			{inp[0], inp[1]},
			{inp[1], inp[0]},
		}
	}

	res := [][]int{}
	for i := 0 ; i < len(inp) ; i++ {
		row := []int{}
		rem := []int{}
		for j:=0; j < len(inp); j++ {
			if j != i {
				rem = append(rem, inp[j])
			}
		}

		for _, v := range perm(rem) {
			row = append([]int{inp[i]}, v... )
			res = append(res, row)
		}

	}

	return res
}
