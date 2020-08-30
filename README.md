# cache-as-a-service

## Setup 
``` gcloud components install app-engine-go ```

## Deploy 

``` gcloud app deploy```

### Set Cache

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"key":"name","value":"test"}' \
  https://cache-as-a-service-dot-demoneil.df.r.appspot.com
  
```

### Get Cache 

```
curl "https://cache-as-a-service-dot-demoneil.df.r.appspot.com?key=name1"
```