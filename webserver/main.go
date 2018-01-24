package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/muehlburger/heat/api"
	"google.golang.org/grpc"
)

const (
	livingroom = "livingroom"
	bedroom    = "bedroom"
	bathroom   = "bathroom"
	kidsroom   = "kidsroom"
)

var backend = flag.String("b", "hassbian.local:8080", "address of the heat backend")

func main() {
	log.Fatal(http.ListenAndServe(":80", handler()))
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/temp", tempHandler)
	return r
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	room := r.FormValue("r")
	if room == "" || room != livingroom && room != bedroom && room != kidsroom && room != bathroom {
		http.Error(w, "missing room", http.StatusBadRequest)
		return
	}

	value := r.FormValue("v")
	if value == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, "not a number: "+value, http.StatusBadRequest)
		return
	}
	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "could not connect to backend", http.StatusBadRequest)
	}
	defer conn.Close()

	client := pb.NewHeatClient(conn)

	temp := &pb.Temp{
		Value: int32(b),
		Room:  room,
	}
	res, err := client.Set(context.Background(), temp)
	if err != nil {
		log.Fatalf("could not set temp value %d: %v", temp.Value, err)
	}
	fmt.Fprintf(w, "%v", res.Value)
}
