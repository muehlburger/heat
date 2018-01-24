package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	pb "github.com/muehlburger/heat/api"
	"google.golang.org/grpc"
)

const (
	livingroom = "livingroom"
	bedroom    = "bedroom"
	bath       = "bath"
	kidsroom   = "kidsroom"
)

func checkRoom(room string) error {
	if room == livingroom || room == bedroom || room == bath || room == kidsroom {
		return nil
	}
	return errors.New("could not find room: %s")
}

func main() {
	backend := flag.String("b", "hassbian.local:8080", "address of the heat backend")
	temperature := flag.Int("t", 22, "set temperature to this value")
	room := flag.String("r", "", "room (livingroom, bedroom, bath, kidsroom) to be updated")
	flag.Parse()

	if err := checkRoom(*room); err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewHeatClient(conn)

	temp := &pb.Temp{
		Value: int32(*temperature),
		Room:  *room,
	}
	res, err := client.Set(context.Background(), temp)
	if err != nil {
		log.Fatalf("could not set temp value %d: %v", temp.Value, err)
	}
	log.Printf("Temperature set to %d°C in %s.", res.Value, res.Room)
}
