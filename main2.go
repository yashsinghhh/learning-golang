package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	totalTickets = 50
)

var (
	availableTickets = totalTickets
	mutex            sync.Mutex
)

func main() {
	http.HandleFunc("/", ticketBookingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ticketBookingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		displayTicketAvailability(w)
	case "POST":
		bookTicket(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func displayTicketAvailability(w http.ResponseWriter) {
	fmt.Fprintf(w, "Available Tickets: %d\n", availableTickets)
}

func bookTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	quantityStr := r.Form.Get("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	if quantity <= 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if availableTickets < quantity {
		http.Error(w, "Insufficient tickets available", http.StatusBadRequest)
		return
	}

	availableTickets -= quantity
	fmt.Fprintf(w, "Successfully booked %d ticket(s)! Available Tickets: %d\n", quantity, availableTickets)
}
