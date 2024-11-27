package mirasvittest

import (
	"math"
)

// На вхід приймає тільки лист однородних мап
// Це значить значення доступних об'єктів визначається по першому об'єкту листа
// Якщо в мапі наступних об'єктів відсутній певний об'єкт, який був зазначений у першому об'єкті листа то його значення
// рахується за False
// Якщо передаємо пустий лист то повернется 0
func FindSolution(blocks []map[string]bool) (bestBlockId int) {

	//Перевіряємо чи масив не пустий Якщо пустий повертаємо 0
	if len(blocks) == 0 {
		return
	}
	//Створюємо мапу з данних першого елементу масива з максимально можливим капасіті
	blockTable := make(map[string][]int, len(blocks[0]))
	for k, v := range blocks[0] {
		blockTable[k] = make([]int, 0, len(blocks))
		//Якщо об'єкт присутній у блоці то ставимо номер блоку
		if v {
			blockTable[k] = append(blockTable[k], 0)
		}
	}
	// Розкладуємо інщі об'єкти у створену мапу
	for i := 1; i < len(blocks); i++ {
		for k, _ := range blockTable {
			if blocks[i][k] {
				blockTable[k] = append(blockTable[k], i)
			}
		}
	}

	// Створюємо ідеальний вектор наближення до якого будемо рахувати (ідеальний вектор це вектор з нулями)
	dimentions := len(blockTable)
	basicVec := make([]int, dimentions)

	// Проходимо у циклі усі блоки, створюємо вектор відстаней по кожному параметру і зрівнюємо з basic

	bestIndex := -1
	bestDistance := math.MaxInt

	for i := 0; i < len(blocks); i++ {
		thisVec := make([]int, 0, dimentions)
		for k, v := range blockTable {
			minDist := math.MaxInt
			if blocks[i][k] {
				thisVec = append(thisVec, 0)
			} else {
				for _, x := range v {
					dist := AbsInt(x - i)
					if dist < minDist {
						minDist = dist
					}
				}
				thisVec = append(thisVec, minDist)

			}
		}
		dist := GetSquaredDistance(thisVec, basicVec, dimentions)
		if dist < bestDistance {
			bestIndex = i
			bestDistance = dist
		}

	}

	return bestIndex

}

// Евклідова відстань між векторами
func GetSquaredDistance(vec1, vec2 []int, dimensions int) int {

	var sum int
	for i := 0; i < dimensions; i++ {
		diff := vec1[i] - vec2[i]
		sum += diff * diff
	}
	return sum
}
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
