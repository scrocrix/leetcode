package main

import (
	"bytes"
	"fmt"
)

func leastInterval(tasks []byte, unitOfTime int) int {
	tasks = bytes.Replace(tasks, []byte(` `), []byte(``), len(tasks))

	if unitOfTime == 0 {
		return len(tasks)
	}

	var intervals []string
	var preInterval []string

	tasksToRun := tasks

	i := 0

	for len(tasksToRun) > 0 {
		if i >= len(tasksToRun) {
			i = 0
		}

		task := string(tasksToRun[i])

		if i+1 < len(tasksToRun) && task == string(tasksToRun[i+1]) {
			i += 1
			continue
		}

		if len(preInterval) >= unitOfTime {
			intervals = append(intervals, preInterval...)
			fmt.Println(intervals)
			preInterval = preInterval[:0]
		}

		preInterval = append(preInterval, task)

		tasksToRun = append(tasksToRun[:i], tasksToRun[i+1:]...)

		i += 1
	}

	if len(preInterval) <= unitOfTime {
		intervals = append(intervals, preInterval...)
	}

	fmt.Println(intervals)

	return len(intervals)
}

func main() {
	//fmt.Println(leastInterval([]byte(`A A A B B B`), 0))
	fmt.Println(leastInterval([]byte(`A A A B B B`), 2))
	//fmt.Println(leastInterval([]byte(`A A A A A A B C D E F G`), 2))
	//fmt.Println(leastInterval([]byte(`A B C D E F G`), 2))
}
