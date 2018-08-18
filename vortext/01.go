package main

import "net"
import "fmt"

func main() {
    conn, _ := net.Dial("tcp", "vortex.labs.overthewire.org:5842")

    buff := make([]byte, 1024)
    n, _ := conn.Read(buff)

    numbers := ReadNumbers(buff)

    fmt.Printf("Received %d bytes: %#x\n", n, buff[:n])
    fmt.Printf("The 4 unsigned integers: %d\n", numbers[:4])

    sum := Sum64(numbers)

    fmt.Printf("Sum is: %d (%#x)\n", sum, sum)

    response := PrepareResponse64(sum)

    fmt.Printf("Sending result: %#x\n", response)
    conn.Write(response)

    n, _ = conn.Read(buff)
    fmt.Printf("Received %d bytes: %#x\n", n, buff[:n])
    fmt.Printf("Response as ASCII text: %s\n", buff)
}

func Sum64(numbers []uint32) uint64 {
    var result uint64;

    for i := 0; i < 4; i++ {
        result += uint64(numbers[i]);
    }

    return result
}

// convert to little-endian []byte
func PrepareResponse64(number uint64) []byte {
    response := make([]byte, 8)
    
    for i := 0; i < 8; i++ {
        response[i] = byte(number & uint64(0xff))
        number = number >> 8
    }

    return response
}

// interpret buffer as 4 unsigned integers in little-endian format
func ReadNumbers(buff []byte) []uint32 {
    numbers := make([]uint32, 4)

    // read the 16byte buffer into a 4-element uint32 array
    for j := 0; j < 4; j++ {
        for i := 3; i >= 0; i-- {
            numbers[j] = (numbers[j] << 8) | uint32(buff[4*j + i])
        }
    }

    return numbers
}
