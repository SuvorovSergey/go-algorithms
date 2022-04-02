package sorting

/*
	Bubble Sort, пузырьковая сортировка
	Худшее время: O(n^2)
	Среднее время: O(n^2)
	Лучшее время: O(n)
	Память: O(1)

	Нужно последовательно сравнивать значения соседних элементов и менять числа местами, если предыдущее оказывается больше последующего. Таким образом элементы с большими значениями оказываются в конце списка, а с меньшими остаются в начале.

	Этот алгоритм считается учебным и почти не применяется на практике из-за низкой эффективности: он медленно работает на тестах, в которых маленькие элементы (их называют «черепахами») стоят в конце массива.
*/
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		swapped := false
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
			swapped = true
		}
		if !swapped {
			break
		}
	}
}
