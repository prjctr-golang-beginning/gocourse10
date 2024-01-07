package database

import (
	"fmt"
	"math/rand"
)

// Observer інтерфейс для спостерігачів, які реагують на події.
type Observer interface {
	OnNotify(event string)
}

// Observable інтерфейс для об'єктів, що можуть нотифікувати спостерігачів.
type Observable interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers(event string)
}

type DefaultObserver struct {
	observers []Observer
}

// Реалізація методів Observable.
func (d *DefaultObserver) RegisterObserver(o Observer) {
	fmt.Println(`Observer registered`)
	d.observers = append(d.observers, o)
}

func (d *DefaultObserver) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range d.observers {
		if observer == o {
			indexToRemove = i
			fmt.Println(`Observer removed`)
			break
		}
	}
	d.observers = append(d.observers[:indexToRemove], d.observers[indexToRemove+1:]...)
}

func (d *DefaultObserver) NotifyObservers(event string) {
	for _, observer := range d.observers {
		observer.OnNotify(event)
	}
}

// DatabaseActivityLogger спостерігач, який логує події бази даних.
type DatabaseActivityLogger struct{}

func (l *DatabaseActivityLogger) OnNotify(event string) {
	fmt.Println("Database Activity:", event)
}

// DatabaseActivityMetric спостерігач, який вимірює події бази даних.
type DatabaseActivityMetric struct{}

func (l *DatabaseActivityMetric) OnNotify(event string) {
	fmt.Printf("The process '%s' took %d milliseconds\n", event, rand.Int31n(100))
}
