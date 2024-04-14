// Adding bindings for ctype.h functions

package goncurses

// #include <ctype.h>
import "C"

// Checks whether the passed character is alphanumeric.
func IsAlNum(ch int) bool {
	return C.isalnum(C.int(ch)) != 0
}

// Checks whether the passed character is alphabetic.
func IsAlpha(ch int) bool {
	return C.isalpha(C.int(ch)) != 0
}

// Checks whether the passed character is control character.
func IsCntrl(ch int) bool {
	return C.iscntrl(C.int(ch)) != 0
}

// Checks whether the passed character is decimal digit.
func IsDigit(ch int) bool {
	return C.isdigit(C.int(ch)) != 0
}

// Checks whether the passed character has graphical representation using locale.
func IsGraph(ch int) bool {
	return C.isgraph(C.int(ch)) != 0
}

// Checks whether the passed character is lowercase letter.
func IsLower(ch int) bool {
	return C.islower(C.int(ch)) != 0
}

// Checks whether the passed character is printable.
func IsPrint(ch int) bool {
	return C.isprint(C.int(ch)) != 0
}

// Checks whether the passed character is a punctuation character.
func IsPunct(ch int) bool {
	return C.ispunct(C.int(ch)) != 0
}

// Checks whether the passed character is white-space.
func IsSpace(ch int) bool {
	return C.isspace(C.int(ch)) != 0
}

// Checks whether the passed character is an uppercase letter.
func IsUpper(ch int) bool {
	return C.isupper(C.int(ch)) != 0
}

// Checks whether the passed character is a hexadecimal digit.
func IsXDigit(ch int) bool {
	return C.isxdigit(C.int(ch)) != 0
}

// Converts uppercase letters to lowercase.
func ToLower(ch int) int {
	return int(C.tolower(C.int(ch)))
}

// Converts lowercase letters to uppercase.
func ToUpper(ch int) int {
	return int(C.toupper(C.int(ch)))
}
