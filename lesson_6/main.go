package main

import (
	"flag"
	"fmt"

	copyutil "github.com/mpak344/vtelyatnikov_homeworks/lesson_6/CopyUtil"
)

/*Копирование файлов
Цель: Реализовать утилиту копирования файлов Утилита должна принимать следующие аргументы
* файл источник (From)
* файл копия (To)
* Отступ в источнике (Offset), по умолчанию - 0
* Количество копируемых байт (Limit),
по умолчанию - весь файл из From Выводить в консоль прогресс копирования в %,
например с помощью github.com/cheggaaa/pb Программа может НЕ обрабатывать файлы,
у которых не известна длинна (например /dev/urandom).

Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать функцию вида Copy(from string, to string, limit int, offset int) error
Написать unit-тесты на функцию Copy
Реализовать функцию main, анализирующую параметры командной строки и вызывающую Copy
Проверить установку и работу утилиты руками
Критерии оценки: Функция должна проходить все тесты
Все необходимые для тестов файлы должны создаваться в самом тесте
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать, проверить и установить пакет с помощью go get / go test / go install
Рекомендуем сдать до: 07.10.2019*/

func main() {
	var from, to string
	var limit, offset int

	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to read to")
	flag.IntVar(&limit, "limit", 0, "Limit for copyed data")
	flag.IntVar(&offset, "offset", 0, "Offset for copyed data")

	flag.Parse()

	err := copyutil.Copy(from, to, int64(limit), int64(offset))
	if err != nil {
		fmt.Printf("have error %v", err)
	}
}
