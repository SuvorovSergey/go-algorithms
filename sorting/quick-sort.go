package sorting

/*
	Quiq Sort - Быстрая сортировка https://www.geeksforgeeks.org/quick-sort/
	Худшее время: O(n^2)
	Среднее время: O(n log n)
	Лучшее время: O(n)
	Память: O(n)

	Этот алгоритм состоит из трёх шагов. Сначала из массива нужно выбрать один элемент — его обычно называют опорным. Затем другие элементы в массиве перераспределяют так, чтобы элементы меньше опорного оказались до него, а большие или равные — после. А дальше рекурсивно применяют первые два шага к подмассивам справа и слева от опорного значения.

	Быструю сортировку изобрели в 1960 году для машинного перевода: тогда словари хранились на магнитных лентах, а сортировка слов обрабатываемого текста позволяла получить переводы за один прогон ленты, без перемотки назад.
*/
func QuickSort(slice []int) {
	quiqSort(slice, 0, len(slice)-1)
}

func quiqSort(slice []int, low, high int) {
	if low < high {
		pivotIndex := partition(slice, low, high)
		quiqSort(slice, low, pivotIndex-1)
		quiqSort(slice, pivotIndex+1, high)
	}
}

func partition(slice []int, low int, high int) int {
	pivot := slice[high]
	s := low - 1
	for i := low; i <= high-1; i++ {
		if slice[i] < pivot {
			s++
			slice[i], slice[s] = slice[s], slice[i]
		}
	}
	s++
	slice[s], slice[high] = slice[high], slice[s]
	return s
}
