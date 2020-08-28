// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 3
		case r == 38: // ['&','&']
			return 4
		case r == 42: // ['*','*']
			return 5
		case r == 43: // ['+','+']
			return 6
		case r == 45: // ['-','-']
			return 7
		case r == 46: // ['.','.']
			return 8
		case r == 47: // ['/','/']
			return 9
		case r == 48: // ['0','0']
			return 10
		case 49 <= r && r <= 57: // ['1','9']
			return 11
		case r == 58: // [':',':']
			return 12
		case r == 59: // [';',';']
			return 13
		case r == 60: // ['<','<']
			return 14
		case r == 61: // ['=','=']
			return 15
		case r == 62: // ['>','>']
			return 16
		case r == 63: // ['?','?']
			return 17
		case r == 94: // ['^','^']
			return 18
		case r == 102: // ['f','f']
			return 19
		case r == 116: // ['t','t']
			return 20
		case r == 124: // ['|','|']
			return 21
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 25
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 26
		case r == 61: // ['=','=']
			return 27
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 28
		case r == 61: // ['=','=']
			return 29
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 30
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 31
		case r == 61: // ['=','=']
			return 32
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 33
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 47: // ['/','/']
			return 34
		case r == 61: // ['=','=']
			return 35
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		case r == 120: // ['x','x']
			return 36
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 60: // ['<','<']
			return 37
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 62: // ['>','>']
			return 38
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 39
		case r == 94: // ['^','^']
			return 40
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 41
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 42
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 43
		case r == 124: // ['|','|']
			return 44
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 45
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 46
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 31
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 47
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 48
		case 65 <= r && r <= 70: // ['A','F']
			return 48
		case 97 <= r && r <= 102: // ['a','f']
			return 48
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 49
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 50
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 51
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 52
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 53
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 54
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 48
		case 65 <= r && r <= 70: // ['A','F']
			return 48
		case 97 <= r && r <= 102: // ['a','f']
			return 48
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 55
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 56
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 57
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		}
		return NoState
	},
}