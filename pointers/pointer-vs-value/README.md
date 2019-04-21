* When you pass as value, golang will copy this parameter to pass to the function
* When you pass as pointer, golang simply pass the address of the original variable
* With pointers the receiver function can change the original parameter value
* With pointers you increase program performance and memory utilization
* It happen even with structs