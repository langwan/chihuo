package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void echoInt(int n) {
    printf("c print: %d\n", n);
}
void echoStr(char* str) {
    printf("c print: %s\n", str);
}
*/
import "C"
import "unsafe"

func basic() {
	C.echoInt(C.int(100))
	str := C.CString("chihuo")
	defer C.free(unsafe.Pointer(str))
	C.echoStr(str)
}
