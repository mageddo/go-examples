# Problem

sql.Scan have a problem, if the column at the database is null then you will got a error at runtime, 
there are some solutions for this problem:
 
* Use pointer variables inside struct 
* Not allow nullable columns
* Treat nullable columns at insert

At these subpackages are the implementation of some solutions, good luck