package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof" }

type Robot struct{}
func (r Robot) Speak() string { return "Beep" }

func Announce(s Speaker) {
    fmt.Println(s.Speak())
}
func main(){
// Both Dog and Robot satisfy Speaker — no declaration required.
Announce(Dog{})
Announce(Robot{})
 }
