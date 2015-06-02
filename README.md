iFooz-engine
====

This is the engine for a Foosball IoT rig called iFooz.

Run the engine
----
- Ensure Go and GOPATH is setup and configured
- clone
- navigate to repo
- `go get`
- Start redis-server at localhost:6379 or set REDIS_URL and REDIS_PORT
- `go run web.go`
- Service runs at localhost:8066 (or :$PORT)
