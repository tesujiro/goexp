package main

type binary struct {
	tables []table
}

/* Binary Search */
func newBinary(ts []table) binary {
	return binary{tables: ts}
}

func (b binary) find(r rune) bool {
	for _, t := range b.tables {
		if b.findTable(r, t) {
			return true
		}
		//fmt.Println("return false")
	}
	return false
}

func (b binary) findTable(r rune, t table) bool {
	// func (t table) IncludesRune(r rune) bool {
	if r < t[0].first {
		return false
	}

	bot := 0
	top := len(t) - 1
	for top >= bot {
		mid := (bot + top) / 2

		switch {
		case t[mid].last < r:
			bot = mid + 1
		case t[mid].first > r:
			top = mid - 1
		default:
			//fmt.Printf("first=%x last=%x\n", t[mid].first, t[mid].last)
			return true
		}
	}

	return false
}
