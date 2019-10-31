package mars_rover

type direction int
const (
	N  direction = 0 + iota
	E
	S
	W
)

var directions = [...] rune {
	'N',
	'E',
	'S',
	'W',
}



func (d direction) String() rune { return directions[d] }

func (d direction) int() int { return d }

func getDirConstant(d rune) direction {
	switch d {
	case 'N':
		return N
		break
	case 'S':
		return S
		break
	case 'W':
		return W
		break
	case 'E':
		return E
		break
	}
	return 0
}
