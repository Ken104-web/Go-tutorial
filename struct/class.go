// go has no classes instead they use structs
// := (short-declare operator) gives you a way to declare a variable and assign its value in a single statement, without needing to specify the type
// used to create a collection of members of different data types, into a single variable. 
// While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

package main

import "fmt"

type car struct {
    gas_pedal uint16     //min: 0,      max: 65535
	brake_pedal uint16   //min: 0,      max: 65535
	steering_wheel int16 //min: -32768  max: 32768
	top_speed_kmh float64
}

func main (){
    a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
    fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
}

