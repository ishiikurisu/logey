package main

import "fmt"
import "github.com/ishiikurisu/moneylog"

func main() {
    var running bool = true
    var option int = 0

    for running {
        menu()
        fmt.Scanln(&option)

        if option == 0 {
            running = false
        } else if option == 1 {
            createEntry()
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

func createEntry() {
    var description string
    var value float64

    fmt.Println("Description:")
    fmt.Scanln(&description)
    fmt.Println("Value:")
    fmt.Scanln(&value)
    entry := moneylog.NewEntry(description, value)
    fmt.Printf("%#v\n", entry)
}
