# Reference Links
* http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
* https://golang.org/doc/faq#pointer_to_interface

# Explanation
Pointers to interfaces are allowed but no needle most of the time, 
if you need to work with pointers specify the pointer in the struct, not on the interface, understood what you have
specified on the struct

# Why the Person pointer is differente on every method call?
Every time that you call a method golang make a copy of the passed pointer, so the two pointers point to the same struct but are two different pointers, it works like in c

