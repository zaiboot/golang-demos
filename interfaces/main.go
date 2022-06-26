package main
import "fmt"

type greeter interface {
    getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}
type russianBot struct {}

func main(){
    eb := englishBot{}
    sp := spanishBot{}
    rb := russianBot{}

    printGreeting(eb)
    printGreeting(sp)
    printGreeting(rb)
}

func printGreeting(b greeter){
   fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string{
    return "Hello"
}

func (spanishBot) getGreeting() string {
    return "Hola"
}

func (russianBot) getGreeting(dialect ...string) string{
    return "Hello"
}

