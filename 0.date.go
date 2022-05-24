package main
 
import "fmt"
import "time"
 
func main() {
    dt := time.Now()
    //Format MM-DD-YYYY
    fmt.Println(dt.AddDate(0,0,7).Format("01-02-2006"))
    // 04-11-2022

    fmt.Println(dt.AddDate(0,0,7).Format("20060102"))
    // 04-11-2022


    // fmt.Println(dt.Format("01-02-2006"))
}