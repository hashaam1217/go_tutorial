package main
import ("fmt")

func main() {
    var a int = 34
    var array = [3]string{"Hello", "World", "!"}

    fmt.Printf("%s\n", a)
    fmt.Printf("%8s\n", a)
    fmt.Printf("%-8s\n", a)
    fmt.Printf("% x\n", a)
    fmt.Println(array)
}
