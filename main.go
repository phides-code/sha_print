/*write a program that prints the SHA256 hash of its standard input by default, but supports a command line flag to print the sha384 or sha512 hash instead. */

package main

import (
    "fmt"
    "flag"
    "crypto/sha256"
    "crypto/sha512"
    "os"
)

var f384 = flag.Bool("384", false, "print the sha384 hash of the argument")
var f512 = flag.Bool("512", false, "print the sha512 hash of the argument")

func main() {
    flag.Parse()    
    printHash(os.Args, f384, f512)
}

func printHash(arg []string, f384, f512 *bool ) {
    if *f384 {
        fmt.Printf("sha384: %x\n", sha512.Sum384([]byte(arg[2])))
        return
    }
    
    if *f512 {
        fmt.Printf("sha512: %x\n", sha512.Sum512([]byte(arg[2])))
        return
    }
    
    fmt.Printf("sha256: %x\n", sha256.Sum256([]byte(arg[1])))
    return
}