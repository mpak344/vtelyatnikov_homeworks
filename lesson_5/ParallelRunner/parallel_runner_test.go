package parallelrunner

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	task := []func() error{
		func() error {
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("Test1")
			return nil
		},
		func() error {
			fmt.Println("Test2")
			return nil
		},
		func() error {
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("Test3")
			return nil
		},
		func() error {
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("Test4")
			return nil
		},
	}

	res := Run(task, 2, 2) == nil
	require.Equal(t, res, true, "len(task) > N")

	res = Run(task, 10, 2) == nil
	require.Equal(t, res, true, "len(task) < N")

	res = Run(task, 4, 2) == nil
	require.Equal(t, res, true, "len(task) == N")

	task2 := []func() error{
		func() error {
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("Test1")
			return fmt.Errorf("test error")
		},
		func() error {
			fmt.Println("Test2")
			return fmt.Errorf("test error")
		},
		func() error {
			time.Sleep(3000 * time.Millisecond)
			return fmt.Errorf("test error")
		},
		func() error {
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("Test4")
			return nil
		},
	}

	res = Run(task2, 2, 2) == nil
	require.Equal(t, res, false, "ErrCount > M")

	res = Run(task2, 2, 10) == nil
	require.Equal(t, res, true, "ErrCount < M")

	res = Run(task2, 2, 3) == nil
	require.Equal(t, res, false, "ErrCount == M")
}
