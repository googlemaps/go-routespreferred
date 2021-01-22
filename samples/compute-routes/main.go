// [START maps_routes_preferred_compute_routes]
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	routespreferred "developers.google.com/maps/go/routespreferred/v1"
	"google.golang.org/api/option"
	routespb "google.golang.org/genproto/googleapis/maps/routes/v1"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc/metadata"
)

const (
	// https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys
	credentialsFile = "service-account.json"
	// Note that setting the field mask to * is OK for testing, but discouraged in
	// production.
	// For example, for ComputeRoutes, set the field mask to
	// "routes.distanceMeters,routes.duration,routes.polyline.encodedPolyline"
	// in order to get the route distances, durations, and encoded polylines.
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
	origin := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Latitude:  37.417670,
					Longitude: -122.0827784,
				},
			},
		},
	}

	// create the destination using a latitude and longitude
	destination := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Latitude:  37.417670,
					Longitude: -122.079595,
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

	// set the field mask
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", fieldMask)

	// execute rpc
	resp, err := c.ComputeRoutes(ctx, req)

	if err != nil {
		// "rpc error: code = InvalidArgument desc = Request contains an invalid
		// argument" may indicate that your project lacks access to Routes Preferred
		log.Fatal(err)
	}

	fmt.Printf("Duration of route %d", resp.Routes[0].Duration.Seconds)
}

// [END maps_routes_preferred_compute_routes]
