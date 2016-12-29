package main

import "fmt"
import "github.com/ishiikurisu/moneylog"
import "io/ioutil"

func main() {
    var running bool = true
    var option int = 0
    var log *moneylog.Log = nil

    for running {
        menu()
        fmt.Scanln(&option)

        if option == 0 {
            running = false
        } else if option == 1 {
            log = createEntry(log)
        } else if option == 2 {
            fmt.Printf("%s\n", log.ToString())
        } else if option == 3 {
            fmt.Printf("Your balance is $%.2f\n", log.CalculateBalance())
        } else if option == 4 {
            log = load()
        }
    }
}

func menu() {
    fmt.Println("# MONEY LOGS")
    fmt.Println("0. EXIT")
    fmt.Println("1. ADD ENTRY")
    fmt.Println("2. SHOW ENTRIES")
    fmt.Println("3. GET BALANCE")
    fmt.Println("Choose an option:")
}

func createEntry(log *moneylog.Log) *moneylog.Log {
    var description string
    var value float64

    fmt.Println("Description:")
    fmt.Scanf("%s\n", &description)
    fmt.Println("Value:")
    fmt.Scanln(&value)

    if log == nil {
        l := moneylog.StartLog(description, value)
        log = &l
    } else {
        log.Add(description, value)
    }

    return log
}

func load() *moneylog.Log {
    var path string

    fmt.Println("Write path to file:")
    fmt.Scanln(&path)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", string(content))
    l := moneylog.LogFromString(string(content))
    return &l
}
