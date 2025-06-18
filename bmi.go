package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jephterSmart/go-basics/inputs"
)

func main(){
	fmt.Println("BMI (Body Mass Index) Calculator")
	fmt.Println("-----------------------------------")
	fmt.Print("Enter your weight (kg): ")
	weightInput,_ := inputs.Reader.ReadString('\n')
	fmt.Print("Enter your height (m): ")
	heightInput,_ := inputs.Reader.ReadString('\n')
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput,64 )

	bmi := weight/ (height * height)

	fmt.Printf("Your BMI is %.2f\n", bmi)
}
