# H3 reverse geolocation for Finland

This project contains a Google Coloud Function that works as a reverse geolocation API for h3 indices for some places in Finland.

You can pass it an h3 id and it will return best guess for city and neighborhood.

e.g.

```
curl -i -X POST --data '{"h3id":"890888534d3ffff"}' https://us-central1-osloveni.cloudfunctions.net/H3RevGeoLocFi
```

