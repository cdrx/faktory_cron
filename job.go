package main

import (
	faktory "github.com/contribsys/faktory/client"
	"github.com/robfig/cron"
)

type Job struct {
	Schedule string        `yaml:"schedule"`
	Name     string        `yaml:"job"`
	Args     []interface{} `yaml:"args"`
	Queue    string        `yaml:"queue"`
	Retries  int           `yaml:"retries"`
	Priority uint8         `yaml:"priority"`
}

func (j *Job) Start() {
	// send the task to faktory
	if len(j.Queue) == 0 {
		j.Queue = "default"
	}

	log.Debugf("Running %v (queue: %v, args: %v)", j.Name, j.Queue, j.Args)

	client, err := faktory.Open()
	if err != nil {
		log.Warnf("Failed to send %v to faktory: %v", j.Name, err)
		return
	}
	defer client.Close()

	job := faktory.NewJob(j.Name, j.Args...)
	if j.Retries > 0 {
		job.Retry = j.Retries
	}
	if j.Priority > 0 {
		job.Priority = j.Priority
	}
	if len(job.Queue) > 0 {
		job.Queue = j.Queue
	}

	err = client.Push(job)

	if err != nil {
		log.Warnf("Failed to push %v to faktory: %v", j.Name, err)
		return
	}
}

func (j *Job) GetFunc() cron.FuncJob {
	return func() {
		go j.Start()
	}
}

func (j *Job) AddToScheduler() {
	scheduler.AddFunc(j.Schedule, j.GetFunc())
}
