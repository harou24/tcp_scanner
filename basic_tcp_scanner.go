package main

import(
    "fmt"
    "net"
)

func main(){
    for i := 1; i <= 1024; i++ {
        addr := fmt.Sprintf("scanme.nmap.org:%d", i)

        // the Dial function connects to a server.
        conn, err := net.Dial("tcp", addr)

        if err != nil {
            // port is closed or filtered.
            continue
        }
        conn.Close()
        fmt.Printf("%d open\n", i)
    }
}
