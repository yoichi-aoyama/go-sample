package main

import (
	"context"
	"github.com/olivere/jobqueue"
	"github.com/olivere/jobqueue/mysql"
	"time"
)

/*
"github.com/olivere/jobqueue" example
*/
func main() {

	store, err := mysql.NewStore("root@tcp(127.0.0.1:3306)/jobqueue_e2e?loc=UTC&parseTime=true")

	if err != nil {
		panic(err)
	}

	// Create a manager with the MySQL store and 10 concurrent workers.
	m := jobqueue.New(
		jobqueue.SetStore(store),
		//jobqueue.SetConcurrency(10),
	)

	// Register one or more topics and their processor
	//m.Register("clicks", func(args ...interface{}) error {
	//	// Handle "clicks" topic
	//})

	// Start the manager
	err = m.Start()
	if err != nil {
		panic(err)
	}

	// Add a job: It'll be added to the store and processed eventually.
	err = m.Add(context.Background(), &jobqueue.Job{Topic: "clicks", Args: []interface{}{640, 480}})
	if err != nil {
		panic(err)
	}

	// Stop the manager, either via Stop/Close (which stops after all workers
	// are finished) or CloseWithTimeout (which gracefully waits for a specified
	// time span)
	err = m.CloseWithTimeout(15 * time.Second) // wait for 15 seconds before forced stop
	if err != nil {
		panic(err)
	}
}
