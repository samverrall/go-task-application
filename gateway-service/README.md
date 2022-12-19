# gateway-service

A reverse proxy (REST -> gRPC) Golang microservice.

This service's sole purpose is to act as a gateway with a reverse proxy mapping a REST API to gRPC functions.

## TODO

[] Read a session cookie from the incoming request

[] Read session from the users-service via a sessionID (from cookie)

[] Build a JWT with the userID + authorisation data 

This JWT will be encoded per request, and will be shortlived (only alive for the request). The tokens claims should be serialised with user data + the `sub` should contain the UserID.

[] Pass JWT via a `ctx` through to gRPC calls in services down the line (eg task-service).


