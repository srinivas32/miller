// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

// CharLiteralValue returns the string value of the char literal.
func (t *Token) CharLiteralValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

// Float32Value returns the float32 value of the token or an error if the token literal does not
// denote a valid float32.
func (t *Token) Float32Value() (float32, error) {
	if v, err := strconv.ParseFloat(string(t.Lit), 32); err != nil {
		return 0, err
	} else {
		return float32(v), nil
	}
}

// Float64Value returns the float64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Float64Value() (float64, error) {
	return strconv.ParseFloat(string(t.Lit), 64)
}

// IDValue returns the string representation of an identifier token.
func (t *Token) IDValue() string {
	return string(t.Lit)
}

// Int32Value returns the int32 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int32Value() (int32, error) {
	if v, err := strconv.ParseInt(string(t.Lit), 10, 64); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

// Int64Value returns the int64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int64Value() (int64, error) {
	return strconv.ParseInt(string(t.Lit), 10, 64)
}

// UTF8Rune decodes the UTF8 rune in the token literal. It returns utf8.RuneError if
// the token literal contains an invalid rune.
func (t *Token) UTF8Rune() (rune, error) {
	r, _ := utf8.DecodeRune(t.Lit)
	if r == utf8.RuneError {
		err := fmt.Errorf("Invalid rune")
		return r, err
	}
	return r, nil
}

// StringValue returns the string value of the token literal.
func (t *Token) StringValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		";",
		"{",
		"}",
		"=",
		"md_token_filter",
		"md_token_emit",
		"md_token_dump",
		"md_token_edump",
		"md_token_field_name",
		"$[",
		"]",
		"md_token_braced_field_name",
		"md_token_full_srec",
		"md_token_oosvar_name",
		"@[",
		"md_token_braced_oosvar_name",
		"md_token_full_oosvar",
		"md_token_non_sigil_name",
		"||=",
		"^^=",
		"&&=",
		"|=",
		"^=",
		"<<=",
		">>=",
		">>>=",
		"+=",
		".=",
		"-=",
		"*=",
		"/=",
		"//=",
		"%=",
		"**=",
		"?",
		":",
		"||",
		"^^",
		"&&",
		"=~",
		"!=~",
		"==",
		"!=",
		">",
		">=",
		"<",
		"<=",
		"|",
		"^",
		"&",
		"<<",
		">>",
		">>>",
		"+",
		"-",
		".+",
		".-",
		".",
		"*",
		"/",
		"//",
		"%",
		".*",
		"./",
		".//",
		"!",
		"~",
		"**",
		"(",
		")",
		"md_token_string_literal",
		"md_token_int_literal",
		"md_token_float_literal",
		"md_token_boolean_literal",
		"md_token_M_PI",
		"md_token_M_E",
		"md_token_panic",
		"[",
		",",
		"md_token_IPS",
		"md_token_IFS",
		"md_token_IRS",
		"md_token_OPS",
		"md_token_OFS",
		"md_token_ORS",
		"md_token_NF",
		"md_token_NR",
		"md_token_FNR",
		"md_token_FILENAME",
		"md_token_FILENUM",
		"md_token_int_type",
		"md_token_float_type",
		"md_token_begin",
		"md_token_end",
		"md_token_if",
		"md_token_elif",
		"md_token_else",
		"md_token_while",
		"md_token_do",
		"md_token_for",
		"md_token_in",
		"md_token_func",
		"md_token_return",
	},

	idMap: map[string]Type{
		"INVALID":                     0,
		"$":                           1,
		"empty":                       2,
		";":                           3,
		"{":                           4,
		"}":                           5,
		"=":                           6,
		"md_token_filter":             7,
		"md_token_emit":               8,
		"md_token_dump":               9,
		"md_token_edump":              10,
		"md_token_field_name":         11,
		"$[":                          12,
		"]":                           13,
		"md_token_braced_field_name":  14,
		"md_token_full_srec":          15,
		"md_token_oosvar_name":        16,
		"@[":                          17,
		"md_token_braced_oosvar_name": 18,
		"md_token_full_oosvar":        19,
		"md_token_non_sigil_name":     20,
		"||=":                         21,
		"^^=":                         22,
		"&&=":                         23,
		"|=":                          24,
		"^=":                          25,
		"<<=":                         26,
		">>=":                         27,
		">>>=":                        28,
		"+=":                          29,
		".=":                          30,
		"-=":                          31,
		"*=":                          32,
		"/=":                          33,
		"//=":                         34,
		"%=":                          35,
		"**=":                         36,
		"?":                           37,
		":":                           38,
		"||":                          39,
		"^^":                          40,
		"&&":                          41,
		"=~":                          42,
		"!=~":                         43,
		"==":                          44,
		"!=":                          45,
		">":                           46,
		">=":                          47,
		"<":                           48,
		"<=":                          49,
		"|":                           50,
		"^":                           51,
		"&":                           52,
		"<<":                          53,
		">>":                          54,
		">>>":                         55,
		"+":                           56,
		"-":                           57,
		".+":                          58,
		".-":                          59,
		".":                           60,
		"*":                           61,
		"/":                           62,
		"//":                          63,
		"%":                           64,
		".*":                          65,
		"./":                          66,
		".//":                         67,
		"!":                           68,
		"~":                           69,
		"**":                          70,
		"(":                           71,
		")":                           72,
		"md_token_string_literal":     73,
		"md_token_int_literal":        74,
		"md_token_float_literal":      75,
		"md_token_boolean_literal":    76,
		"md_token_M_PI":               77,
		"md_token_M_E":                78,
		"md_token_panic":              79,
		"[":                           80,
		",":                           81,
		"md_token_IPS":                82,
		"md_token_IFS":                83,
		"md_token_IRS":                84,
		"md_token_OPS":                85,
		"md_token_OFS":                86,
		"md_token_ORS":                87,
		"md_token_NF":                 88,
		"md_token_NR":                 89,
		"md_token_FNR":                90,
		"md_token_FILENAME":           91,
		"md_token_FILENUM":            92,
		"md_token_int_type":           93,
		"md_token_float_type":         94,
		"md_token_begin":              95,
		"md_token_end":                96,
		"md_token_if":                 97,
		"md_token_elif":               98,
		"md_token_else":               99,
		"md_token_while":              100,
		"md_token_do":                 101,
		"md_token_for":                102,
		"md_token_in":                 103,
		"md_token_func":               104,
		"md_token_return":             105,
	},
}
