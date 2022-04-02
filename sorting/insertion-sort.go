package sorting

/*
	Insertion Sort - Сортировка вставкой
	Худшее время: O(n^2)
	Среднее время: O(n^2)
	Лучшее время: O(n)
	Память: O(n)

	При сортировке вставками массив постепенно перебирается слева направо.
	При этом каждый последующий элемент размещается так, чтобы он оказался между
	ближайшими элементами с минимальным и максимальным значением.
*/
func InsertionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > v {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = v

	}
}