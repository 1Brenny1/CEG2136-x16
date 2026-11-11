# TinyC
Because C is a bit complex and is expected to run on more powerful hardware, I am going to make my own sublanguage which I am gonna call `TinyC` or `.tc` since it will be a subset of features from C but dumbed down so I can make a compiler relatively fast to compile its code to assembly for Mano's Basic Computer. In this document, I will outline all documentation relevant to TinyC.

Due to limitations in Mano's Basic Computer, I will be stripping types and leave everything as a uint16 which can be setup to act as though it is a int16. Furthurmore, any variable will also be able to be treated as a pointer, so be careful.

Program entry point:
```mc
#include <std.tc> // Includes will just copy and paste the content of the file into this file

// Program Entry Point
func main() {
    
}
```