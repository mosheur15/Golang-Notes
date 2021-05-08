package main

import (
    "bufio";
    "fmt";
    "os";
    "strings";
    "math/rand";
    "strconv"
    //"reflect";
)

func input(msg ...string) string {
    // Takes user input
    if len(msg) > 0 {
        fmt.Print(msg[0])
    }
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text()
    return str
}

func main() {
    str := func(ans string) string {
        if len(ans) > 0 {
            ans = strings.ToLower(strings.TrimSpace(ans))
        }
        return ans
    }

    game := func() {
        for true {
            num := rand.Intn(100)

            for true {
                ans := input("> guess the number: ")
                guess,_ := strconv.Atoi(ans)
                if num > guess {
                    fmt.Println("> too low.")
                } else if num < guess {
                    fmt.Println("> too high.")
                } else {
                    fmt.Println("> congratulations you have guessed the correct number!")
                    break
                }
            }
        }
    }

    func() {
        fmt.Println("<--------[ guess the number ]------->")
        fmt.Println("")
        ans := str(input("> do you want to play the game? (y/n): "))

        if ans == "y" {
            game()
        }
    }()

}