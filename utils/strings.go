package utils

import (
	"fmt"
	"unicode/utf8"
)

// A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside double quotes " ".

func Strings() {
	name := "Hello, World!"
	fmt.Println(name)
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		for i := 0; i < len(s); i++ {
			fmt.Printf("%x ", s[i])
		}
	}
}

func AccesingIndividualBytesOfAString() {
	name := "Hello, World!"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func AccessingIndividualCharactersOfAString() {
	name := "Hello, World!"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}

func printCharsForRune(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

// A rune is a builtin type in Go and it’s the alias of int32. Rune represents a Unicode code point in Go. It doesn’t matter how many bytes the code point occupies, it can be represented by a rune. Let’s modify the above program to print characters using a rune.
func Rune() {
	name := "Hello, World!"
	fmt.Printf("String: %s\n", name)
	printCharsForRune(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printCharsForRune(name)
	fmt.Printf("\n")
	printBytes(name)
}

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte position %d\n", rune, index)
	}
}

// The above program is a perfect way to iterate over the individual runes of a string. But Go offers us a much easier way to do this using the for range loop.

func AccessingIndividualRunesUsingForRangeLoop() {
	name := "Señor"
	charsAndBytePosition(name)
}

// byteSlice in line no. 8 of the program above contains the UTF-8 Encoded hex bytes of the string Café. The program prints

func CreatingAStringFromSliceOfBytes() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

// Decimal values also work and the above program will also print Café.
func DecimalEquivalentOfHexValues() {
	byteSlice := []byte{67, 97, 102, 195, 169} // decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}

// In the above program runeSlice contains the Unicode code points of the string Señor in hexadecimal. The program outputs
func CreatingAStringFromSliceOfRunes() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

// The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string. This method takes a string as an argument and returns the number of runes in it.
// As we discussed earlier, len(s) is used to find the number of bytes in the string and it doesn’t return the string length. As we already discussed, some Unicode characters have code points that occupy more than 1 byte. Using len to find out the length of those strings will return the incorrect string length.
func StringLength() {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))
}

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

// The == operator is used to compare two strings for equality. If both the strings are equal, then the result is true else it’s false.
func StringComparison() {
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "Hello"
	string4 := "World"
	compareStrings(string3, string4)
}

func StringConcatenation() {
	string1 := "Go"
	string2 := "is awesome"
	result := string1 + " " + string2 // first approach
	// result := fmt.Sprintf("%s %s", string1, string2) // second approach
	fmt.Println(result)
}

//	func mutate(s string) string {
//		s[0] = 'a' //any valid unicode character within single quote is a rune
//		return s
//	}
func mutate(s []rune) string {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	return string(s)
}

// Strings are immutable in Go. Once a string is created it’s not possible to change it.

func StringsAreImmutable() {
	h := "Hello"
	fmt.Println(mutate([]rune(h)))
}

// In line no.7 of the above program, the mutate function accepts a rune slice as an argument. It then changes the first element of the slice to 'a', converts the rune back to string and returns it. This method is called from line no. 13 of the program. h is converted to a slice of runes and passed to mutate in line no. 13. This program outputs aello
