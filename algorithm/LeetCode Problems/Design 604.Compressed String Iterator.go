package main

// 604. Compressed String Iterator
// https://leetcode.com/problems/compressed-string-iterator/
/*
restate :
Design and implement a data structure for a compressed string iterator.
The compressed string will be in the form of each letter followed by a positive integer
representing the number of times the letter appears in the original uncompressed string.

Implement the StringIterator class:
- StringIterator(string compressedString) Initializes the object with the compressed string compressedString.
- char next() Returns the next character if the original string still has uncompressed characters, otherwise returns a white space character ' '.
- boolean hasNext() Returns true if there are still uncompressed characters, otherwise returns false.

Example 1:
Input
["StringIterator", "next", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[["L1e2t1C1o1d1e1"], [], [], [], [], [], [], [], []]
Output
[null, "L", "e", "e", true, "t", true, "C", true]

Explanation
StringIterator stringIterator = new StringIterator("L1e2t1C1o1d1e1");
stringIterator.next(); // return 'L'
stringIterator.next(); // return 'e'
stringIterator.next(); // return 'e'
stringIterator.hasNext(); // return True
stringIterator.next(); // return 't'
stringIterator.hasNext(); // return True
stringIterator.next(); // return 'C'
stringIterator.hasNext(); // return True
*/

/*
	Thinking process:
	1. We can use a pointer to keep track of our current position in the compressed string.
	2. For each call to next(), we check if the current character is a letter or a digit.
	   - If it is a letter, we store it as the current character and move the pointer to the next character.
	   - If it is a digit, we calculate the count of the current character by parsing the digits until we reach a non-digit character.
	3. We then return the current character and decrement the count. If the count reaches zero, we move to the next character in the compressed string.

	Time complexity: O(n) where n is the length of the compressed string.
	Space complexity: O(1) since we are using only a constant amount of extra space.
*/

type StringIterator struct {
	s   string
	pos int
}

func Constructor(compressedString string) StringIterator {
	return StringIterator{
		s:   compressedString,
		pos: 0,
	}
}

func (this *StringIterator) Next() string {
	if this.pos >= len(this.s) {
		return " "
	}

	ch := this.s[this.pos]
	this.pos++

	count := 0
	for this.pos < len(this.s) && this.s[this.pos] >= '0' && this.s[this.pos] <= '9' {
		count = count*10 + int(this.s[this.pos]-'0')
		this.pos++
	}

	if count > 0 {
		this.pos -= 1 // Move back to the last digit
	}
	return string(ch)
}

func (this *StringIterator) HasNext() bool {
	return this.pos < len(this.s)
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	iterator := Constructor("L1e2t1C1o1d1e1")
	println(iterator.Next())    // return 'L'
	println(iterator.Next())    // return 'e'
	println(iterator.Next())    // return 'e'
	println(iterator.HasNext()) // return true
	println(iterator.Next())    // return 't'
	println(iterator.HasNext()) // return true
	println(iterator.Next())    // return 'C'
	println(iterator.HasNext()) // return true
}
