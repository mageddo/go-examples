#include <stdio.h>

struct Books {
   char  title[50];
   char  author[50];
   char  subject[100];
   int   book_id;
} book; 

int main()
{
    struct Books Book1;
    printf("book1=%p\n", &Book1);
    increaseId(&Book1);
    printf("book1=%p\n", &Book1);

    return 0;
}

void increaseId(struct Books *book){
    book->book_id++;
    printf("book=%p\n", &book);
}
