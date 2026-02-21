package models

import (
	"time"

	"go.uber.org/zap"
)

type Task struct {
	Title    string
	Job      func() error
	Interval time.Duration
}

type Scheduler struct {
	slogger *zap.SugaredLogger
	tasks   []Task
	quit    chan struct{}
}

func NewScheduler(slogger *zap.SugaredLogger) *Scheduler {
	return &Scheduler{
		slogger: slogger,
		tasks:   []Task{},
		quit:    make(chan struct{}),
	}
}

func (s *Scheduler) AddTask(title string, interval time.Duration, job func() error) {
	s.tasks = append(s.tasks, Task{
		Title:    title,
		Job:      job,
		Interval: interval,
	})
}

// Start runs all tasks on their respective intervals in separate goroutines.
func (s *Scheduler) Start() {
	for _, task := range s.tasks {
		// Launch a goroutine per task to allow them to run independently
		go func(t Task) {
			// Create a ticker that ticks at the task's interval
			ticker := time.NewTicker(t.Interval)
			defer ticker.Stop()

			for {
				select {
				// On every tick, run the task in its own goroutine
				case <-ticker.C:
					go func() {
						s.slogger.Infow("started scheduled job", "job", t.Title)
						err := t.Job()
						if err != nil {
							s.slogger.Infow("scheduled job failed", "job", t.Title, "err", err)
							return
						}
						s.slogger.Infow("finished scheduled job", "job", t.Title)
					}()

					// If we receive from the quit channel, exit the loop and stop the task
				case <-s.quit:
					return
				}
			}
		}(task)
	}
}

// Stop stops all running tasks by closing the quit channel.
// All goroutines listening to it will exit cleanly.
func (s *Scheduler) Stop() {
	close(s.quit)
}
