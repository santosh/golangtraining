package main

import "fmt"

func main() {
	myFriendsName := "Mar"

    switch {
    case len(myFriendsName) == 2:
        fmt.Println("Wassaup my friend with name of length 2")
    case myFriendsName == "Tim":
        fmt.Println("Wassap Tim")
    case myFriendsName == "Jenny":
        fmt.Println("Wassap Jenny")
    case myFriendsName == "Marcus", myFriendsName == "Medhi":
        fmt.Println("Your name is either Marcus or Medhi")
    case myFriendsName == "Julian":
        fmt.Println("Wassap Julian")
    case myFriendsName == "Sushant":
        fmt.Println("Wassap Sushant")
    default:
        fmt.Println("nothing matches")
    }
}
