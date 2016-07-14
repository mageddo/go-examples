# Introduction
This example have a shared array that is modified by two goroutines at the "same time", 
using lock resource the application prevents that the `find` method reads the `ids` while remove method is removing 
then `find` will not count a wrong array size because they are synchronized   