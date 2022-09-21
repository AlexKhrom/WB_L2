package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type EventsHand struct {
	DB *sql.DB
}

func NewEventsHand(db *sql.DB) *EventsHand {
	hand := new(EventsHand)
	hand.DB = db
	return hand
}

type Event struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Date   int    `json:"date"`
}

var days = map[string]int64{
	"Monday":    0,
	"Tuesday":   1,
	"Wednesday": 2,
	"Thursday":  3,
	"Friday":    4,
	"Saturday":  5,
	"Sunday":    6,
}

var months = map[string]int64{
	"January":   31,
	"February":  28,
	"March":     31,
	"April":     30,
	"May":       31,
	"June":      30,
	"July":      31,
	"August":    31,
	"September": 30,
	"October":   31,
	"November":  30,
	"December":  31,
}

func (h *EventsHand) CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi createEvent")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't read request body")
		return
	}
	fmt.Println("request  = ", string(body))
	event := &Event{}

	err1 := json.Unmarshal(body, event)
	if err1 != nil {
		fmt.Println("some bad")
		JSONError(w, http.StatusBadRequest, "cant unpack payload")
		return
	}
	fmt.Println("event = ", event)

	var lastInsertId int
	err = h.DB.QueryRow(
		`INSERT INTO events("user_id","title","date") VALUES ($1,$2,$3) RETURNING id`,
		event.UserId,
		event.Title,
		event.Date,
	).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("err new event sql = ", err)
		return
	}

	fmt.Println("event  last insert Id = ", lastInsertId)

}

func (h *EventsHand) SomeHand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi someHand")

}

func (h *EventsHand) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi updateEvent")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't read request body")
		return
	}
	fmt.Println("request  = ", string(body))
	event := &Event{}

	err1 := json.Unmarshal(body, event)
	if err1 != nil {
		fmt.Println("some bad")
		JSONError(w, http.StatusBadRequest, "cant unpack payload")
		return
	}
	fmt.Println("event = ", event)

	_, err = h.DB.Query(`UPDATE "events" set title='` + event.Title + `',date=` + strconv.Itoa(event.Date) + ` where user_id=` + strconv.Itoa(event.UserId) + ` and id=` + strconv.Itoa(event.Id))
	if err != nil {
		fmt.Println("err delete events sql = ", err)
		JSONError(w, http.StatusInternalServerError, "err delete events sql")
		return
	}

}
func (h *EventsHand) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi deleteEvent")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't read request body")
		return
	}
	fmt.Println("request  = ", string(body))
	event := &Event{}

	err1 := json.Unmarshal(body, event)
	if err1 != nil {
		fmt.Println("some bad")
		JSONError(w, http.StatusBadRequest, "cant unpack payload")
		return
	}
	fmt.Println("event = ", event)

	fmt.Println(`DELETE  FROM "events" where user_id=` + strconv.Itoa(event.UserId) + ` and id=` + strconv.Itoa(event.Id))
	_, err = h.DB.Query(`DELETE  FROM "events" where user_id=0 and id=` + strconv.Itoa(event.Id))
	if err != nil {
		fmt.Println("err delete events sql = ", err)
		JSONError(w, http.StatusInternalServerError, "err delete events sql")
		return
	}

}
func (h *EventsHand) GetEventsDay(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi getEventsDay")

	date := r.URL.Query().Get("date")
	if date == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get param date")
		return
	}
	fmt.Println("date = ", date)

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't parse date")
		return
	}

	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get param userId")
		return
	}

	fmt.Println("date time = ", days[dateTime.Weekday().String()])

	dateStart := strconv.FormatInt(dateTime.Unix()*1000, 10)
	dateEnd := strconv.FormatInt(dateTime.Unix()*1000+24*60*60*1000, 10)

	fmt.Println("datestart = ", dateStart)
	fmt.Println("dateEnd = ", dateEnd)

	rows, err := h.DB.Query(`SELECT "id","user_id","title","date" FROM "events" where user_id=` + userId + ` and date>=` + dateStart + " and date<" + dateEnd)
	if err != nil {
		fmt.Println("err get events sql = ", err)
		JSONError(w, http.StatusInternalServerError, "err get events sql")
		return
	}

	var events []Event
	defer rows.Close()
	for rows.Next() {
		var event Event

		err = rows.Scan(
			&event.Id,
			&event.UserId,
			&event.Title,
			&event.Date,
		)
		if err != nil {
			fmt.Println("err get events scan = ", err)
			JSONError(w, http.StatusInternalServerError, "err get events scan")
			return
		}
		events = append(events, event)
	}
	fmt.Println("events = ", events)

	eventJson, err1 := json.Marshal(events)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events json.marshal")
		return
	}
	_, err1 = w.Write(eventJson)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events write response")
		return
	}

}
func (h *EventsHand) GetEventsWeek(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi GetEventsWeek")

	date := r.URL.Query().Get("date")
	if date == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get param date")
		return
	}
	fmt.Println("date = ", date)

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't parse date")
		return
	}

	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get userId date")
		return
	}

	fmt.Println("date time = ", days[dateTime.Weekday().String()])

	dateStart := strconv.FormatInt(dateTime.Unix()*1000-days[dateTime.Weekday().String()]*24*60*60*1000, 10)
	dateEnd := strconv.FormatInt(dateTime.Unix()*1000-days[dateTime.Weekday().String()]*24*60*60*1000+7*24*60*60*1000, 10)

	fmt.Println("datestart = ", dateStart)
	fmt.Println("dateEnd = ", dateEnd)

	rows, err := h.DB.Query(`SELECT "id","user_id","title","date" FROM "events" where user_id=` + userId + ` and date>=` + dateStart + " and date<" + dateEnd)
	if err != nil {
		fmt.Println("err get events sql = ", err)
		JSONError(w, http.StatusInternalServerError, "err get events sql")
		return
	}

	var events []Event
	defer rows.Close()
	for rows.Next() {
		var event Event

		err = rows.Scan(
			&event.Id,
			&event.UserId,
			&event.Title,
			&event.Date,
		)
		if err != nil {
			fmt.Println("err get events scan = ", err)
			JSONError(w, http.StatusInternalServerError, "err get events scan")
			return
		}
		events = append(events, event)
	}
	fmt.Println("events = ", events)
	eventJson, err1 := json.Marshal(events)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events json.marshal")
		return
	}
	_, err1 = w.Write(eventJson)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events write response")
		return
	}
}
func (h *EventsHand) GetEventsMonth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi GetEventsMonth")

	date := r.URL.Query().Get("date")
	if date == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get param date")
		return
	}
	fmt.Println("date = ", date)

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		JSONError(w, http.StatusBadRequest, "can't parse date")
		return
	}

	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		JSONError(w, http.StatusServiceUnavailable, "no get param userId")
		return
	}

	fmt.Println("date time = ", dateTime.Month().String())

	dateStart := strconv.FormatInt(dateTime.Unix()*1000-int64(dateTime.Day()-1)*24*60*60*1000, 10)
	dateEnd := strconv.FormatInt(dateTime.Unix()*1000-int64(dateTime.Day()-1)*24*60*60*1000+months[dateTime.Month().String()]*24*60*60*1000, 10)

	fmt.Println("datestart = ", dateStart)
	fmt.Println("dateEnd = ", dateEnd)

	rows, err := h.DB.Query(`SELECT "id","user_id","title","date" FROM "events" where user_id=` + userId + ` and date>=` + dateStart + " and date<" + dateEnd)
	if err != nil {
		fmt.Println("err get events sql = ", err)
		JSONError(w, http.StatusInternalServerError, "err get events sql")
		return
	}

	var events []Event
	defer rows.Close()
	for rows.Next() {
		var event Event

		err = rows.Scan(
			&event.Id,
			&event.UserId,
			&event.Title,
			&event.Date,
		)
		if err != nil {
			fmt.Println("err get events scan = ", err)
			JSONError(w, http.StatusInternalServerError, "err get events scan")
			return
		}
		events = append(events, event)
	}
	fmt.Println("events = ", events)
	eventJson, err1 := json.Marshal(events)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events json.marshal")
		return
	}
	_, err1 = w.Write(eventJson)
	if err1 != nil {
		JSONError(w, http.StatusInternalServerError, "some bad get events write response")
		return
	}
}

func JSONError(w http.ResponseWriter, status int, msg string) {
	resp, err := json.Marshal(map[string]interface{}{
		"status": status,
		"error":  msg,
	})
	w.WriteHeader(status)
	if err != nil {
		fmt.Println("error in JSONError ")
		return
	}
	_, err2 := w.Write(resp)
	if err2 != nil {
		fmt.Println("some bad in JSONError write response")
	}
}
