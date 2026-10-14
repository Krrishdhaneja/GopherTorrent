// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//
// )
//
//	func main(){
//		// scanner := bufio.NewScanner(os.Stdin)
//		// // // hi:= scanner.Scan()
//		// // hi := scanner.Text()
//		// // fmt.Println(hi+"i")
//		// for scanner.Scan(){
//		// 	l := scanner.Text()
//		// 	println(l)
//		// }
//		// if e := scanner.Err(); e != nil{
//		// 	println(os.Stderr, "err: ", e)
//		// }
//		r := bufio.NewReader(os.Stdin)
//		// fmt.Print("Keep writing...")
//		// r.ReadString('\n')
//		// fmt.Println(n)
//		// fmt.Println(r)
//		for {
//	    line, err := r.ReadString('\n')
//
//	    if len(line) > 0 {
//	        fmt.Println(line)
//	    }
//
//	    if err != nil {
//	        break
//	    }
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	r := bufio.NewReader(os.Stdin)
	for{
		line, err := r.ReadString('\n')
		if len(line) > 0{
			fmt.Printf("%q",line)
		}
		if err != nil{
			fmt.Println("Err: ", err)
			return
		}
		
	}
}
