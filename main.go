package main

import (
	"fmt"
	"net/http"
	time2 "time"
)

type PingPongHandler struct {
    Pingcounter int
    buscounter int
    hunderedcounter int
}

func (h *PingPongHandler) PingPongMsg(w http.ResponseWriter, r *http.Request) {
	Mytime:= time2.Now()

	h.Pingcounter = h.Pingcounter+1
	h.buscounter = h.buscounter +1
	h.hunderedcounter = h.hunderedcounter + 100

	fmt.Println("port run on 8080 at %s",Mytime.String())

	//fmt.Fprintf(w, "Sunil, %s , %s %d", r.URL.Path[1:],Mytime.String(),counter)
	if h.buscounter % 5 == 0{
		fmt.Fprintf(w, "Sunil, %s , you have requests %d bas\n", r.URL.Path[1:],h.Pingcounter)
		fmt.Fprintf(w, "Sunil, %s , you have requests %d times\n", r.URL.Path[1:],h.hunderedcounter)

	} else{
		fmt.Fprintf(w, "Sunil, %s , you have requests %d times\n", r.URL.Path[1:],h.Pingcounter)
		fmt.Fprintf(w, "Sunil, %s , you have requests %d times\n", r.URL.Path[1:],h.hunderedcounter)
	}


	fmt.Println(h.Pingcounter)
}

func main() {
	//http.Handle("/pingpong",PingPongMsg)
	a := PingPongHandler{}
	http.Handle("/ping",http.HandlerFunc(a.PingPongMsg))//for registering the route
	http.ListenAndServe(":65354", nil)
}

