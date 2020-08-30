# cache-as-a-service

## Setup 
``` gcloud components install app-engine-go ```

## Deploy 

``` gcloud app deploy```

## APIs

### Set Cache (201 Created)

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"key":"name","value":"test"}' \
  https://cache-as-a-service-dot-demoneil.df.r.appspot.com
  
```

### Get Cache (200 OK)

```
curl "https://cache-as-a-service-dot-demoneil.df.r.appspot.com?key=name1"
```

## Errors

* Cache Miss - `404 Not Found`

## TODO
- [x] Error code for cache miss
- [ ] Stict JSON validation for set cache
- [ ] Multi Set
- [ ] Multi Get 
- [ ] Explicit TTL 

