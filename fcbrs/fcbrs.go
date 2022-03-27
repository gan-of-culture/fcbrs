package fcbrs

const tagetValue = 8
const amountMoves = 5

var moveNames = []string{
	"Investment",
	"Blackmail",
	"Donation",
	"Orgy",
	"Extra Time",
	"Teambuilding",
}

var moves = [][]int{
	{
		3, 0, -1,
	},
	{
		-3, -2, 3,
	},
	{
		2, 0, -2,
	},
	{
		-2, 2, 0,
	},
	{
		0, -3, 2,
	},
	{
		0, 3, -2,
	},
}

func nextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

func calcMove(vals []int, moveIdx int) []int {
	for idx := range vals {
		vals[idx] = vals[idx] + moves[moveIdx][idx]
		if vals[idx] < 1 {
			vals[idx] = 1
		}
		if vals[idx] > 16 {
			vals[idx] = 16
		}
	}
	return vals
}

func movesToInstructions(moves []int) []string {
	out := []string{}
	for _, move := range moves {
		out = append(out, moveNames[move])
	}
	return out
}

func Solve(a, b, c int) []string {
	// len(moves)^amountMoves
	numPossibleCombinations := func() int {
		base := len(moves)
		out := base
		for x := amountMoves - 1; x > 0; x-- {
			out *= base
		}
		return out
	}()
	np := nextProduct([]int{0, 1, 2, 3, 4, 5}, amountMoves)
	var tmpVals []int
	for i := 0; i < numPossibleCombinations; i++ {
		product := np()
		tmpVals = []int{a, b, c}
		for _, move := range product {
			tmpVals = calcMove(tmpVals, move)
		}
		if tmpVals[0] == tagetValue && tmpVals[1] == tagetValue && tmpVals[2] == tagetValue {
			return movesToInstructions(product)
		}
	}

	return nil
}
