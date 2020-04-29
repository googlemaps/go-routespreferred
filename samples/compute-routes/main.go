// [START maps_routes_preferred_compute_routes]
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	routespreferred "developers.google.com/maps/go/routespreferred/v1"
	routespb "google.golang.org/genproto/googleapis/maps/routes/v1"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// instantiate a client
	c, err := routespreferred.NewRoutesPreferredClient(ctx)
	defer c.Close()

	if err != nil {
		log.Fatal(err)
	}

	// create the origin using a latitude and longitude
	origin := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Longitude: 37.417670,
					Latitude:  -122.0827784,
				},
			},
		},
	}

	// create the destination using a latitude and longitude
	destination := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Longitude: 37.417670,
					Latitude:  -122.079595,
				},
			},
		},
	}

	// create the request with additional options
	req := &routespb.ComputeRoutesRequest{
		Origin:                   origin,
		Destination:              destination,
		TravelMode:               routespb.RouteTravelMode_DRIVE,
		RoutingPreference:        routespb.RoutingPreference_TRAFFIC_AWARE,
		ComputeAlternativeRoutes: true,
		Units:                    routespb.Units_METRIC,
		LanguageCode:             "en-us",
		RouteModifiers: &routespb.RouteModifiers{
			AvoidTolls:    false,
			AvoidHighways: true,
			AvoidFerries:  true,
		},
		PolylineQuality: routespb.PolylineQuality_OVERVIEW,
	}

	// add an API key to the request
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Api-Key", "YOUR KEY")

	// set the field mask
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", "*")

	// execute rpc
	resp, err := c.ComputeRoutes(ctx, req)

	if err != nil {
		// "rpc error: code = InvalidArgument desc = Request contains an invalid
		// argument" may indicate an invalid API key or project not having
		// access to Routes Preferred
		log.Fatal(err)
	}

	fmt.Printf("Duration of route %d", resp.Routes[0].Duration.Seconds)
}

// [END maps_routes_preferred_compute_routes]
