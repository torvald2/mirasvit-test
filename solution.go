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
	blockLen := len(blocks)
	if blockLen == 0 {
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
	for i := 1; i < blockLen; i++ {
		for k := range blockTable {
			if blocks[i][k] {
				blockTable[k] = append(blockTable[k], i)
			}
		}
	}

	dimentions := len(blockTable)

	// Проходимо у циклі усі блоки, створюємо вектор відстаней по кожному параметру і зрівнюємо з basic
	//ssfdsf
	bestIndex := -1
	bestDistance := math.MaxInt
	thisVec := make([]int, dimentions)

	for i := 0; i < len(blocks); i++ {
		for i := 0; i < dimentions; i++ {
			thisVec[i] = 0
		}
		idx := 0
		minDist := math.MaxInt
		for k, v := range blockTable {
			minDist = math.MaxInt
			if !blocks[i][k] {
				for _, x := range v {
					dist := AbsInt(x - i)
					if dist < minDist {
						minDist = dist
					}
				}
				thisVec[idx] = minDist
			}
			idx++
		}
		dist := GetSquaredLength(thisVec)
		if dist < bestDistance {
			bestIndex = i
			bestDistance = dist
		}

	}

	return bestIndex

}

// Обчислюємо просто квадрат відстані замість порівняння векторів

func GetSquaredLength(vec []int) int {
	var sum int
	for _, val := range vec {
		sum += val * val
	}
	return sum
}
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
