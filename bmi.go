package main

import (
	"fmt"

	"github.com/jephterSmart/go-basics/inputs"
)

func main(){
	inputs.PrintWelcome()
	userMetric := inputs.GetUserMetrics()
	bmi := calculateBMI(userMetric.Weight, userMetric.Height)
	printBMI(bmi)	
}

func calculateBMI(weight float64, height float64) (bmi float64) {
	bmi = weight/ (height * height)
	return
}

func printBMI(bmi float64) {
	fmt.Printf("Your BMI is %.2f\n", bmi)
}
