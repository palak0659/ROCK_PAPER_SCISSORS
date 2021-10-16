 package main
  import "math/rand"
  import  "os"
  import "bufio"
  import "fmt"
  import "strings"
  var playerScore int
  var computerScore int
  
   func main() {
   fmt.Println("ROCK PAPER SCISSOR")
   fmt.Println("ENTER YOUR NAME PLEASE:")
   reader := bufio.NewReader(os.Stdin)
   name, _ := reader.ReadString('\n')
   name = strings.Replace(name,"\n", "",-1)
 
   fmt.Printf("%s %s\n", " Welcome,Do you want to start Playing",name)
   for{
   fmt.Println("Press s to start")
   fmt.Println("Press e to exit")
   fmt.Println("Press v to View score")

 reader := bufio.NewReader(os.Stdin)
 char, _ , error := reader.ReadRune()
 
 if error != nil{
 fmt.Println("Error Occured")
 fmt.Println(error)
 break
 }
 switch char {
 case 's' :
   fmt.Println("Game Starting...")
   startGame()
   printScores(name)
 case 'e' :
   fmt.Println("Exiting...")
   os.Exit(1)
  case 'v':
   printScores(name)
           }
   }
 }

func printScores(name string) {
  fmt.Printf("\n%s %d : %d %s\n\n",name,playerScore, computerScore, "Computer")
 }
func startGame() {
   fmt.Println("Press r for Rock")
   fmt.Println("Press p for Paper ")
   fmt.Println("Press s for Scissor")

   reader := bufio.NewReader(os.Stdin)
   userChoice, _, error := reader.ReadRune()

  if error != nil{
       fmt.Println(error)
 }
   computerChoice := getComputerChoice()
   choices :=map[string]string{"r" : "Rock","p": "Paper","s": "Scissor"}

  fmt.Println("computer chooses ",choices[ string(computerChoice)])
  results(userChoice, computerChoice)
 }
 func getComputerChoice() rune {
   choices := [3]rune{'r', 's', 'p'}
   return choices[rand.Intn(len(choices))]
  }
 
 func results (user rune, computer rune) {
  combinedresults := string(user) + string(computer)
  switch combinedresults {
 case "rs" , "pr" ,"sp":
  fmt.Printf("YOU WON")
  playerScore++
case "rp" , "ps" ,"sr":
  fmt.Printf("YOU Lost")
  computerScore++
 case "rr" , "pp" ,"ss":
  fmt.Printf("MATCH IS DRAW")
         }
  }
