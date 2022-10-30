package main

// Problem Description

/**************************************************************************************

6. Zigzag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000`


**************************************************************************************/

// Solution steps
/*******************************

1. ColumnStep = numRows-1 * 2
2. DiagnolStep = ColumnStep - 2
3. Iterate given numRows
	` for row:=0;row<numRows; row =row+1`
4. Check does this row have diagonal cells ? except first and last low ,
	`diag := row > 0 && row < numRows-1`
5. Add characters for each row by incrementing ColumnStep and DiagnolStep
	`for j:=row; j < len(s) ; j+=ColumnStep
		append s[j] to result

		if diag
		append s[j+DiagnolStep] to result


	`
6. If diag Decrement DiagnolStep by 2
	`diagStep -= 2`

********************************/
func main() {

}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	ss := make([]uint8, len(s), len(s))
	colStep := (numRows - 1) * 2
	diagStep := colStep - 2

	i := 0 // i = position to write to in ss

	for row := 0; row < numRows; row = row + 1 {
		diag := row > 0 && row < numRows-1

		for j := row; j < len(s); j += colStep {
			ss[i] = s[j]
			i = i + 1
			if diag && j+diagStep < len(s) {
				ss[i] = s[j+diagStep]
				i = i + 1
			}
		}
		if diag {
			diagStep = diagStep - 2
		}
	}
	return string(ss)
}
