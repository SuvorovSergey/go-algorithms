package sorting

/*
	Selection Sorting - Сортировка выбором
	Худшее время: O(n^2)
	Среднее время: O(n^2)
	Лучшее время: O(n^2)
	Память: O(n)

	Сначала нужно рассмотреть подмножество массива и найти в нём максимум
	или минимум). Затем выбранное значение меняют местами со значением первого
	неотсортированного элемента. Этот шаг нужно повторять до тех пор, пока в
	массиве не закончатся неотсортированные подмассивы.
*/
func SelectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}
}
