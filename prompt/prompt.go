package main

import "fmt"

func main() {
    var running bool = true
    var option int = 0

    for running {
        menu()
        fmt.Scanf("%d\n", &option)

        if option == 0 {
            running = false
        }
    }
}

func menu() {
    fmt.Println("# MONEY LOGS")
    fmt.Println("0. EXIT")
    fmt.Println("1. ADD ENTRY")
    fmt.Println("2. SHOW ENTRIES")
    fmt.Println("Choose an option:")
}
