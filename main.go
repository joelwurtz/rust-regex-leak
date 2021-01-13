package main

/*
#cgo LDFLAGS: -lrustregexleakrepro -Ltarget/x86_64-unknown-linux-musl/debug
#include <stdlib.h>

typedef struct {} Regex;

#include "target/rust-regex-leak-repro.h"

*/
import "C"
import "time"

func main() {
	var regex *C.Regex = C.create_regex()

	for i := 0; i < 10000000; i++ {
		// Uncomment this line to avoid leaking
		//C.regex_match_leak(regex)

		go func() {
			C.regex_match_leak(regex)
		}()
	}

	time.Sleep(60 * time.Second)
}