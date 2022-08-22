# H3 reverse geolocation for Finland

This project contains a Google Coloud Function that works as a reverse geolocation API for h3 indices for some places in Finland. It should cover most areas where people live.

Uber H3 is a 1D geospatial index, it splits the surface of earth to hierarchical hexagons, and addresses the hexagons by 64bit integers. If the integers are similar, the hexagons are most likely close to each other.

You can pass it an h3 id and it will return best guess for city and neighborhood.

It's loading the locations from sqlite file deployed together with the Cloud Function to Google Cloud.

e.g.

```
tomk@xps ~ » curl -X POST --data '{"h3id":"890888534d3ffff"}' https://us-central1-osloveni.cloudfunctions.net/H3RevGeoLocFi 
{"city":"Pori","neighborhood":"Porin seutukunta"}

```

.. or locally, after `cd cmd; go run main.go`

```
tomk@xps ~/kscripts ±master⚡ » curl -i -X POST --data '{"h3id":"890888534d3ffff"}' 127.0.0.1:8084
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Mon, 22 Aug 2022 08:47:34 GMT
Content-Length: 50

{"city":"Pori","neighborhood":"Porin seutukunta"}
```

