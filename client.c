#include <stdio.h>
#include "cnats.h"

int main() {
    printf("Demo of invoking the cnats lib from C:\n");    
    Pub("nats://localhost:4555", "test.subject", "hello world");
}