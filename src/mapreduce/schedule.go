package mapreduce

import (
	"fmt"
	"sync"
)

//
// schedule() starts and waits for all tasks in the given phase (mapPhase
// or reducePhase). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers. Once all tasks
	// have completed successfully, schedule() should return.
	//
	// Your code here (Part III, Part IV).
	//
	var (
		wg       sync.WaitGroup              // wait until all the tasks have been executed
		doneChan = make(chan string, ntasks) // inform scheduler of the termination of one task
	)
	for i := 0; i < ntasks; i++ {
		var workerID string
		select {
		case x := <-registerChan:
			workerID = x
		case x := <-doneChan:
			workerID = x
		}

		wg.Add(1)
		// assign a task to an available worker via rpc
		go func(i int, workerID string) {
			defer wg.Done()

			var args DoTaskArgs
			switch phase {
			case mapPhase:
				args = DoTaskArgs{
					JobName:       jobName,
					File:          mapFiles[i],
					Phase:         phase,
					TaskNumber:    i,
					NumOtherPhase: n_other,
				}
			case reducePhase:
				args = DoTaskArgs{
					JobName:       jobName,
					Phase:         phase,
					TaskNumber:    i,
					NumOtherPhase: n_other,
				}
			}
			call(workerID, "Worker.DoTask", &args, nil)
			doneChan <- workerID
		}(i, workerID)
	}
	wg.Wait()
	fmt.Printf("Schedule: %v done\n", phase)
}

//	issues found during the debugging:
//	- The task number i should be passed as parameter to the goroutine,
//	  otherwise, data race may occur on i and mapFiles[i].
//	- doneChan should be created by calling make() to avoid a nil channel
//	  leading to the deadlock.
//	- doneChan must be a buffered channel. If it was a non-buffered one,
//	  there would be workers being blocked forever on the doneChan since
//	  the scheduler main goroutine would not read from the channel any
//	  more after it has gone out of the for loop. The buffer size is set
//	  to ntasks since the maximum number of parallel workers can not
//	  exceed ntasks.
