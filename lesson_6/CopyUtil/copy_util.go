package copyutil

import (
	"fmt"
	"io"
	"os"
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

var chankSize int = 16

func copy(from *os.File, to *os.File, lim int64) error {
	buf := make([]byte, lim)
	var _lim int64 = 0
	for _lim < lim {
		n, err := from.Read(buf[_lim:])
		lim += int64(n)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Copy is
func Copy(from string, to string, limit int64, offset int64) error {
	fmt.Printf("start copy from %v to %v with limit %v and offset %v \n", from, to, limit, offset)
	file, err := os.Open(from)

	if err != nil {
		fmt.Printf("Failed open file %v with error %v", from, err)
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed get file information %v with error %v", from, err)
		return err
	}

	fileSize := fileInfo.Size()
	if limit <= 0 {
		limit = fileSize - offset
	}

	//if fileSize <= 0 {
	//	return fmt.Errorf("Can't determine file size")
	//}

	if int64(offset) > fileSize {
		return fmt.Errorf("Can't copy file %v offset > filesize", from)
	}

	newOffset, err := file.Seek(int64(offset), 0)
	if err != nil {
		return fmt.Errorf("Failed seek %v %v", from, newOffset)
	}

	toFile, err := os.Create(to)
	if err != nil {
		return err
	}

	file.Seek(offset, 0)

	err = copy(file, toFile, limit)

	file.Close()
	toFile.Close()

	fmt.Printf("have error %v", err)
	return nil
}
