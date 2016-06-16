# Introduction
This test objective is compare 
The summary of the result is to compare [vestigo](https://github.com/husobee/vestigo) perfomance vs [google native handler (with a simple adapter)](https://golang.org/doc/articles/wiki/#tmp_13) at the same application

# Summary 

## CPU Performance
Every app was tested 3 times alternately with the follow result (details on benchmark section)

	VESTIGO = 589.42 trans/sec + 600.20 trans/sec + 594.51 trans/sec = 594,71 trans/sec
	NATIVE  = 591.66 trans/sec + 606.87 trans/sec + 597.61 trans/sec = 598,71 trans/sec

## Memory consumition

	VESTIGO app starts with `1.0M` and take `4.8M` at the peak

	NATIVE app starts with the same `1.0M` and take the same `4.8` at the peak

	source: `ps aux`

Conclusion

Native is `0,67%` fastest than Vestigo framework so if you want to use Vestigo perfomance will not be a problem

# Hardware details
This benchmark runned at:

	Linux x64 4.4.0-22-generic
	Ubuntu 16.04 
	8 GB RAM
	CPU
		product: Intel(R) Core(TM) i3-3227U CPU @ 1.90GHz
		size: 1065MHz
		capacity: 1900MHz
		width: 64 bits

# The benchmark test 

The test was start the two apps alternately and run the follow command

	siege -c300 -t30s http://localhost:8080/edit/

the result was the report below

# Vestigo

### 1
	Transactions:		       17659 hits
	Availability:		      100.00 %
	Elapsed time:		       29.96 secs
	Data transferred:	        3.44 MB
	Response time:		        0.00 secs
	Transaction rate:	      589.42 trans/sec
	Throughput:		        0.11 MB/sec
	Concurrency:		        2.43
	Successful transactions:       17659
	Failed transactions:	           0
	Longest transaction:	        1.00
	Shortest transaction:	        0.00

### 2
	Transactions:		       17628 hits
	Availability:		      100.00 %
	Elapsed time:		       29.37 secs
	Data transferred:	        3.43 MB
	Response time:		        0.00 secs
	Transaction rate:	      600.20 trans/sec
	Throughput:		        0.12 MB/sec
	Concurrency:		        1.83
	Successful transactions:       17628
	Failed transactions:	           0
	Longest transaction:	        0.07
	Shortest transaction:	        0.00

### 3
	Transactions:		       17657 hits
	Availability:		      100.00 %
	Elapsed time:		       29.70 secs
	Data transferred:	        3.44 MB
	Response time:		        0.00 secs
	Transaction rate:	      594.51 trans/sec
	Throughput:		        0.12 MB/sec
	Concurrency:		        2.40
	Successful transactions:       17657
	Failed transactions:	           0
	Longest transaction:	        1.00
	Shortest transaction:	        0.00

# Native

### 1
	Transactions:		       17300 hits
	Availability:		      100.00 %
	Elapsed time:		       29.24 secs
	Data transferred:	        3.37 MB
	Response time:		        0.01 secs
	Transaction rate:	      591.66 trans/sec
	Throughput:		        0.12 MB/sec
	Concurrency:		        3.26
	Successful transactions:       17300
	Failed transactions:	           0
	Longest transaction:	        3.01
	Shortest transaction:	        0.00

### 2

	Transactions:		       17757 hits
	Availability:		      100.00 %
	Elapsed time:		       29.26 secs
	Data transferred:	        3.45 MB
	Response time:		        0.00 secs
	Transaction rate:	      606.87 trans/sec
	Throughput:		        0.12 MB/sec
	Concurrency:		        1.75
	Successful transactions:       17757
	Failed transactions:	           0
	Longest transaction:	        0.07
	Shortest transaction:	        0.00

### 3
	Transactions:		       17779 hits
	Availability:		      100.00 %
	Elapsed time:		       29.75 secs
	Data transferred:	        3.46 MB
	Response time:		        0.00 secs
	Transaction rate:	      597.61 trans/sec
	Throughput:		        0.12 MB/sec
	Concurrency:		        1.89
	Successful transactions:       17779
	Failed transactions:	           0
	Longest transaction:	        0.09
	Shortest transaction:	        0.00
