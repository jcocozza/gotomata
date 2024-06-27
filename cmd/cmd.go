package cmd

import "fmt"

func Printrow(row []bool) {
    rowstr := ""
    for _, elm := range row {
       if elm {
        rowstr += "█"
       } else {
        rowstr += "░"
       } 
    }
    fmt.Println(rowstr)
}
