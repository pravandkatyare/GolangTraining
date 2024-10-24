
#ifndef _GREETER_H_
#define _GREETER_H_

struct Greetee {
    const char *name;
    int year;
};

// void greet();
void greetStruct(struct Greetee *g);
// void greetString(char *a);
#endif