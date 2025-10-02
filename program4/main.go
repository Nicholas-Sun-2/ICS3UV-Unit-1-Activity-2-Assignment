/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-10-01
 * @fileoverview This program calculates the gross pay for an employee named Fred, who worked 40 hours at an hourly wage of $8.20
 */

package main

import "fmt"

func main() {
	// This line displays the introduction to the calculation.
	fmt.Println("Fred worked 40 hours at an hourly wage of $8.20.")

	// This line uses the gross pay formula gross pay = hours worked * hourly wage to calculate Fred's gross pay.
	fmt.Println("His gross pay is $" + fmt.Sprint(40 * 8.20) + ".")

	fmt.Println("\nDone.")
}

