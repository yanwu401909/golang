package main

import (
	"fmt"
	"log"
)

func init() {
	log.Println("Main init ...")
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// func (u user) notify() {
// 	fmt.Printf("Sending User Email To %s<%s>\n",
// 		u.name,
// 		u.email)
// }

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

// func sendPrinter(p printer) {
// 	p.print()
// }

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(u)
	// bill.notify()
	// lisa := &user{"Lisa", "lisa@email.com"}
	// lisa.notify()
	// bill.changeEmail("bill@newdomain.com")
	// bill.notify()
	// lisa.changeEmail("lisa@newdomain.com")
	// lisa.notify()

	// sendPrinter(u1)
	// sendPrinter(u2)
}
