// INPUT USING SCANNER
// package main
// import(
// 	"bufio"
// 	"fmt"
// 	"os"
// )
// func main(){
// 	sc := bufio.NewScanner(os.Stdin)
// 	for sc.Scan(){
// 		line := sc.Text()
// 		fmt.Println(line)
// 	}
// 	if e:= sc.Err(); e != nil{
// 		fmt.Println("Err: ", e)
// 	}
//
// }
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
// }

// CONTINOUS INPUT USING NEWREADER()

// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )
// func main(){
// 	r := bufio.NewReader(os.Stdin)
// 	for{
// 		line, err := r.ReadString('\n')
// 		if len(line) > 0{
// 			fmt.Printf("%q",line)
// 		}
// 		if err != nil{
// 			fmt.Println("Err: ", err)
// 			return
// 		}
//
// 	}
// }

// WRITING TO A FILE
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Print("Write file name: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fname := sc.Text()
	f, e := os.Create(fname)
	if e != nil{
		fmt.Println("Err: ",e)
		return
	}
	w := bufio.NewWriter(f)
	fmt.Println()
	fmt.Print("Write the things you want in da file: ")
	sc.Scan()
	content := sc.Text()
	_, e = w.WriteString(content)
	if e != nil{
		fmt.Println("Err2: ",e)
		return
	}
	w.Flush()
}
