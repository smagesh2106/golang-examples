package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	rand.Seed(time.Now().Unix())

	projects := make(chan string, 10)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go employee(projects, i)
	}

	for j := 0; j < 50; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}
	close(projects)
	wg.Wait()
}

func employee(proj chan string, emp int) {
	defer wg.Done()

	for {
		project, ok := <-proj
		if !ok {
			fmt.Printf("Employee : %d Exit\n", emp)
			return
		}

		fmt.Printf("Employee %d, Started Project %s\n", emp, project)
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("\nTime to sleep, ", sleep, "ms\n")
		fmt.Printf("Employee %d, Completed Project %s\n", emp, project)
	}
}
