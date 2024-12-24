package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки букв в слове и получения ключа для анаграмм
func sortString(s string) string {
	// Приводим слово к нижнему регистру
	s = strings.ToLower(s)
	// Преобразуем строку в срез рун
	runes := []rune(s)
	// Сортируем руны
	sort.Sort(runesByValue(runes))
	// Возвращаем отсортированное слово как строку
	return string(runes)
}

// Тип для сортировки рун
type runesByValue []rune

func (r runesByValue) Len() int           { return len(r) }
func (r runesByValue) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runesByValue) Less(i, j int) bool { return r[i] < r[j] }

// Функция для поиска всех множеств анаграмм
func findAnagramSets(words []string) map[string][]string {
	// Карта для хранения множества анаграмм
	anagrams := make(map[string][]string)

	// Проходим по всем словам
	for _, word := range words {
		// Получаем ключ (отсортированное слово)
		key := sortString(word)
		// Добавляем слово в соответствующее множество анаграмм
		anagrams[key] = append(anagrams[key], word)
	}

	// Результат: карта, где ключ - это слово, первое в множестве
	result := make(map[string][]string)

	// Перебираем все множества анаграмм
	for _, group := range anagrams {
		// Если множество состоит из более чем одного слова
		if len(group) > 1 {
			// Сортируем множество
			sort.Strings(group)
			// Первое слово будет ключом
			result[group[0]] = group
		}
	}

	return result
}

func main() {
	// Пример словаря
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"}

	// Ищем множества анаграмм
	anagramSets := findAnagramSets(words)

	// Выводим результат
	for key, group := range anagramSets {
		fmt.Printf("Ключ: %s, Множество: %v\n", key, group)
	}
}
