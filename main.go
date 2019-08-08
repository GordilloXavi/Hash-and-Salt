package main

import (
    "fmt"
    "crypto/sha256"
    "os"
    "bufio"
    "strconv"
)

var last_id int

type user struct{
    username string
    password [32]byte
    id int
} 

var db map[string]user

func main() {
    reader := bufio.NewReader(os.Stdin)
    db = map[string]user{}

    for {
        fmt.Println("1: Sign in\n2: Log in")
        text, _ := reader.ReadString('\n')
        fmt.Println(text)

        if text == "1\n" {
            fmt.Println("Pick username")
            username, _ := reader.ReadString('\n')
            _, exists := db[username]
            if exists {
                fmt.Println("Username already exists")
                continue
            } 
            fmt.Println("Pick password")
            password, _ := reader.ReadString('\n')
            str := []byte(password)
            str = append(str, []byte(strconv.Itoa(last_id))...)
            u := user{username, sha256.Sum256(str), last_id}
            db[username] = u
            last_id ++
        } else if text == "2\n" {
            fmt.Println("Username: ")
            u, _ := reader.ReadString('\n')
            us, f := db[u]
            if !f {
                fmt.Println("User does not exist")
                continue
            } 
            fmt.Println("Password: ")
            p, _ := reader.ReadString('\n')
            str := []byte(p+strconv.Itoa(us.id))
            enc := sha256.Sum256(str)
            if enc == us.password {
                fmt.Println("Correct password!!!! ", us.password)
            } else {
                fmt.Println("Wrong password!!!!")
            } 

        } else {
            fmt.Println("not 1 or 2")
        }
    } 


} 
