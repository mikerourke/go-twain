# go-twain

Go library mapping for the TWAIN protocol.

This library provides a means of writing TWAIN applications that communicate via the DSM with Go (instead of C/C++).
My future goal is to port the DSM to Go as well. 
This would avoid the need to link to a DLL/shared object, enabling you to build a Go app that communicates with a datasource (`.ds`) file directly.

## Caveats

The amount of testing I've done is limited. I literally just went through the `twain.h` file line by line and created equivalent types/structs/constants in Go.

The following issues are bound to come up, most of which will be rectified with a DSM -> Go port:

1. Since Go doesn't support `union`, I'm using a byte array (per the cgo documentation)
2. I didn't port the entry points/callbacks (e.g. `DSM_Entry` and `DS_Entry`) 

## License

Licensed under the [MIT license](https://opensource.org/licenses/MIT)
