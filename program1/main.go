/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-10-01
 * @fileoverview This program converts a temperature of 34 degrees Celsius to its Fahrenheit equivalent using the formula: F = 9 / 5 * C + 32.
 */

package main

import "fmt"

func main() {
	// This line uses the Celsius to Fahrenheit conversion (F = 9 / 5 * C + 32) to calculate the temperature. Every number ends in .0 to convert them into floats and prevent integer division.
	fmt.Println("34 degrees Celsius = " + fmt.Sprint(9.0 / 5.0 * 34.0 + 32.0) + " degrees Fahrenheit.")
	
	fmt.Println("\nDone.")
}

