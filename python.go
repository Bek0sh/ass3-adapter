package main

import "log"

type IPython interface {
	Run()
}

type Python struct{}

func (p *Python) Run() {
	log.Println("running python")
}
