package main

import "fmt"

type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contactInfo
}

func main() {
    p := person{
        firstName: "Jim",
        lastName:  "Henderson",
        contactInfo: contactInfo{
            email:   "test@test.test",
            zipCode: 94000,
        },
    }

    p.updateName("Alex")
    p.print()
}

func (p *person) updateName(n string) {
    (*p).firstName = n
}

func (p person) print() {
    fmt.Printf("%+v", p)
}
