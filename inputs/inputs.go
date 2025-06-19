package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UserMetric struct {
	Weight float64
	Height float64
}

var reader = bufio.NewReader(os.Stdin)

func PrintWelcome() {
	fmt.Println("BMI (Body Mass Index) Calculator")
	fmt.Println("-----------------------------------")
}

func GetUserMetrics() UserMetric  {
	weight := getInputFromCli("Enter your weight (kg): ")
	height := getInputFromCli("Enter your height (m): ")
	return UserMetric{weight, height}
}

func getInputFromCli(promtText string) (value float64) {
   fmt.Print(promtText)
   input, _ := reader.ReadString('\n')
   input = strings.Replace(input, "\n", "", -1)
   value,_ = strconv.ParseFloat(input, 64)
   return
} 


