package main
import "fmt"


func main (){
    grades := make(map[string]float32)

    grades["Ken"] = 42
    grades["Jess"] = 92
    grades["Sam"] = 71
    
    for k, v := range grades{
        fmt.Println(k,":" , v)
    }
}




