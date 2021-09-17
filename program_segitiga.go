package main
import "fmt"

func main() {
    fmt.Printf("Hello World, Double Skill!!!\n")
    for brs:=1; brs<=30; brs++{
        //melakukan pengulangan bintang(*) sampai 30

        for spasi:=30 ; spasi>=brs; spasi--{
            fmt.Printf(" ");
        }

        for klm:=1; klm<=brs; klm++{
            fmt.Printf("*"); /

        }
        for klm:=(brs-1); klm>=1; klm--{
            fmt.Printf("*");

        }
        fmt.Printf("\n");
    }
}