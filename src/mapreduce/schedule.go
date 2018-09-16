package mapreduce

import (
	"fmt"
	"sync"
)

type done struct {
	workerID   string
	taskNumber int
	status     bool
}

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
		// wait until all the tasks have been executed
		wg sync.WaitGroup
		// inform scheduler of the termination of one task
		doneChan = make(chan done, ntasks)
		// list of failing taskNumber
		failedTasks []int
	)
	for i := 0; i < ntasks || len(failedTasks) > 0; {
		// determine who runs which task
		workerID, taskNumber, ok := func() (workerID string, taskNumber int, ok bool) {
			// next workerID
			select {
			case x := <-registerChan:
				workerID = x
			case x := <-doneChan:
				if !x.status {
					// blacklist the worker due to execution failure
					// The worker won't be visible until its re-registration on the master.
					// Notice that a failued task might fail again on a new worker.
					// If it is the case, simply append it to the end of the failedTasks
					// and wait for a healthy worker becomes available.
					failedTasks = append(failedTasks, x.taskNumber)
					return "", -1, false
				}
				workerID = x.workerID
			}

			// next taskNumber
			if i < ntasks {
				taskNumber = i
				i++
			} else {
				// all the original tasks have been executed
				// start reprocessing the failure
				taskNumber = failedTasks[0]
				failedTasks = failedTasks[1:]
			}

			return workerID, taskNumber, true
		}()

		// current worker has failed to run the task assigned
		if !ok {
			continue
		}

		// assign the task to an available worker via rpc
		wg.Add(1)
		go func(workerID string, taskNumber int) {
			defer wg.Done()

			var args DoTaskArgs
			switch phase {
			case mapPhase:
				args = DoTaskArgs{
					JobName:       jobName,
					File:          mapFiles[taskNumber],
					Phase:         phase,
					TaskNumber:    taskNumber,
					NumOtherPhase: n_other,
				}
			case reducePhase:
				args = DoTaskArgs{
					JobName:       jobName,
					Phase:         phase,
					TaskNumber:    taskNumber,
					NumOtherPhase: n_other,
				}
			}
			status := call(workerID, "Worker.DoTask", &args, nil)
			doneChan <- done{workerID, taskNumber, status}
		}(workerID, taskNumber)
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
