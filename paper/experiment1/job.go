package main

type Job interface {
	execute(jobType int)
}
