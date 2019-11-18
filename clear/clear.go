package clear

import (
    "os"
    "os/exec"
)


func CallClear() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}

// func main() {
//     fmt.Println("I will clean the screen in 2 seconds!")
//     time.Sleep(2 * time.Second)
//     CallClear()
//     fmt.Println("I'm alone...")
// }