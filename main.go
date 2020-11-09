package main

import (
	"container/list"
	"fmt"
	"time"
)

//Proceso comentario para que no de advertencia
type Proceso struct {
	ID  int
	Inc int
	C   bool
}

//Start comentario para que no de advertencia
func (p *Proceso) Start(id int, i int, c chan bool) {

	p.ID = id
	p.Inc = i
	for {
		if p.C {
			break
		} else {
			p.Inc = p.Inc + 1
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//Mostrar comentario para que no de advertencia
func (p *Proceso) Mostrar() {
	fmt.Println(p.ID, " : ", p.Inc)
}

//Stop comentario para que no de advertencia
func (p *Proceso) Stop(c chan bool) {
	p.C = true
}

func main() {
	lista := list.New()
	opc := 0
	aux := 0
	elim := 0
	c := make(chan bool)
	for opc != 4 {
		fmt.Println("1.- Agregar proceso")
		fmt.Println("2.- Mostrar proceso")
		fmt.Println("3.- Terminar proceso")
		fmt.Println("4.- Salir")
		fmt.Scan(&opc)
		switch opc {
		case 1:
			pr := &Proceso{ID: aux, Inc: 0}
			go pr.Start(aux, 0, c)
			lista.PushBack(pr)
			aux++

		case 2:
			for e := lista.Front(); e != nil; e = e.Next() {
				pro := e.Value.(*Proceso)
				pro.Mostrar()
			}

		case 3:
			fmt.Println("ID del proceso a eliminar: ")
			fmt.Scan(&elim)
			for e := lista.Front(); e != nil; e = e.Next() {
				pro := e.Value.(*Proceso)
				if elim == pro.ID {
					fmt.Println(pro.ID)
					pro.Stop(c)
					lista.Remove(e)
				}
			}
		}

	}

}
