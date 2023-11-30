package main

import "fmt"

type P struct {
}

func (p *P) Name(prefix string) string {
	return fmt.Sprintf("%s, shared-object", prefix)
}

var Plugin P
