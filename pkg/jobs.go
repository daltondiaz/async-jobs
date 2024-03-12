package pkg

import (
	"log"
	"os/exec"
	"strconv"
	"sync"
)

type Job struct {
	cmd string
	arg int
}

func fetch() []Job {
	// get jobs from database
	var jobs []Job
	for i := 0; i < 1000; i++ {
		var job Job
		job.cmd = "/home/dalton/dev/personal/async-jobs/test.php"
		job.arg = i
		jobs = append(jobs, job)
	}
	return jobs
}

func execution(job Job, wg *sync.WaitGroup) {
	cmd := exec.Command("php", job.cmd, strconv.Itoa(job.arg))
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(stdout))
	defer wg.Done()
}

func Run() {
	// run jobs according to cron
	jobs := fetch()
	var wg sync.WaitGroup
	for i, job := range jobs {
		wg.Add(1)
		log.Println(i)
		go execution(job, &wg)
	}
	wg.Wait()
}
