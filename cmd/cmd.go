package cmd

import "fmt"

func printrow(row []bool) {
    rowstr := ""
    for _, elm := range row {
       if elm {
        rowstr += "1"
       } else {
        rowstr += "0"
       } 
    }
    fmt.Println(rowstr)
}
