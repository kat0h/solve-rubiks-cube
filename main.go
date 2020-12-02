package main

import "fmt"

type State struct {
	cp [8]int
	co [8]int
	ep [12]int
	eo [12]int
}

func newState(cp [8]int, co [8]int, ep [12]int, eo [12]int) *State {
	s := &State{
		cp: cp,
		co: co,
		ep: ep,
		eo: eo,
	}
	return s
}
func (self State) applyMove(move *State) *State {
	var new_cp [8]int
	for index, new_value := range move.cp {
		new_cp[index] = self.cp[new_value]
	}

	var new_co [8]int
	for index, new_value := range move.co {
		new_co[index] = (self.co[new_value] + move.co[index]) % 3
	}

	var new_ep [12]int
	for index, new_value := range move.ep {
		new_ep[index] = self.ep[new_value]
	}

	var new_eo [12]int
	for index, new_value := range move.eo {
		new_eo[index] = (self.eo[new_value]+move.eo[index]) % 2
	}

	return newState(new_cp, new_co, new_ep, new_eo)
}

func main() {
	r_state := newState(
		[8]int{0, 2, 6, 3, 4, 1, 5, 7},
		[8]int{0, 1, 2, 0, 0, 2, 1, 0},
		[12]int{0, 5, 9, 3, 4, 2, 6, 7, 8, 1, 10, 11},
		[12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)


	fmt.Print(r_state.applyMove(r_state))
}
