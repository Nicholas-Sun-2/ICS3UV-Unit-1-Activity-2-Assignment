/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-10-01
 * @fileoverview This program outputs the volume of a sphere with a radius of 4 cm.
 */

package main

import "fmt"

func main() {
	// This line uses the volume of a sphere formula (v = (4/3) * Pi * r^3) with Pi approximated as 3.14. Every integer ends in .0 to convert them into floats and prevent integer division.
	fmt.Println("The volume of a sphere with a radius of 4 cm = " + fmt.Sprint(4.0 / 3.0 * 3.14 * 4.0 * 4.0 * 4.0) + " cm^3.")
	
	fmt.Println("\nDone.")
}

