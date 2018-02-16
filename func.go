package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
#include <stdio.h>
int add(int a,int b){
	return a+b;
}
*/
import (
	"C"
)

type Person struct {
	Name string
}

func main() {
	p := &Person{Name: "World"}
	json.NewDecoder(os.Stdin).Decode(p)
	mapD := map[string]string{"message": fmt.Sprintf("Hello %s", p.Name)}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	fmt.Println("demo log")
	log.Println("applog")
	fmt.Println(C.add(1, 3), "call from Cgo")
}
