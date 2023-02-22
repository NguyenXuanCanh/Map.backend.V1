package compute_routes

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/NguyenXuanCanh/go-starter/config"
	routespb "google.golang.org/genproto/googleapis/maps/routing/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	fieldMask  = "*"
	apiKey     = config.API_KEY
	serverAddr = "routes.googleapis.com:443"
)

func GetComputeRoutes() any {
	config := tls.Config{}
	conn, err := grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(credentials.NewTLS(&config)))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := routespb.NewRoutesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Api-Key", apiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", fieldMask)
	defer cancel()

	// create the origin using a place id
	origin := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_PlaceId{
			PlaceId: "ChIJD-2wOQAvdTERQMcHMUK0xPQ",
		},
	}

	// create the destination using a place id
	destination := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_PlaceId{
			PlaceId: "ChIJg0HGgRwvdTERPHWPenqdENM",
		},
	}
	req := &routespb.ComputeRoutesRequest{
		Origin:                   origin,
		Destination:              destination,
		TravelMode:               routespb.RouteTravelMode_DRIVE,
		RoutingPreference:        routespb.RoutingPreference_TRAFFIC_AWARE,
		ComputeAlternativeRoutes: true,
		Units:                    routespb.Units_METRIC,
		RouteModifiers: &routespb.RouteModifiers{
			AvoidTolls:    false,
			AvoidHighways: true,
			AvoidFerries:  true,
		},
		PolylineQuality: routespb.PolylineQuality_OVERVIEW,
	}

	// execute rpc
	resp, err := client.ComputeRoutes(ctx, req)

	if err != nil {
		// "rpc error: code = InvalidArgument desc = Request contains an invalid
		// argument" may indicate that your project lacks access to Routes
		log.Fatal(err)
	}

	return resp
}
