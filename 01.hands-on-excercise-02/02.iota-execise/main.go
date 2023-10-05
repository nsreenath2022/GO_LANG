package main

import (
    "fmt"
)

type Bytesize int

const (
    _  = iota
    KB Bytesize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
)
func main(){
    fmt.Printf("%d  \t\t %b\n", KB, KB)
    fmt.Printf("%d  \t\t %b\n", MB, MB)
    fmt.Printf("%d  \t\t %b\n", GB, GB)
    fmt.Printf("%d  \t\t %b\n", TB, TB)
    fmt.Printf("%d  \t\t %b\n", PB, PB)
}