func foo(array [4]byte) {
     array[1]++
     fmt.Println(array) // now array is [1, 3, 3, 4]
}

func main() {
    array := [4]byte{1, 2, 3, 4}
    foo(array)
    fmt.Println(array) // still [1, 2, 3, 4]
}
