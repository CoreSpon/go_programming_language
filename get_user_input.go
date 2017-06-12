
package main
import "fmt"

func main() {
    var in_put_name string
    str1 := "What is your name?"
    fmt.Println(str1)
    fmt.Scanf("%s", &in_put_name)
    fmt.Println("Hello if you spelled your name right then hello", in_put_name)
}
