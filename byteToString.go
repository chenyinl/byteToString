package main

import (
"fmt" 
"strings"
//"reflect"
//"unsafe"
"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
/*
func BytesToString(bytes []byte) (s string) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = slice.Data
	str.Len = slice.Len
	return s
}*/
func (ip IPAddr) String() string{
    
	var returnstring []string
	//returnstring = BytesToString(ip)
	//fmt.Println(returnstring)
	
	//s := string(ip)
	//fmt.Println(s)
	//return "abc"
    for _, v:= range ip {
	
	  returnstring = append( returnstring, strconv.Itoa(int(v)))
	  //fmt.Println(reflect.TypeOf(v));
	}
	//fmt.Println(returnstring);
	r :=strings.Join(returnstring,".")
	//fmt.Println(reflect.TypeOf(v))
	return r
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
