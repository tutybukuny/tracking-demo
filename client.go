package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"time"
	"tracking-demo/models"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := models.NewTrackingClient(conn)
	rand.Seed(time.Now().UnixNano())
	vehicle := &models.Vehicle{Id: rand.Int31(), Name: fmt.Sprintf("xe %d", rand.Intn(100)), Vehicle: "o to"}
	st, err := c.SyncTrip(context.Background(), &models.JoinTrip{
		Trip:    &models.Trip{Id: 1},
		Vehicle: vehicle,
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	go func() {
		for {
			_, err := c.SendTrack(context.Background(), &models.Track{Vehicle: vehicle, Point: &models.Point{Latitude: rand.Int31n(100), Longitude: rand.Int31n(100)}})
			if err != nil {
				log.Println("something went wrong")
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()
	for {
		message, err := st.Recv()
		if err == io.EOF {
			return
		}
		log.Printf("Response from server: track %s at point %s", message.Vehicle.Name, message.Point.String())
	}

}
