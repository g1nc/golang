package main

func getPointer(n *int)  {
    *n = *n * *n
}

func returnPointer(n int) *int {
    v := n * n
    return &v
}

func main() {

}
