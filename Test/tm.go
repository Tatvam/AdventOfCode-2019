package main

import (
    tm "github.com/buger/goterm"
    "time"
)

func main() {
    tm.Clear() // Clear current screen

    for {
        // By moving cursor to top-left position we ensure that console output
        // will be overwritten each time, instead of adding new.
        //tm.MoveCursor(1,1)
        i := 0
        for i < 100 {
            tm.MoveCursor(i%10 + 1, 1 )
            tm.Printf("%d",i)
            time.Sleep(1 * time.Second)
            i++
            tm.Flush() // Call it every time at the end of rendering
        }
        time.Sleep(time.Second)
    }
}