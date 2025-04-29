package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Taxes(firstSel float64, Prize float64, ExtraExp float64, ch chan float64) {
	defer wg.Done()
	defer close(ch)
	totalIncome := firstSel + Prize + ExtraExp
	totalTaxes := float64(totalIncome) * 0.15
	ch <- totalTaxes
}

func Salary(WorkHours float64, RatePerHour float64, Prize float64, Extra float64, Deductions float64, ch chan float64) {
	defer wg.Done()
	defer close(ch)
	totalSalIncome := WorkHours * RatePerHour
	totalSalary := totalSalIncome + Prize + Extra - Deductions
	ch <- totalSalary

}

func CycleInput(visibleText string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(visibleText)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("Некорректно введено число, попробуйте ещё раз: ")
			continue
		}
		return value
	}
}

func main() {
	chTax := make(chan float64)
	chSal := make(chan float64)

	firstSel := CycleInput("Сначала рассчитаем налоги. Введите оклад: ") //Salary - зарплата и оклад одновременно =))))))))))))))))))))))
	prize := CycleInput("Введите премию: ")
	extraExp := CycleInput("Введите надбавку за стаж: ")

	wg.Add(1)

	go Taxes(firstSel, prize, extraExp, chTax)

	workHours := CycleInput("Считаем зарплату. Введите количество отработанных часов: ")
	ratePerHour := CycleInput("Введите ставку за час: ")
	extra := CycleInput("Введите надбавки: ")
	deduct := CycleInput("Введите удержания: ")

	wg.Add(1)

	go Salary(workHours, ratePerHour, prize, extra, deduct, chSal)
	taxes := <-chTax
	salary := <-chSal
	result := salary - taxes

	wg.Wait()

	fmt.Println("Чистая зарплата: ", result)
}
