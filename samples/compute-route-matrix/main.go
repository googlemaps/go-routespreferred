// [START maps_routes_preferred_compute_routes]
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	routespreferred "developers.google.com/maps/go/routespreferred/v1"
	routespb "google.golang.org/genproto/googleapis/maps/routes/v1"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc/metadata"
	"google.golang.org/api/option"
)

const (
	// https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys
	credentialsFile = "service-account.json"
	// Note that setting the field mask to * is OK for testing, but discouraged in
	// production.
	// For example, for ComputeRouteMatrix, set the field mask to
	// "distanceMeters,duration,status" in order to get the route distances,
	// durations, and status.
	fieldMask = "*"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// instantiate a client
	c, err := routespreferred.NewRoutesPreferredClient(ctx,
		option.WithCredentialsFile(credentialsFile))
	defer c.Close()

	if err != nil {
		log.Fatal(err)
	}

	// create the origin using a latitude and longitude
	origins := []*routespb.RouteMatrixOrigin{{
		Waypoint: &routespb.Waypoint{
			LocationType: &routespb.Waypoint_Location{
				Location: &routespb.Location{
					LatLng: &latlng.LatLng{
						Latitude: 37.417670,
						Longitude:  -122.0827784,
					},
				},
			},
		},
		RouteModifiers: &routespb.RouteModifiers{
			AvoidTolls:    false,
			AvoidHighways: true,
			AvoidFerries:  true,
		},
	}}

	// create the destination using a latitude and longitude
	destinations := []*routespb.RouteMatrixDestination{{
		Waypoint: &routespb.Waypoint{
			LocationType: &routespb.Waypoint_Location{
				Location: &routespb.Location{
					LatLng: &latlng.LatLng{
						Latitude: 37.417670,
						Longitude:  -122.079595,
					},
				},
			},
		},
	}}

	// create the request with additional options
	req := &routespb.ComputeRouteMatrixRequest{
		Origins:                  origins,
		Destinations:             destinations,
		TravelMode:               routespb.RouteTravelMode_DRIVE,
		RoutingPreference:        routespb.RoutingPreference_TRAFFIC_AWARE,
	}

	// set the field mask
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", fieldMask)

	// execute rpc
	stream, err := c.ComputeRouteMatrix(ctx, req)

	if err != nil {
		// "rpc error: code = InvalidArgument desc = Request contains an invalid
		// argument" may indicate that your project lacks access to Routes Preferred
		log.Fatal(err)
	}

	for {
		element, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Duration of route: %d\n", element.Duration.Seconds)
	}
}

// [END maps_routes_preferred_compute_routes]
