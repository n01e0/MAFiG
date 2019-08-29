#include <stdio.h>
#include <stdlib.h>

int inpt(void) {
int c = getchar();
if (c == EOF) {
exit(0);
}
return c;
}

int main() {
char vmem[32767];
char *ptr = vmem;
(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
while (*ptr) {
(*ptr)--;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;
ptr++;
ptr++;
ptr++;
ptr++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
}

ptr++;(*ptr)--;(*ptr)--;
ptr++;(*ptr)--;(*ptr)--;(*ptr)--;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)--;(*ptr)--;(*ptr)--;
ptr++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;
ptr++;
ptr++;
ptr++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
while (*ptr) {

ptr++;
ptr++;
ptr++;(*ptr)++;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;
while (*ptr) {
(*ptr)--;
ptr++;
ptr++;(*ptr)++;
ptr++;(*ptr)++;
ptr--;
ptr--;
ptr--;
}

ptr++;
ptr++;
ptr++;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;
ptr--;(*ptr)++;
ptr++;
ptr++;
ptr++;
}
(*ptr)++;
ptr--;
while (*ptr) {

while (*ptr) {
(*ptr)--;
}

ptr++;(*ptr)--;
ptr--;
ptr--;
while (*ptr) {
(*ptr)--;
ptr++;(*ptr)++;
ptr++;(*ptr)++;
ptr--;
ptr--;
}

ptr++;
ptr++;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;(*ptr)++;
ptr++;
ptr++;
}
(*ptr)++;
ptr--;
while (*ptr) {

while (*ptr) {
(*ptr)--;
}

ptr++;(*ptr)--;
ptr--;
ptr--;
ptr--;(*ptr)++;
ptr++;(*ptr)--;
ptr++;
}

ptr++;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;
ptr--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;(*ptr)--;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;
ptr++;
ptr++;
ptr++;
ptr++;(*ptr)++;
ptr--;
ptr--;
ptr--;
}

ptr--;
}

ptr++;
while (*ptr) {
(*ptr)--;
ptr--;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
while (*ptr) {

ptr--;
ptr--;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;
ptr++;(*ptr)--;
}

ptr--;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr++;
ptr++;
}

ptr++;
ptr++;
ptr++;
}

ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
while (*ptr) {

ptr--;
ptr--;
ptr--;
ptr--;
}

ptr++;(*ptr)--;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;(*ptr)++;
ptr++;(*ptr)++;
ptr++;
}

ptr--;
while (*ptr) {
(*ptr)--;
ptr++;(*ptr)++;
ptr--;
}
(*ptr)++;
ptr--;
while (*ptr) {

while (*ptr) {
(*ptr)--;
}

ptr++;(*ptr)--;
ptr--;
}

ptr++;
while (*ptr) {
(*ptr)--;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
putchar(*ptr);

ptr++;
putchar(*ptr);

ptr++;
ptr++;
ptr++;
putchar(*ptr);

putchar(*ptr);

ptr++;
ptr++;(*ptr)++;
ptr++;
ptr++;
}

ptr++;
ptr++;(*ptr)--;
while (*ptr) {
(*ptr)--;
ptr--;
ptr--;
ptr--;(*ptr)++;
ptr++;(*ptr)++;
ptr++;
ptr++;
}

ptr--;
ptr--;
while (*ptr) {
(*ptr)--;
ptr++;
ptr++;(*ptr)++;
ptr--;
ptr--;
}
(*ptr)++;
ptr--;
while (*ptr) {

while (*ptr) {
(*ptr)--;
}

ptr++;(*ptr)--;
ptr--;
}

ptr++;
while (*ptr) {
(*ptr)--;
ptr++;
ptr++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;(*ptr)++;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
ptr--;
putchar(*ptr);

ptr++;
putchar(*ptr);

ptr++;
putchar(*ptr);

putchar(*ptr);

ptr++;
ptr++;(*ptr)++;
ptr++;
ptr++;
}

ptr--;(*ptr)++;
ptr--;
while (*ptr) {

while (*ptr) {
(*ptr)--;
}

ptr++;(*ptr)--;
ptr--;
}

ptr++;
while (*ptr) {
(*ptr)--;
ptr++;
ptr++;
ptr++;
ptr++;
ptr++;
while (*ptr) {

ptr++;
ptr++;
ptr++;
ptr++;
}

ptr--;
ptr--;
ptr--;
ptr--;
while (*ptr) {

putchar(*ptr);

ptr--;
ptr--;
ptr--;
ptr--;
}

ptr--;
}

ptr--;
ptr--;
putchar(*ptr);

ptr++;
ptr++;
ptr++;
ptr++;
ptr++;
ptr++;(*ptr)--;
}
putchar(13);putchar(10);
return 0;
}
