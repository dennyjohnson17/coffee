package main

import 	"fmt"

func main() {
	var water, milk, coffee, cups, money int
	water, milk, coffee, cups, money = 500, 500, 500, 500, 500
	for {
		var command string
		fmt.Println("Введите команду (buy, fill, take, exit):")
		fmt.Scan(&command)

		switch command {
		case "buy":
			var choice string
			fmt.Println("Выберите напиток (1 - эспрессо, 2 - латте, 3 - капучино):")
			fmt.Scan(&choice)

			switch choice {
			case "1": // эспрессо
				if water >= 250 && coffee >= 16 && cups >= 1 {
					water -= 250
					coffee -= 16
					cups--
					money += 4
					fmt.Println("Вы получили эспрессо")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления эспрессо")
				}
			case "2": // латте
				if water >= 350 && milk >= 75 && coffee >= 20 && cups >= 1 {
					water -= 350
					milk -= 75
					coffee -= 20
					cups--
					money += 7
					fmt.Println("Вы получили латте")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления латте")
				}
			case "3": // капучино
				if water >= 200 && milk >= 100 && coffee >= 12 && cups >= 1 {
					water -= 200
					milk -= 100
					coffee -= 12
					cups--
					money += 6
					fmt.Println("Вы получили капучино")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления капучино")
				}
			default:
				fmt.Println("Неверный выбор")
			}

		case "fill":
			var addWater, addMilk, addCoffee, addCups int
			fmt.Println("Введите количество воды для добавления:")
			fmt.Scan(&addWater)
			water += addWater

			fmt.Println("Введите количество молока для добавления:")
			fmt.Scan(&addMilk)
			milk += addMilk

			fmt.Println("Введите количество кофе для добавления:")
			fmt.Scan(&addCoffee)
			coffee += addCoffee

			fmt.Println("Введите количество стаканчиков для добавления:")
			fmt.Scan(&addCups)
			cups += addCups

		case "take":
			fmt.Println("Вы забрали $", money)
			money = 0

		case "exit":
			return

		default:
			fmt.Println("Неверная команда")
		}
	}
}
