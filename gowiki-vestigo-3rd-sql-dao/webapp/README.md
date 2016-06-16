A DAO application example with sql persistence

# Installing

up the postgresql database

    $ docker-compose up

build and run the application

	$ go get
	$ ./build.sh
	$ ./build/wiki
	
# Benchmark

runnet at a I7, 16GB machine(I will update to the common machine)

    $ siege -c400 -t10s http://localhost:8989/edit/elvis-23
    Transactions:		        7590 hits
    Availability:		      100.00 %
    Elapsed time:		        9.57 secs
    Data transferred:	        1.47 MB
    Response time:		        0.01 secs
    Transaction rate:	      793.10 trans/sec
    Throughput:		        0.15 MB/sec
    Concurrency:		        6.04
    Successful transactions:        7590
    Failed transactions:	           0
    Longest transaction:	        0.13
    Shortest transaction:	        0.00
