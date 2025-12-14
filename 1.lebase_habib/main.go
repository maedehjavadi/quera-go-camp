package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// this is first practice
func readInputLines(scanner *bufio.Scanner) ([]string, map[string][]string) {
	lines := make([]string, 0, 6)           // length = 0, capacity = 6
	categories := make(map[string][]string) // category: colors list[]

	for i := 0; i < 6; i++ { // Inja Line hamono gereftim va zakhire kardim ba formati le mikhastim:
		scanner.Scan()
		trimText := strings.TrimSpace(scanner.Text())
		lines = append(lines, trimText)
	}

	for _, line := range lines { // Inja har Line ro omadim datash ro pardazesh kardim va map sakhtim [index, value]
		keyValue := strings.Split(line, ":")
		if len(keyValue) < 2 {
			continue
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		colorList := strings.Fields(value) //***** string ro be []string tabdil mikone
		categories[key] = colorList
	}
	return lines, categories
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines, categories := readInputLines(scanner)

	shirts := categories["SHIRT"]
	pants := categories["PANTS"]
	caps := categories["CAP"]
	jackets := categories["JACKET"]
	coats := categories["COAT"]
	season := lines[5]

	for _, shirt := range shirts {
		for _, pant := range pants {
			if season == "SUMMER" {
				for _, CAP := range caps {
					fmt.Printf("SHIRT: %v PANTS: %v CAP: %v\n", shirt, pant, CAP)
				}
			} else {
				coats := append(coats, "")
				caps := append(caps, "")

				for _, coat := range coats {
					if season == "WINTER" {
						if coat == "" {
							for _, j := range jackets {
								fmt.Printf("SHIRT: %v PANTS: %v JACKET: %v\n", shirt, pant, j)

							}
						} else {
							fmt.Printf("COAT: %v SHIRT: %v PANTS: %v\n", coat, shirt, pant)
						}

					}
					if season == "SPRING" || season == "FALL" {
						lCoat := strings.ToLower(coat)
						if season == "FALL" && (lCoat == "yellow" || lCoat == "orange") {
							continue
							// وقتی coat == "yellow" یا coat == "orange" باشه
							// 								کلاً چاپ نمی‌شه

							// حلقه‌ی cap اجرا نمی‌شه

							// می‌ره سراغ coat بعدی
							// fmt.Printf("SHIRT: %v PANTS: %v CAP: %v\n", shirt, pant, ca)
						}
						for _, ca := range caps {
							if coat == "" && ca == "" {
								fmt.Printf("SHIRT: %v PANTS: %v\n", shirt, pant)
							} else {
								fmt.Printf("COAT: %v SHIRT: %v PANTS: %v CAP: %v\n", coat, shirt, pant, ca)
							}
						}

					}

				}
			}
			// fmt.Printf("SHIRT: %v PANTS: %v\n", shirt, pant)
		}
	}
}
