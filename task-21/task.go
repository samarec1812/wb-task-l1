package main

import "fmt"

/*
	Реализовать паттерн «адаптер» на любом примере.
 */

// паттерн адаптер позволяет двум несовместным интерфейсам взаимодействовать друг с
// с другом
// допустим у нас есть два интерфейса и они различны
// нам необходимо чтобы они смогли обмениваться данными, для этого создаётся промежуточный адаптер
// который будет отвечать за переход с одного интерфейса на другой

type animal interface {
	say()
	run()
}

type wolf struct{}
func (w *wolf)say() {
	fmt.Println("Woooo")
}

func (w *wolf)run() {
	fmt.Println("I'm wolf. I can run")
}

type bird interface {
	say()
	fly()
}

type duck struct {}

func (d *duck) say() {
	fmt.Println("Quack")
}

func (d *duck) fly() {
	fmt.Println("I'm duck. I can fly")
}

type duckAdapter struct {
	bird bird
}

func newDuckAdapter(b bird) *duckAdapter {
	return &duckAdapter{
		bird: b,
	}
}

func (d *duckAdapter) run() {
	d.bird.fly()
}

func (d *duckAdapter) say() {
	d.bird.say()
}

func testWolf(a animal) {
	a.say()
	a.run()
}

func main() {
	wolf := &wolf{}
	duck := &duck{}

	duckAdapterEx := newDuckAdapter(duck)
	testWolf(wolf)

	testWolf(duckAdapterEx)

}