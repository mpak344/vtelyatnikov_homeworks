package FrequencyAnalyser
import (
	"fmt"
	"strings"
	"sort"
)

type WordType struct {
	str string
	count int
};

// Написать функцию, которая получает на вход текст и возвращает 10 самых часто встречающихся слов без учета словоформ
func GetPopularWords(words string) (res []string) {
	// разбиваем строку на отдельные слова
	splitedWords := strings.Split(words, " ")
	wordsMap := make(map[string]int)
	// проходим по массиву слов, приводим их к нижнему регистру и убираем знаки препинания,
	// затем складываем в словарь со счётчиками
	for _, v := range splitedWords {
		rawString := getRawString(v)
		wordsMap[rawString] = wordsMap[rawString] + 1  
	} 

	// Проходим по словорю и полученные данные помещаем в слайс для сортировки
	wordsList := []WordType {}
	for key, value := range wordsMap {
		var word WordType
		word.str = key
		word.count = value
		wordsList  = append(wordsList, word)
	}
	
	// Сортируем слова по убыванию колличества упоминаний
	sort.Slice(wordsList, func(i, j int) bool { 
		return wordsList[i].count > wordsList[j].count 
	})

	length := len(wordsList)
	if (length > 10 ) {
		length = 10
	}
	for _, value := range wordsList[:length] {
		res = append(res, value.str)
	}
	fmt.Printf("%v", res)
	return res
}

// На входе функция получает слово не обработанное слово(слово, которое может начинаться с большой буквы и иметь знаки препинания в конце) и
// возвращает обработанное, которое может использоваться для дальнейшего анализа.
func getRawString(w string) (res string) {
	res = strings.ToLower(w)
	for i:= len(res) - 1; i >= 0; i-- {
		symb := res[i]
		if (symb >= 'A') {
			break
		}
		res = res[0: i]
	}
	return res
}