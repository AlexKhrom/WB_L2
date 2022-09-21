package main

import (
	"L2/task_11/handlers"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"net/http"

	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		fmt.Println("\nhi middleware")
		fmt.Println("////////////////////////")
		fmt.Println(time.Now().String()[:23])
		fmt.Println("url: ", r.URL.Path)
		fmt.Println("method: ", r.Method)
		fmt.Println("////////////////////////\n")

		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	dsn := "postgres://postgres:root@localhost:5432/WB_L2?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("errpr!!!!", err)
		return
	}

	db.SetMaxOpenConns(10)

	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		panic(err)
	}

	eventsHandler := handlers.NewEventsHand(db)

	createEventHandler := http.HandlerFunc(eventsHandler.CreateEvent)
	updateEventHandler := http.HandlerFunc(eventsHandler.UpdateEvent)
	deleteEventHandler := http.HandlerFunc(eventsHandler.DeleteEvent)
	getEventsDayHandler := http.HandlerFunc(eventsHandler.GetEventsDay)
	getEventsWeekHandler := http.HandlerFunc(eventsHandler.GetEventsWeek)
	getEventsMonthHandler := http.HandlerFunc(eventsHandler.GetEventsMonth)
	//someHand := http.HandlerFunc(eventsHandler.SomeHand)

	mux.Handle("/api/createEvent", Middleware(createEventHandler))
	//mux.Handle("/", Middleware(someHand))
	mux.Handle("/api/updateEvent", Middleware(updateEventHandler))
	mux.Handle("/api/deleteEvent", Middleware(deleteEventHandler))
	mux.Handle("/api/events_for_day", Middleware(getEventsDayHandler))
	mux.Handle("/api/events_for_week", Middleware(getEventsWeekHandler))
	mux.Handle("/api/events_for_month", Middleware(getEventsMonthHandler))

	//http.HandleFunc("/updateEvent", updateEvent)
	//http.HandleFunc("/deleteEvent", deleteEvent)
	//http.HandleFunc("/events_for_day", getEvents)
	//http.HandleFunc("/events_for_week", getEvents)
	//http.HandleFunc("/events_for_month", getEvents)

	//_ = handlers.SpaHandler{StaticPath: ".././static/dist", IndexPath: "index.html"}
	fs := http.FileServer(http.Dir("./static/dist"))
	mux.Handle("/", fs)

	port := "8085"
	fmt.Println("start serv on port " + port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("can't Listen and server")
		return
	}
}
