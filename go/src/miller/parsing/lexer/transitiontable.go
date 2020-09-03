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
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 36: // ['$','$']
			return 4
		case r == 37: // ['%','%']
			return 5
		case r == 38: // ['&','&']
			return 6
		case r == 40: // ['(','(']
			return 7
		case r == 41: // [')',')']
			return 8
		case r == 42: // ['*','*']
			return 9
		case r == 43: // ['+','+']
			return 10
		case r == 45: // ['-','-']
			return 11
		case r == 46: // ['.','.']
			return 12
		case r == 47: // ['/','/']
			return 13
		case r == 48: // ['0','0']
			return 14
		case 49 <= r && r <= 57: // ['1','9']
			return 15
		case r == 58: // [':',':']
			return 16
		case r == 59: // [';',';']
			return 17
		case r == 60: // ['<','<']
			return 18
		case r == 61: // ['=','=']
			return 19
		case r == 62: // ['>','>']
			return 20
		case r == 63: // ['?','?']
			return 21
		case r == 70: // ['F','F']
			return 22
		case r == 73: // ['I','I']
			return 23
		case r == 78: // ['N','N']
			return 24
		case r == 79: // ['O','O']
			return 25
		case r == 91: // ['[','[']
			return 26
		case r == 93: // [']',']']
			return 27
		case r == 94: // ['^','^']
			return 28
		case r == 102: // ['f','f']
			return 29
		case r == 116: // ['t','t']
			return 30
		case r == 123: // ['{','{']
			return 31
		case r == 124: // ['|','|']
			return 32
		case r == 125: // ['}','}']
			return 33
		case r == 126: // ['~','~']
			return 34
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
		case r == 61: // ['=','=']
			return 35
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 36
		case r == 33: // ['!','!']
			return 36
		case r == 34: // ['"','"']
			return 37
		case r == 35: // ['#','#']
			return 36
		case r == 36: // ['$','$']
			return 36
		case r == 37: // ['%','%']
			return 36
		case r == 38: // ['&','&']
			return 36
		case r == 39: // [''',''']
			return 36
		case r == 40: // ['(','(']
			return 36
		case r == 41: // [')',')']
			return 36
		case r == 42: // ['*','*']
			return 36
		case r == 43: // ['+','+']
			return 36
		case r == 44: // [',',',']
			return 36
		case r == 45: // ['-','-']
			return 36
		case r == 46: // ['.','.']
			return 36
		case r == 47: // ['/','/']
			return 36
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 58: // [':',':']
			return 36
		case r == 59: // [';',';']
			return 36
		case r == 60: // ['<','<']
			return 36
		case r == 61: // ['=','=']
			return 36
		case r == 62: // ['>','>']
			return 36
		case r == 63: // ['?','?']
			return 36
		case r == 64: // ['@','@']
			return 36
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 91: // ['[','[']
			return 36
		case r == 92: // ['\','\']
			return 36
		case r == 93: // [']',']']
			return 36
		case r == 94: // ['^','^']
			return 36
		case r == 95: // ['_','_']
			return 36
		case r == 96: // ['`','`']
			return 36
		case 97 <= r && r <= 122: // ['a','z']
			return 36
		case r == 123: // ['{','{']
			return 36
		case r == 124: // ['|','|']
			return 36
		case r == 125: // ['}','}']
			return 36
		case r == 126: // ['~','~']
			return 36
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 36
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 38
		case 65 <= r && r <= 90: // ['A','Z']
			return 39
		case r == 91: // ['[','[']
			return 40
		case r == 95: // ['_','_']
			return 41
		case 97 <= r && r <= 122: // ['a','z']
			return 39
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case r == 37: // ['%','%']
			return 42
		case r == 61: // ['=','=']
			return 43
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 44
		case r == 61: // ['=','=']
			return 45
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 46
		case r == 61: // ['=','=']
			return 47
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 48
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 49
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 50
		case r == 43: // ['+','+']
			return 51
		case r == 45: // ['-','-']
			return 52
		case r == 47: // ['/','/']
			return 53
		case 48 <= r && r <= 57: // ['0','9']
			return 54
		case r == 61: // ['=','=']
			return 55
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 47: // ['/','/']
			return 56
		case r == 61: // ['=','=']
			return 57
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 58
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case r == 69: // ['E','E']
			return 59
		case r == 101: // ['e','e']
			return 59
		case r == 120: // ['x','x']
			return 60
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 58
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case r == 69: // ['E','E']
			return 59
		case r == 101: // ['e','e']
			return 59
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
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
		case r == 60: // ['<','<']
			return 61
		case r == 61: // ['=','=']
			return 62
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 63
		case r == 126: // ['~','~']
			return 64
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 65
		case r == 62: // ['>','>']
			return 66
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 73: // ['I','I']
			return 67
		case r == 78: // ['N','N']
			return 68
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 70: // ['F','F']
			return 69
		case r == 80: // ['P','P']
			return 70
		case r == 82: // ['R','R']
			return 71
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 70: // ['F','F']
			return 72
		case r == 82: // ['R','R']
			return 73
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 70: // ['F','F']
			return 74
		case r == 80: // ['P','P']
			return 75
		case r == 82: // ['R','R']
			return 76
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
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
			return 77
		case r == 94: // ['^','^']
			return 78
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 79
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 80
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 81
		case r == 124: // ['|','|']
			return 82
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
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 126: // ['~','~']
			return 83
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 36
		case r == 33: // ['!','!']
			return 36
		case r == 34: // ['"','"']
			return 37
		case r == 35: // ['#','#']
			return 36
		case r == 36: // ['$','$']
			return 36
		case r == 37: // ['%','%']
			return 36
		case r == 38: // ['&','&']
			return 36
		case r == 39: // [''',''']
			return 36
		case r == 40: // ['(','(']
			return 36
		case r == 41: // [')',')']
			return 36
		case r == 42: // ['*','*']
			return 36
		case r == 43: // ['+','+']
			return 36
		case r == 44: // [',',',']
			return 36
		case r == 45: // ['-','-']
			return 36
		case r == 46: // ['.','.']
			return 36
		case r == 47: // ['/','/']
			return 36
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 58: // [':',':']
			return 36
		case r == 59: // [';',';']
			return 36
		case r == 60: // ['<','<']
			return 36
		case r == 61: // ['=','=']
			return 36
		case r == 62: // ['>','>']
			return 36
		case r == 63: // ['?','?']
			return 36
		case r == 64: // ['@','@']
			return 36
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 91: // ['[','[']
			return 36
		case r == 92: // ['\','\']
			return 36
		case r == 93: // [']',']']
			return 36
		case r == 94: // ['^','^']
			return 36
		case r == 95: // ['_','_']
			return 36
		case r == 96: // ['`','`']
			return 36
		case 97 <= r && r <= 122: // ['a','z']
			return 36
		case r == 123: // ['{','{']
			return 36
		case r == 124: // ['|','|']
			return 36
		case r == 125: // ['}','}']
			return 36
		case r == 126: // ['~','~']
			return 36
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 36
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 38
		case 65 <= r && r <= 90: // ['A','Z']
			return 39
		case r == 95: // ['_','_']
			return 41
		case 97 <= r && r <= 122: // ['a','z']
			return 39
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 38
		case 65 <= r && r <= 90: // ['A','Z']
			return 39
		case r == 95: // ['_','_']
			return 41
		case 97 <= r && r <= 122: // ['a','z']
			return 39
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 38
		case 65 <= r && r <= 90: // ['A','Z']
			return 39
		case r == 95: // ['_','_']
			return 41
		case 97 <= r && r <= 122: // ['a','z']
			return 39
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 37: // ['%','%']
			return 84
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
			return 85
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
		case r == 61: // ['=','=']
			return 86
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
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 47: // ['/','/']
			return 87
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 54
		case r == 69: // ['E','E']
			return 88
		case r == 101: // ['e','e']
			return 88
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 89
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 90
		case r == 69: // ['E','E']
			return 91
		case r == 101: // ['e','e']
			return 91
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 92
		case 48 <= r && r <= 57: // ['0','9']
			return 93
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 94
		case 65 <= r && r <= 70: // ['A','F']
			return 94
		case 97 <= r && r <= 102: // ['a','f']
			return 94
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 95
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 96
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 76: // ['L','L']
			return 97
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 82: // ['R','R']
			return 98
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 99
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 100
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 101
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 102
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 103
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case r == 83: // ['S','S']
			return 104
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 105
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 106
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 107
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 108
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 109
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 110
		case 48 <= r && r <= 57: // ['0','9']
			return 111
		}
		return NoState
	},
	// S89
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S90
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 90
		case r == 69: // ['E','E']
			return 112
		case r == 101: // ['e','e']
			return 112
		}
		return NoState
	},
	// S91
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 113
		case 48 <= r && r <= 57: // ['0','9']
			return 114
		}
		return NoState
	},
	// S92
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 115
		}
		return NoState
	},
	// S93
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 93
		}
		return NoState
	},
	// S94
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 94
		case 65 <= r && r <= 70: // ['A','F']
			return 94
		case 97 <= r && r <= 102: // ['a','f']
			return 94
		}
		return NoState
	},
	// S95
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S96
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S97
	func(r rune) int {
		switch {
		case r == 69: // ['E','E']
			return 116
		}
		return NoState
	},
	// S98
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S99
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S100
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S101
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S102
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S103
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S104
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S105
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S106
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 117
		}
		return NoState
	},
	// S107
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 118
		}
		return NoState
	},
	// S108
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S109
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 119
		}
		return NoState
	},
	// S110
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 120
		}
		return NoState
	},
	// S111
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 111
		}
		return NoState
	},
	// S112
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 121
		case 48 <= r && r <= 57: // ['0','9']
			return 122
		}
		return NoState
	},
	// S113
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 123
		}
		return NoState
	},
	// S114
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 114
		}
		return NoState
	},
	// S115
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 115
		}
		return NoState
	},
	// S116
	func(r rune) int {
		switch {
		case r == 78: // ['N','N']
			return 124
		}
		return NoState
	},
	// S117
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 125
		}
		return NoState
	},
	// S118
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S119
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 126
		}
		return NoState
	},
	// S120
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 120
		}
		return NoState
	},
	// S121
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 127
		}
		return NoState
	},
	// S122
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 122
		}
		return NoState
	},
	// S123
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 123
		}
		return NoState
	},
	// S124
	func(r rune) int {
		switch {
		case r == 65: // ['A','A']
			return 128
		case r == 85: // ['U','U']
			return 129
		}
		return NoState
	},
	// S125
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S126
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 130
		}
		return NoState
	},
	// S127
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 127
		}
		return NoState
	},
	// S128
	func(r rune) int {
		switch {
		case r == 77: // ['M','M']
			return 131
		}
		return NoState
	},
	// S129
	func(r rune) int {
		switch {
		case r == 77: // ['M','M']
			return 132
		}
		return NoState
	},
	// S130
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 133
		}
		return NoState
	},
	// S131
	func(r rune) int {
		switch {
		case r == 69: // ['E','E']
			return 134
		}
		return NoState
	},
	// S132
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S133
	func(r rune) int {
		switch {
		case r == 37: // ['%','%']
			return 135
		}
		return NoState
	},
	// S134
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S135
	func(r rune) int {
		switch {
		case r == 37: // ['%','%']
			return 136
		}
		return NoState
	},
	// S136
	func(r rune) int {
		switch {
		case r == 37: // ['%','%']
			return 137
		}
		return NoState
	},
	// S137
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
