package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	in, _ := ioutil.ReadAll(os.Stdin)
	s := bufio.NewScanner(bytes.NewReader(in))
	var b bytes.Buffer
	for s.Scan() {
		l := s.Text()
		fmt.Printf("//\t%s\n", l)
		l = l[1 : len(l)-1]
		single := strings.Index(l, "),") == -1
		if !single {
			fmt.Fprintln(&b, "\t{")
		}
		for {
			i := strings.Index(l, "),")
			if i == -1 {
				if len(l) == 0 {
					break
				}
				i = len(l)
			} else {
				i++
			}
			if !single {
				fmt.Fprint(&b, "\t")
			}
			fmt.Fprintf(&b, "\tswp(a,%s\n", l[1:i])
			if i == len(l) {
				i--
			}
			l = l[i+1:]
		}
		if !single {
			fmt.Fprintln(&b, "\t}")
		}
	}
	fmt.Printf("func SN%02d(a []int) {\n", n)
	fmt.Printf("\t_ = a[%d]\n", n-1)
	_, _ = b.WriteTo(os.Stdout)
	fmt.Println("}")
}
