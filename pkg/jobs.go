package pkg

import (
	"daltondiaz/async-jobs/db"
	"daltondiaz/async-jobs/models"
	"log"
	"os/exec"
	"sync"

	"github.com/robfig/cron/v3"
)

type Job struct {
	cmd string
	arg int
}

/*func fetch() []Job {
	// get jobs from database
	/* var jobs []Job
	for i := 0; i < 5; i++ {
		var job Job
		job.cmd = "/home/dalton/dev/personal/async-jobs/test.php"
		job.arg = i
		jobs = append(jobs, job)
	}
	return jobs
}*/

func execution(job models.Job, wg *sync.WaitGroup) {
	cmd := exec.Command("php", "echo", job.Args)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(stdout))
	defer wg.Done()
}

func Run() {
	// run jobs according to cron
	conn := db.Connect()
    // db.Insert(conn)
	jobs, err := db.GetAvailableJobs(conn)
	if err != nil {
		log.Println(err)
	}
	c := cron.New()
	for i, job := range jobs {
		log.Println(i)
		id, _ := c.AddFunc(job.Cron, func() {
			cmd := exec.Command("php", "echo", job.Args)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(stdout)
		})
        log.Println("id cron:", id)
	}
	/*var wg sync.WaitGroup
	for i, job := range jobs {
		wg.Add(1)
		log.Println(i)
		go execution(job, &wg)
	}
	wg.Wait()*/
}
