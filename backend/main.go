package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"

	"golang.org/x/net/context"

	pb "github.com/muehlburger/heat/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	livingroom = getenv("LIVINGROOM")
	bedroom    = getenv("BEDROOM")
	bath       = getenv("BATH")
	kidsroom   = getenv("KIDSROOM")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func getDeviceID(deviceID string) string {
	switch deviceID {
	case "kidsroom":
		return kidsroom
	case "bedroom":
		return bedroom
	case "bath":
		return bath
	default:
		return livingroom
	}
}

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterHeatServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve %v", err)
	}
}

type server struct{}

func (server) Set(ctx context.Context, temp *pb.Temp) (*pb.Temp, error) {
	deviceID := getDeviceID(temp.Room)
	logrus.Infof("Set temperature to %d°C for %s (%s)", temp.Value, temp.Room, deviceID)
	cmd := exec.Command("sudo", "/home/pi/heaterControl.exp", deviceID, "00000000", fmt.Sprintf("%d", temp.Value))
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("temp updated failed: %s", data)
	}
	return &pb.Temp{Value: temp.Value, Room: temp.Room}, nil
}
