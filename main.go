package main

import (
    "fmt"
    "neutrinos"
    "os"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./206neutrinos n a h sd")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tn\t\tnumber of values")
    fmt.Println("\ta\t\tarithmetic mean")
    fmt.Println("\th\t\tharmonic mean")
    fmt.Println("\tsd\t\tstandard deviation")
}

func main() {
    if neutrinos.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !neutrinos.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    neutrinos.Neutrinos()
    os.Exit(0)
}