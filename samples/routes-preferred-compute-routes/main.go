// [START maps_routes_preferred_compute_routes]
package main

import (
	"context"
	"fmt"
	"log"

	"developers.google.com/maps/go/routes/v1"
	routespb "google.golang.org/genproto/googleapis/maps/routes/v1"
	"google.golang.org/genproto/googleapis/type/latlng"
)

func main() {
	ctx := context.Background()

	// instantiate client with default credentials
	c, err := routes.NewRoutesPreferredClient(ctx)

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

	// execute rpc
	resp, err := c.ComputeRoutes(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Duration of route %d", resp.Routes[0].Duration.Seconds)
}
// [END maps_routes_preferred_compute_routes]
