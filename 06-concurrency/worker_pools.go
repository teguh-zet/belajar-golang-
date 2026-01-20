package concurrency

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	// Loop ini akan terus berjalan SELAMA channel 'jobs' masih terbuka
	// dan masih ada isinya.
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second) // Simulasi kerja berat (1 detik)
		fmt.Println("worker", id, "finished job", j)

		// Kirim hasil kerja ke channel results
		results <- j * 2
	}
}

func RunWorkerPool() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
