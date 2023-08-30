# Golang bindings for Strava V3 API
This is the standard golang bindings for Strava's V3 API (current API as of 2023).  This is made available by Strava via Swagger with no modifications here, but it was sufficiently difficult to get working that I thoughts I'd post it on github for others, along with an explanation for why I did it this way (see the Makefile for a simple description of "what" I did).

For general use feel free to require this repos in golang.  I don't intend to update it in the future, but I also make no promises.  If it breaks then your recourse will be to clone and revert this repository and update your requires to your own repository.

Strava uses the same string values in multiple enums.  Golang, which has no built in support for enums, uses the value of these strings to generate static variable names.  The variables from these different enums then collide and produce compiler errors.  You can, of course, simply rename these variables to address the conflict.  I instead used openapi-generator (compatible with Swagger 2 which is what Strava requires) with its -p enumClassPrefix=true option to do this for me.  Openapi-generator complains that the provided swagger input is invalid, so I silence this with --skip-validate-spec.

To use this client you will first need to perform oauth2 authorization, create an http client with the oauth2 token, and assign this client into the strava HTTPClient struct.  I use github.com/stevenpelley/oauth2-accesstoken-golang to create a token for development use.

See example/athlete for an example.  This assumes you have the input oauth_client_config.json used in github.com/stevenpelley/oauth2-accesstoken-golang as well as the token.json that results.  These files are both listed in .gitignore to prevent commiting secrets.

What follows is the README generated from the strava swagger itself.

# Go API client for strava3golang

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import strava3golang "github.com/stevenpelley/strava3golang"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), strava3golang.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), strava3golang.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), strava3golang.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), strava3golang.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://www.strava.com/api/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivitiesAPI* | [**CreateActivity**](docs/ActivitiesAPI.md#createactivity) | **Post** /activities | Create an Activity
*ActivitiesAPI* | [**GetActivityById**](docs/ActivitiesAPI.md#getactivitybyid) | **Get** /activities/{id} | Get Activity
*ActivitiesAPI* | [**GetCommentsByActivityId**](docs/ActivitiesAPI.md#getcommentsbyactivityid) | **Get** /activities/{id}/comments | List Activity Comments
*ActivitiesAPI* | [**GetKudoersByActivityId**](docs/ActivitiesAPI.md#getkudoersbyactivityid) | **Get** /activities/{id}/kudos | List Activity Kudoers
*ActivitiesAPI* | [**GetLapsByActivityId**](docs/ActivitiesAPI.md#getlapsbyactivityid) | **Get** /activities/{id}/laps | List Activity Laps
*ActivitiesAPI* | [**GetLoggedInAthleteActivities**](docs/ActivitiesAPI.md#getloggedinathleteactivities) | **Get** /athlete/activities | List Athlete Activities
*ActivitiesAPI* | [**GetZonesByActivityId**](docs/ActivitiesAPI.md#getzonesbyactivityid) | **Get** /activities/{id}/zones | Get Activity Zones
*ActivitiesAPI* | [**UpdateActivityById**](docs/ActivitiesAPI.md#updateactivitybyid) | **Put** /activities/{id} | Update Activity
*AthletesAPI* | [**GetLoggedInAthlete**](docs/AthletesAPI.md#getloggedinathlete) | **Get** /athlete | Get Authenticated Athlete
*AthletesAPI* | [**GetLoggedInAthleteZones**](docs/AthletesAPI.md#getloggedinathletezones) | **Get** /athlete/zones | Get Zones
*AthletesAPI* | [**GetStats**](docs/AthletesAPI.md#getstats) | **Get** /athletes/{id}/stats | Get Athlete Stats
*AthletesAPI* | [**UpdateLoggedInAthlete**](docs/AthletesAPI.md#updateloggedinathlete) | **Put** /athlete | Update Athlete
*ClubsAPI* | [**GetClubActivitiesById**](docs/ClubsAPI.md#getclubactivitiesbyid) | **Get** /clubs/{id}/activities | List Club Activities
*ClubsAPI* | [**GetClubAdminsById**](docs/ClubsAPI.md#getclubadminsbyid) | **Get** /clubs/{id}/admins | List Club Administrators
*ClubsAPI* | [**GetClubById**](docs/ClubsAPI.md#getclubbyid) | **Get** /clubs/{id} | Get Club
*ClubsAPI* | [**GetClubMembersById**](docs/ClubsAPI.md#getclubmembersbyid) | **Get** /clubs/{id}/members | List Club Members
*ClubsAPI* | [**GetLoggedInAthleteClubs**](docs/ClubsAPI.md#getloggedinathleteclubs) | **Get** /athlete/clubs | List Athlete Clubs
*GearsAPI* | [**GetGearById**](docs/GearsAPI.md#getgearbyid) | **Get** /gear/{id} | Get Equipment
*RoutesAPI* | [**GetRouteAsGPX**](docs/RoutesAPI.md#getrouteasgpx) | **Get** /routes/{id}/export_gpx | Export Route GPX
*RoutesAPI* | [**GetRouteAsTCX**](docs/RoutesAPI.md#getrouteastcx) | **Get** /routes/{id}/export_tcx | Export Route TCX
*RoutesAPI* | [**GetRouteById**](docs/RoutesAPI.md#getroutebyid) | **Get** /routes/{id} | Get Route
*RoutesAPI* | [**GetRoutesByAthleteId**](docs/RoutesAPI.md#getroutesbyathleteid) | **Get** /athletes/{id}/routes | List Athlete Routes
*SegmentEffortsAPI* | [**GetEffortsBySegmentId**](docs/SegmentEffortsAPI.md#geteffortsbysegmentid) | **Get** /segment_efforts | List Segment Efforts
*SegmentEffortsAPI* | [**GetSegmentEffortById**](docs/SegmentEffortsAPI.md#getsegmenteffortbyid) | **Get** /segment_efforts/{id} | Get Segment Effort
*SegmentsAPI* | [**ExploreSegments**](docs/SegmentsAPI.md#exploresegments) | **Get** /segments/explore | Explore segments
*SegmentsAPI* | [**GetLoggedInAthleteStarredSegments**](docs/SegmentsAPI.md#getloggedinathletestarredsegments) | **Get** /segments/starred | List Starred Segments
*SegmentsAPI* | [**GetSegmentById**](docs/SegmentsAPI.md#getsegmentbyid) | **Get** /segments/{id} | Get Segment
*SegmentsAPI* | [**StarSegment**](docs/SegmentsAPI.md#starsegment) | **Put** /segments/{id}/starred | Star Segment
*StreamsAPI* | [**GetActivityStreams**](docs/StreamsAPI.md#getactivitystreams) | **Get** /activities/{id}/streams | Get Activity Streams
*StreamsAPI* | [**GetRouteStreams**](docs/StreamsAPI.md#getroutestreams) | **Get** /routes/{id}/streams | Get Route Streams
*StreamsAPI* | [**GetSegmentEffortStreams**](docs/StreamsAPI.md#getsegmenteffortstreams) | **Get** /segment_efforts/{id}/streams | Get Segment Effort Streams
*StreamsAPI* | [**GetSegmentStreams**](docs/StreamsAPI.md#getsegmentstreams) | **Get** /segments/{id}/streams | Get Segment Streams
*UploadsAPI* | [**CreateUpload**](docs/UploadsAPI.md#createupload) | **Post** /uploads | Upload Activity
*UploadsAPI* | [**GetUploadById**](docs/UploadsAPI.md#getuploadbyid) | **Get** /uploads/{uploadId} | Get Upload


## Documentation For Models

 - [ActivityStats](docs/ActivityStats.md)
 - [ActivityTotal](docs/ActivityTotal.md)
 - [ActivityType](docs/ActivityType.md)
 - [ActivityZone](docs/ActivityZone.md)
 - [AltitudeStream](docs/AltitudeStream.md)
 - [BaseStream](docs/BaseStream.md)
 - [CadenceStream](docs/CadenceStream.md)
 - [ClubActivity](docs/ClubActivity.md)
 - [ClubAthlete](docs/ClubAthlete.md)
 - [Comment](docs/Comment.md)
 - [DetailedActivity](docs/DetailedActivity.md)
 - [DetailedAthlete](docs/DetailedAthlete.md)
 - [DetailedClub](docs/DetailedClub.md)
 - [DetailedGear](docs/DetailedGear.md)
 - [DetailedSegment](docs/DetailedSegment.md)
 - [DetailedSegmentEffort](docs/DetailedSegmentEffort.md)
 - [DistanceStream](docs/DistanceStream.md)
 - [Error](docs/Error.md)
 - [ExplorerResponse](docs/ExplorerResponse.md)
 - [ExplorerSegment](docs/ExplorerSegment.md)
 - [Fault](docs/Fault.md)
 - [HeartRateZoneRanges](docs/HeartRateZoneRanges.md)
 - [HeartrateStream](docs/HeartrateStream.md)
 - [Lap](docs/Lap.md)
 - [LatLngStream](docs/LatLngStream.md)
 - [MetaActivity](docs/MetaActivity.md)
 - [MetaAthlete](docs/MetaAthlete.md)
 - [MetaClub](docs/MetaClub.md)
 - [MovingStream](docs/MovingStream.md)
 - [PhotosSummary](docs/PhotosSummary.md)
 - [PhotosSummaryPrimary](docs/PhotosSummaryPrimary.md)
 - [PolylineMap](docs/PolylineMap.md)
 - [PowerStream](docs/PowerStream.md)
 - [PowerZoneRanges](docs/PowerZoneRanges.md)
 - [Route](docs/Route.md)
 - [SmoothGradeStream](docs/SmoothGradeStream.md)
 - [SmoothVelocityStream](docs/SmoothVelocityStream.md)
 - [Split](docs/Split.md)
 - [SportType](docs/SportType.md)
 - [StreamSet](docs/StreamSet.md)
 - [SummaryActivity](docs/SummaryActivity.md)
 - [SummaryAthlete](docs/SummaryAthlete.md)
 - [SummaryClub](docs/SummaryClub.md)
 - [SummaryGear](docs/SummaryGear.md)
 - [SummaryPRSegmentEffort](docs/SummaryPRSegmentEffort.md)
 - [SummarySegment](docs/SummarySegment.md)
 - [SummarySegmentEffort](docs/SummarySegmentEffort.md)
 - [TemperatureStream](docs/TemperatureStream.md)
 - [TimeStream](docs/TimeStream.md)
 - [TimedZoneRange](docs/TimedZoneRange.md)
 - [UpdatableActivity](docs/UpdatableActivity.md)
 - [Upload](docs/Upload.md)
 - [ZoneRange](docs/ZoneRange.md)
 - [Zones](docs/Zones.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### strava_oauth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://www.strava.com/api/v3/oauth/authorize
- **Scopes**: 
 - **read**: Read public segments, public routes, public profile data, public posts, public events, club feeds, and leaderboards
 - **read_all**: Read private routes, private segments, and private events for the user
 - **profile:read_all**: Read all profile information even if the user has set their profile visibility to Followers or Only You
 - **profile:write**: Update the user's weight and Functional Threshold Power (FTP), and access to star or unstar segments on their behalf
 - **activity:read**: Read the user's activity data for activities that are visible to Everyone and Followers, excluding privacy zone data
 - **activity:read_all**: The same access as activity:read, plus privacy zone data and access to read the user's activities with visibility set to Only You
 - **activity:write**: Access to create manual activities and uploads, and access to edit any activities that are visible to the app, based on activity read access level

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



