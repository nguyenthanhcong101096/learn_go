package main

import "fmt"

type Job interface {
	Process()
}

type Worker struct {
	WorkerId   int
	Done       chan bool
	JobRunning chan Job
}

func NewWorker(WorkerId int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerId:   WorkerId,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run() {
	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("Running job: ", w.WorkerId)
				job.Process()
			case <-w.Done:
				fmt.Println("Stop worker")
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.Done <- true
}

type JobQuene struct {
	Worker     []*Worker
	JobRunning chan Job
	Done       chan bool
}

func NewJobQuene(numberOfJob int) JobQuene {
	wokers := make([]*Worker, numberOfJob, numberOfJob)
	JobRunning := make(chan Job)

	for i := 0; i < numberOfJob; i++ {
		wokers[i] = NewWorker(i, JobRunning)
	}

	return JobQuene{
		Worker:     wokers,
		JobRunning: JobRunning,
		Done:       make(chan bool),
	}
}

func (j *JobQuene) Push(job Job) {
	j.JobRunning <- job
}

func (j *JobQuene) Stop() {
	j.Done <- true
}

func (j *JobQuene) Start() {
	go func() {
		for i := 0; i < len(j.Worker); i++ {
			j.Worker[i].Run()
		}
	}()

	go func() {
		for {
			select {
			case <-j.Done:
				for i := 0; i < len(j.Worker); i++ {
					j.Worker[i].Stop()
				}
				return
			}
		}
	}()
}

type Sender struct {
	Email string
}

func (e Sender) Process() {
	fmt.Println(e.Email)
}

func main() {
	emails := []string{
		"A@gmail.com",
		"B@gmail.com",
		"C@gmail.com",
		"D@gmail.com",
	}

	jobQuene := NewJobQuene(5)
	jobQuene.Start()

	for _, email := range emails {
		sender := Sender{email}
		jobQuene.Push(sender)
	}
}
