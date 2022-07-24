package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	conversions "week-4-homework-1-tohanilhan/conversions/conv"
)

var flag bool = true

type Calculator struct {
	functions []conversions.Conversion
}

func (c *Calculator) addConversion(m conversions.Conversion) {
	c.functions = append(c.functions, m)
}

func (c Calculator) doCalculation(name string, arg float64) (float64, error) {

	for _, f := range c.functions {
		// check if name is equal to the name of the function, case insensitive
		a, b := f.GetName()
		if strings.ToUpper(a+" "+b) == strings.ToUpper(name) {
			result := f.Calculate(arg)
			return result, nil
		}
	}
	return 0, fmt.Errorf("Function %s not found", name)
}

func main() {
	startCalculator()
}

func startCalculator() {
	myCalculator := Calculator{}

	// Celcius conversions
	myCalculator.addConversion(conversions.Celsius{Name: "Celcius", ConvertTo: "Farhenheit"})
	myCalculator.addConversion(conversions.Celsius{Name: "Celcius", ConvertTo: "Kelvin"})

	// Farhenheit conversions
	myCalculator.addConversion(conversions.Farhenheit{Name: "Farhenheit", ConvertTo: "Celcius"})
	myCalculator.addConversion(conversions.Farhenheit{Name: "Farhenheit", ConvertTo: "Kelvin"})

	// Kelvin conversions
	myCalculator.addConversion(conversions.Kelvin{Name: "Kelvin", ConvertTo: "Farhenheit"})
	myCalculator.addConversion(conversions.Kelvin{Name: "Kelvin", ConvertTo: "Celcius"})

	// Ask user for input
	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions to convert temperatures:")
	for _, f := range myCalculator.functions {

		fmt.Println(f.GetName())

	}

	for flag {

		// Get input from user
		fmt.Println("\nEnter your input or x to exit:")
		fmt.Print("Conversion: ")
		var fName string

		// scan the input from user with space without newline
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fName = scanner.Text()

		if fName == "x" || fName == "X" {
			flag = false

		} else {
			fmt.Print("Degree: ")
			var argument float64
			fmt.Scan(&argument)

			// Calculate
			result, err := myCalculator.doCalculation(fName, argument)
			if err != nil {
				fmt.Println(err)
			} else {
				a := FixNaming(strings.Split(fName, " "))
				fmt.Printf("Conversion of %s to %s is %.2f\n", a[0], a[1], result)
			}
		}

	}
	fmt.Println("Bye!")

}

// FixNaming converts the name of the function to the correct format
func FixNaming(a []string) []string {
	for i := range a {
		if a[i] == "C" || a[i] == "c" {
			a[i] = "Celsius"
		} else if a[i] == "F" || a[i] == "f" {
			a[i] = "Farhenheit"
		} else if a[i] == "K" || a[i] == "k" {
			a[i] = "Kelvin"
		}
	}
	return a
}
