package parallelrunner

import (
	"fmt"
	"sync"
)

func worker(ind int, jobs <-chan func() error, errors chan<- error, wg *sync.WaitGroup, M int) {
	fmt.Println("Worker", ind, "started")
	defer wg.Done()
	for {
		select {
		// если есть работа, считываем её из канала и запускаем
		case j := <-jobs:
			res := j()
			// если в очереди ошибок больше или сколько нужно убиваем вокер
			if len(errors) >= M {
				fmt.Println("Worker", ind, "killed errors is max")
				return
			}
			// если работа окончилась с ошибкой, кладём её в канал с ошибками
			if res != nil {
				errors <- res
			}
			// если работы больше нету, убиваем воркер
			if len(jobs) == 0 {
				fmt.Println("Worker", ind, "killed")
				return
			}
		default:
			// если в очереди ошибок больше или сколько нужно убиваем вокер
			// если работы больше нету, убиваем воркер
			if len(errors) >= M || len(jobs) == 0 {
				fmt.Println("Worker", ind, "killed")
				return
			}
		}
	}
}

// Run d
func Run(tasks []func() error, N int, M int) (retErr error) {
	retErr = nil
	// создаём буферизированные каналы
	// jobs - используется для передачи задач
	jobs := make(chan func() error, len(tasks))
	// errors - используем для передачи ошибок
	errors := make(chan error, M)

	// Wg - для синхронизации воркеров
	var wg sync.WaitGroup
	// создаём N воркеров
	//если len(tasks) < N создаём чуть меньше воркеров
	var workersCount = N
	if len(tasks) < workersCount {
		workersCount = len(tasks)
	}
	for w := 0; w < workersCount; w++ {
		go worker(w, jobs, errors, &wg, M)
		wg.Add(1)
	}
	// передаём воркерам работу
	for _, j := range tasks {
		jobs <- j
	}

	// дожидаемся, когда воркеры доработают
	wg.Wait()
	close(jobs)
	close(errors)
	// если ошибок >= M генерируем ошибку
	if len(errors) >= M {
		retErr = fmt.Errorf("error")
	}
	return retErr
}
