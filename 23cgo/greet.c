#include <stdio.h>
#include "greeter.h"

// void greet() {
//     printf("Hello\n");
// }

// void greetString(char *a){
//     printf("Hello  %s\n", a);
// }

void greetStruct(struct Greetee *g){
    printf("Hello %s, from %d\n", g->name, g->year);
}