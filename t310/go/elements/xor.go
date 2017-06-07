package elements

import "fmt"

/*
func xor(input1 []int, input2 []int) []int {

var output_slice []int

let n be the length of input1
for i from 0 to n {
    a is the nth element of input1
    b is the nth element of input2
    c is a XOR b
    output_slice = append(output_slice, c)
  }
    return output_slice
}
*/

func xor(input1 []int, input2 []int) []int {

	var output_slice []int

	fmt.Println("xor() has come into being")
	fmt.Println("Input 1 is:", input1)
	fmt.Println("Input 2 is:", input2)

	n := len(input1)
	fmt.Println("Length of input1 is", n)

	for i := 0; i <= n; i++ {
		a := input1[n]
		b := input2[n]
		var c int
		c = a ^ b
		output_slice = append(output_slice, c)

		fmt.Println("a is", a, ", b is", b, "XOR is", c)
		fmt.Println("output_slice is currently", output_slice)
	}

	return output_slice
}

/*
for i := range a {
    if a[i] != b[i] {
        return false
*/

/*
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
*/
