package main

import "code.google.com/p/go-tour/pic"

func picc(dx, dy int) [][]uint8 {
    // slicing using make() funtion
    pic := make([][]uint8, dy)

    for i := 0; i < dy; i++ {
        // making sclice of lenght dx
        pic[i] = make([]uint8, dx)
    }
    for y := range pic {
        for x := range pic[y]{
            pic[y][x] = uint8(x^y)
        }
    }
    return pic
}   


func main(){
    picc.Show(picc)
}
