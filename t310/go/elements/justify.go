package elements

import "fmt"

/*
justify() takes three inputs - two integer slices, and a target length.
It checks that the slices are of the target length.
If a slice is too long, it will return an error message.
If a slice is too short, it will pad the shortfall with zeros.
*/

/*
(which by the way is probably a __terrible__ idea, but i don't have anyone
to discuss this with atm.)
*/

func justify(slice1 []int, slice2 []int, length int) ([]int, []int, string) {

	fmt.Println("justify() is called into being")
	fmt.Println("slice 1 is", slice1)
	fmt.Println("slice 2 is", slice2)
	fmt.Println("length is", length)

	if len(slice1) == len(slice2) && len(slice1) == length {
		return slice1, slice2, "length ok"
	} else {
		return slice1, slice2, "Length Mismatch! All stop!"
	}

}
