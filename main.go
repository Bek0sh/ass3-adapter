package main

import "log"

type Igolang interface {
	ConvertToGolang()
}

type ConvertAdapter struct {
	p IPython
}

func NewAdapter(p IPython) *ConvertAdapter {
	return &ConvertAdapter{p: p}
}

func (c *ConvertAdapter) ConvertToGolang() {
	c.p.Run()
}

type ToGolang struct {
	adapter *ConvertAdapter
}

func NewToGo(p IPython) Igolang {
	return &ToGolang{adapter: NewAdapter(p)}
}

func (t *ToGolang) ConvertToGolang() {
	t.adapter.p.Run()
	log.Println("code is converting into golang")

}

func main() {
	python := Python{}

	gol := NewToGo(&python)
	gol.ConvertToGolang()
}
