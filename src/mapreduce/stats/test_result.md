## Part I: Map/Reduce input and output
```
source setenv
cd src/mapreduce
go test -run Sequential
```
```
master: Starting Map/Reduce task test
Merge: read mrtmp.test-res-0
master: Map/Reduce task completed
master: Starting Map/Reduce task test
Merge: read mrtmp.test-res-0
Merge: read mrtmp.test-res-1
Merge: read mrtmp.test-res-2
master: Map/Reduce task completed
PASS
ok      mapreduce       1.978s
```

## Part II: Single-worker word count
```
source setenv
cd src/main
time go run wc.go master sequential pg-*.txt
```
```
master: Starting Map/Reduce task wcseq
Merge: read mrtmp.wcseq-res-0
Merge: read mrtmp.wcseq-res-1
Merge: read mrtmp.wcseq-res-2
master: Map/Reduce task completed

real    0m2.709s
user    0m2.748s
sys     0m0.426s
```
```
bash ./test-wc.sh
```
```
master: Starting Map/Reduce task wcseq
Merge: read mrtmp.wcseq-res-0
Merge: read mrtmp.wcseq-res-1
Merge: read mrtmp.wcseq-res-2
master: Map/Reduce task completed
Passed test
```

## Part III: Distributing MapReduce tasks
```
source setenv
cd src/mapreduce
go test -run TestParallel
```
```
018/09/16 14:38:50 rpc.Register: method "CleanupFiles" has 1 input parameters; needs exactly three
2018/09/16 14:38:50 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:38:50 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:38:50 rpc.Register: method "Wait" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36159-master: Starting Map/Reduce task test
2018/09/16 14:38:50 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
Schedule: 20 mapPhase tasks (10 I/Os)
2018/09/16 14:38:50 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:38:50 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:38:50 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36159-worker1: given mapPhase task #0 on file 824-mrinput-0.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: given mapPhase task #1 on file 824-mrinput-1.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #1 done
/var/tmp/824-501/mr36159-worker1: mapPhase task #0 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #3 on file 824-mrinput-3.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: given mapPhase task #2 on file 824-mrinput-2.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #2 done
/var/tmp/824-501/mr36159-worker0: mapPhase task #3 done
/var/tmp/824-501/mr36159-worker1: given mapPhase task #4 on file 824-mrinput-4.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: given mapPhase task #5 on file 824-mrinput-5.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #4 done
/var/tmp/824-501/mr36159-worker0: mapPhase task #5 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #6 on file 824-mrinput-6.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: given mapPhase task #7 on file 824-mrinput-7.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #6 done
/var/tmp/824-501/mr36159-worker1: mapPhase task #7 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #8 on file 824-mrinput-8.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: given mapPhase task #9 on file 824-mrinput-9.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #9 done
/var/tmp/824-501/mr36159-worker0: mapPhase task #8 done
/var/tmp/824-501/mr36159-worker1: given mapPhase task #11 on file 824-mrinput-11.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: given mapPhase task #10 on file 824-mrinput-10.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #11 done
/var/tmp/824-501/mr36159-worker0: mapPhase task #10 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #13 on file 824-mrinput-13.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: given mapPhase task #12 on file 824-mrinput-12.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #13 done
/var/tmp/824-501/mr36159-worker1: mapPhase task #12 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #14 on file 824-mrinput-14.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: given mapPhase task #15 on file 824-mrinput-15.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #14 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #16 on file 824-mrinput-16.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #15 done
/var/tmp/824-501/mr36159-worker1: given mapPhase task #17 on file 824-mrinput-17.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #16 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #18 on file 824-mrinput-18.txt (nios: 10)
/var/tmp/824-501/mr36159-worker1: mapPhase task #17 done
/var/tmp/824-501/mr36159-worker1: given mapPhase task #19 on file 824-mrinput-19.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #18 done
/var/tmp/824-501/mr36159-worker1: mapPhase task #19 done
Schedule: mapPhase done
Schedule: 10 reducePhase tasks (20 I/Os)
/var/tmp/824-501/mr36159-worker0: given reducePhase task #1 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: given reducePhase task #0 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #1 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #2 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #0 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #3 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #2 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #4 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #3 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #5 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #4 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #6 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #5 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #7 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #6 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #8 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #7 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #9 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #8 done
/var/tmp/824-501/mr36159-worker1: reducePhase task #9 done
Schedule: reducePhase done
Merge: read mrtmp.test-res-0
Merge: read mrtmp.test-res-1
Merge: read mrtmp.test-res-2
Merge: read mrtmp.test-res-3
Merge: read mrtmp.test-res-4
Merge: read mrtmp.test-res-5
Merge: read mrtmp.test-res-6
Merge: read mrtmp.test-res-7
Merge: read mrtmp.test-res-8
Merge: read mrtmp.test-res-9
/var/tmp/824-501/mr36159-master: Map/Reduce task completed
2018/09/16 14:38:51 rpc.Register: method "CleanupFiles" has 1 input parameters; needs exactly three
2018/09/16 14:38:51 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:38:51 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:38:51 rpc.Register: method "Wait" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36159-master: Starting Map/Reduce task test
2018/09/16 14:38:51 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
Schedule: 20 mapPhase tasks (10 I/Os)
2018/09/16 14:38:51 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:38:51 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:38:51 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36159-worker1: given mapPhase task #0 on file 824-mrinput-0.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: given mapPhase task #1 on file 824-mrinput-1.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #1 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #2 on file 824-mrinput-2.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #2 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #3 on file 824-mrinput-3.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #3 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #4 on file 824-mrinput-4.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #4 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #5 on file 824-mrinput-5.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #5 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #6 on file 824-mrinput-6.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #6 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #7 on file 824-mrinput-7.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #7 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #8 on file 824-mrinput-8.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #8 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #9 on file 824-mrinput-9.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #9 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #10 on file 824-mrinput-10.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #10 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #11 on file 824-mrinput-11.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #11 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #12 on file 824-mrinput-12.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #12 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #13 on file 824-mrinput-13.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #13 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #14 on file 824-mrinput-14.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #14 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #15 on file 824-mrinput-15.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #15 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #16 on file 824-mrinput-16.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #16 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #17 on file 824-mrinput-17.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #17 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #18 on file 824-mrinput-18.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #18 done
/var/tmp/824-501/mr36159-worker0: given mapPhase task #19 on file 824-mrinput-19.txt (nios: 10)
/var/tmp/824-501/mr36159-worker0: mapPhase task #19 done
/var/tmp/824-501/mr36159-worker1: mapPhase task #0 done
Schedule: mapPhase done
Schedule: 10 reducePhase tasks (20 I/Os)
/var/tmp/824-501/mr36159-worker0: given reducePhase task #1 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: given reducePhase task #0 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #1 done
/var/tmp/824-501/mr36159-worker1: reducePhase task #0 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #2 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: given reducePhase task #3 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #3 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #4 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #2 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #5 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #4 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #6 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #5 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #7 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #6 done
/var/tmp/824-501/mr36159-worker1: given reducePhase task #8 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker0: reducePhase task #7 done
/var/tmp/824-501/mr36159-worker0: given reducePhase task #9 on file  (nios: 20)
/var/tmp/824-501/mr36159-worker1: reducePhase task #8 done
/var/tmp/824-501/mr36159-worker0: reducePhase task #9 done
Schedule: reducePhase done
Merge: read mrtmp.test-res-0
Merge: read mrtmp.test-res-1
Merge: read mrtmp.test-res-2
Merge: read mrtmp.test-res-3
Merge: read mrtmp.test-res-4
Merge: read mrtmp.test-res-5
Merge: read mrtmp.test-res-6
Merge: read mrtmp.test-res-7
Merge: read mrtmp.test-res-8
Merge: read mrtmp.test-res-9
/var/tmp/824-501/mr36159-master: Map/Reduce task completed
PASS
ok      mapreduce       3.713s
```

## Part IV: Handling worker failures
```
source setenv
cd src/mapreduce
go test -run Failure
```
```
2018/09/16 14:41:52 rpc.Register: method "CleanupFiles" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Wait" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36485-master: Starting Map/Reduce task test
2018/09/16 14:41:52 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:52 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
Schedule: 20 mapPhase tasks (10 I/Os)
/var/tmp/824-501/mr36485-worker1: given mapPhase task #1 on file 824-mrinput-1.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #0 on file 824-mrinput-0.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #1 done
/var/tmp/824-501/mr36485-worker0: mapPhase task #0 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #2 on file 824-mrinput-2.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #3 on file 824-mrinput-3.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #3 done
/var/tmp/824-501/mr36485-worker1: mapPhase task #2 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #4 on file 824-mrinput-4.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #5 on file 824-mrinput-5.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #5 done
/var/tmp/824-501/mr36485-worker1: mapPhase task #4 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #6 on file 824-mrinput-6.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: given mapPhase task #7 on file 824-mrinput-7.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #6 done
/var/tmp/824-501/mr36485-worker1: mapPhase task #7 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #8 on file 824-mrinput-8.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: given mapPhase task #9 on file 824-mrinput-9.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #8 done
/var/tmp/824-501/mr36485-worker1: mapPhase task #9 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #10 on file 824-mrinput-10.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: given mapPhase task #11 on file 824-mrinput-11.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #11 done
/var/tmp/824-501/mr36485-worker0: mapPhase task #10 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #12 on file 824-mrinput-12.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #13 on file 824-mrinput-13.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #12 done
/var/tmp/824-501/mr36485-worker0: mapPhase task #13 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #15 on file 824-mrinput-15.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: given mapPhase task #14 on file 824-mrinput-14.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #14 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #16 on file 824-mrinput-16.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #15 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #17 on file 824-mrinput-17.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #16 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #18 on file 824-mrinput-18.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #17 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #19 on file 824-mrinput-19.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #18 done
/var/tmp/824-501/mr36485-worker0: mapPhase task #19 done
Schedule: mapPhase done
Schedule: 10 reducePhase tasks (20 I/Os)
/var/tmp/824-501/mr36485-worker1: given reducePhase task #1 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #1 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #2 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #2 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #3 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #3 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #4 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #4 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #5 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #5 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #6 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #6 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #7 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #7 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #8 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #8 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #9 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #9 done
/var/tmp/824-501/mr36485-worker1: given reducePhase task #0 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker1: reducePhase task #0 done
Schedule: reducePhase done
Master: RPC /var/tmp/824-501/mr36485-worker0 shutdown error
Merge: read mrtmp.test-res-0
Merge: read mrtmp.test-res-1
Merge: read mrtmp.test-res-2
Merge: read mrtmp.test-res-3
Merge: read mrtmp.test-res-4
Merge: read mrtmp.test-res-5
Merge: read mrtmp.test-res-6
Merge: read mrtmp.test-res-7
Merge: read mrtmp.test-res-8
Merge: read mrtmp.test-res-9
/var/tmp/824-501/mr36485-master: Map/Reduce task completed
2018/09/16 14:41:54 rpc.Register: method "CleanupFiles" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Wait" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36485-master: Starting Map/Reduce task test
Schedule: 20 mapPhase tasks (10 I/Os)
2018/09/16 14:41:54 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:54 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36485-worker1: given mapPhase task #0 on file 824-mrinput-0.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #1 on file 824-mrinput-1.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #0 done
/var/tmp/824-501/mr36485-worker0: mapPhase task #1 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #2 on file 824-mrinput-2.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: given mapPhase task #3 on file 824-mrinput-3.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #3 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #4 on file 824-mrinput-4.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #2 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #5 on file 824-mrinput-5.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #4 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #6 on file 824-mrinput-6.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #5 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #7 on file 824-mrinput-7.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #7 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #8 on file 824-mrinput-8.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #6 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #9 on file 824-mrinput-9.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #9 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #10 on file 824-mrinput-10.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #8 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #11 on file 824-mrinput-11.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #10 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #12 on file 824-mrinput-12.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #11 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #13 on file 824-mrinput-13.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #12 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #14 on file 824-mrinput-14.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #13 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #15 on file 824-mrinput-15.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #14 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #16 on file 824-mrinput-16.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #15 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #17 on file 824-mrinput-17.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #16 done
/var/tmp/824-501/mr36485-worker0: given mapPhase task #18 on file 824-mrinput-18.txt (nios: 10)
/var/tmp/824-501/mr36485-worker1: mapPhase task #17 done
/var/tmp/824-501/mr36485-worker1: given mapPhase task #19 on file 824-mrinput-19.txt (nios: 10)
/var/tmp/824-501/mr36485-worker0: mapPhase task #18 done
/var/tmp/824-501/mr36485-worker1: mapPhase task #19 done
Schedule: mapPhase done
Schedule: 10 reducePhase tasks (20 I/Os)
2018/09/16 14:41:55 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:55 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
2018/09/16 14:41:55 rpc.Register: method "Lock" has 1 input parameters; needs exactly three
2018/09/16 14:41:55 rpc.Register: method "Unlock" has 1 input parameters; needs exactly three
/var/tmp/824-501/mr36485-worker2: given reducePhase task #3 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker3: given reducePhase task #2 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker2: reducePhase task #3 done
/var/tmp/824-501/mr36485-worker2: given reducePhase task #4 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker3: reducePhase task #2 done
/var/tmp/824-501/mr36485-worker3: given reducePhase task #5 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker2: reducePhase task #4 done
/var/tmp/824-501/mr36485-worker2: given reducePhase task #6 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker3: reducePhase task #5 done
/var/tmp/824-501/mr36485-worker3: given reducePhase task #7 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker2: reducePhase task #6 done
/var/tmp/824-501/mr36485-worker2: given reducePhase task #8 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker3: reducePhase task #7 done
/var/tmp/824-501/mr36485-worker3: given reducePhase task #9 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker2: reducePhase task #8 done
/var/tmp/824-501/mr36485-worker2: given reducePhase task #0 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker3: reducePhase task #9 done
/var/tmp/824-501/mr36485-worker3: given reducePhase task #1 on file  (nios: 20)
/var/tmp/824-501/mr36485-worker2: reducePhase task #0 done
/var/tmp/824-501/mr36485-worker3: reducePhase task #1 done
Schedule: reducePhase done
Master: RPC /var/tmp/824-501/mr36485-worker0 shutdown error
Master: RPC /var/tmp/824-501/mr36485-worker1 shutdown error
Merge: read mrtmp.test-res-0
Merge: read mrtmp.test-res-1
Merge: read mrtmp.test-res-2
Merge: read mrtmp.test-res-3
Merge: read mrtmp.test-res-4
Merge: read mrtmp.test-res-5
Merge: read mrtmp.test-res-6
Merge: read mrtmp.test-res-7
Merge: read mrtmp.test-res-8
Merge: read mrtmp.test-res-9
/var/tmp/824-501/mr36485-master: Map/Reduce task completed
PASS
ok      mapreduce       4.113s
```

## Part V: Inverted index generation
```
source setenv
cd src/main
go run ii.go master sequential pg-*.txt
```
```
master: Starting Map/Reduce task iiseq
Merge: read mrtmp.iiseq-res-0
Merge: read mrtmp.iiseq-res-1
Merge: read mrtmp.iiseq-res-2
master: Map/Reduce task completed
```
```
head -n5 mrtmp.iiseq
```
```
A: 8 pg-being_ernest.txt,pg-dorian_gray.txt,pg-frankenstein.txt,pg-grimm.txt,pg-huckleberry_finn.txt,pg-metamorphosis.txt,pg-sherlock_holmes.txt,pg-tom_sawyer.txt
ABOUT: 1 pg-tom_sawyer.txt
ACT: 1 pg-being_ernest.txt
ACTRESS: 1 pg-dorian_gray.txt
ACTUAL: 8 pg-being_ernest.txt,pg-dorian_gray.txt,pg-frankenstein.txt,pg-grimm.txt,pg-huckleberry_finn.txt,pg-metamorphosis.txt,pg-sherlock_holmes.txt,pg-tom_sawyer.txt
```

## Running all tests
```
source setenv
cd src/main
bash ./test-mr.sh
```
```
==> Part I
ok      mapreduce       1.934s

==> Part II
Passed test

==> Part III
ok      mapreduce       3.412s

==> Part IV
ok      mapreduce       4.158s

==> Part V (inverted index)
Passed test
```