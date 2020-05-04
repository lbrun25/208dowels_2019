package main

import (
    "fmt"
	"dowels"
    "os"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./208dowels O0 O1 O2 O3 O4 O5 O6 O7 O8")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tOi\tsize of the observed class")
}

func main() {
    if dowels.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !dowels.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    dowels.Dowels()
    os.Exit(0)
}