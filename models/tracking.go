package models

import (
	"golang.org/x/net/context"
	"log"
	"sync"
	"time"
)

type Server struct {
	trips map[int32][]*Track
	streams map[int32]*Tracking_SyncTripServer
	mutex sync.Mutex
}

func NewServer() *Server{
	s := &Server{streams: make(map[int32]*Tracking_SyncTripServer)}
	return s
}

func (s *Server) SyncTrip(message *JoinTrip, stream Tracking_SyncTripServer) error  {
	log.Printf("client %s connected", message.Vehicle.Name)
	s.streams[message.Vehicle.Id] = &stream
	for {
		time.Sleep(time.Second * 10)
	}
	return nil
}

func (s *Server) SendTrack(ctx context.Context, track *Track) (*Track, error) {
	for id, stream := range s.streams {
		if id == track.Vehicle.Id {
			continue
		}
		err := (*stream).Send(track)
		if err != nil {
			log.Println("error when send track to client")
		}
	}
	return track, nil
}
