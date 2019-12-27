[![Build Status](https://travis-ci.org/michaljirman/newsapp.svg?branch=master)](https://travis-ci.org/michaljirman/newsapp)

# Overview
This repository contains a demo News App (API) project consiting of two microservices: 

* apigateway-service
* newsfeeder-service

## `Design`


### Some external libraries 

* [go-kit](https://gokit.io/) 
    
    Projects is built on top of the [go-kit](https://gokit.io/) library which provides collection of Go (golang) packages (libraries) that help you build robust, reliable, maintainable microservices.  

* [http-cache](github.com/victorspringer/http-cache) 

    Provides in-memory caching at the transport layer.

* [gofeed](github.com/mmcdole/gofeed)
    
    Provides a feed parser that supports parsing both RSS and Atom feeds.

* [errors](github.com/pkg/errors)

    Provides a simple error handling primitive with support for call stack in error.

* [logrus](github.com/sirupsen/logrus)

    Provides a structured logger for Go (golang), completely API compatible with the standard library logger.

* [Additional dependencies]()

    ```
        github.com/DATA-DOG/go-sqlmock
        github.com/PuerkitoBio/goquery
        github.com/afex/hystrix-go
        github.com/caarlos0/env
        github.com/go-kit/kit
        github.com/golang-migrate/migrate/v4
        github.com/golang/mock
        github.com/golang/protobuf
        github.com/gorilla/mux
        github.com/jmoiron/sqlx
        github.com/joho/godotenv
        github.com/lib/pq
        github.com/mmcdole/gofeed
        github.com/pkg/errors
        github.com/sirupsen/logrus
        github.com/smartystreets/goconvey
        github.com/sony/gobreaker
        github.com/stretchr/testify
        github.com/victorspringer/http-cache
        google.golang.org/grpc v1.26.0
    ```

    ### 3rd party APIs

    * [urlbox.io](https://urlbox.io/docs)

        A screenshot service allowing to request a web page screeshot via API. 


## `Setup`
Clone this repository and navigate to the folder 
```
https://github.com/michaljirman/newsapp
cd newsapp/
```

run docker containers via docker-compose
```
docker-compose up newsfeedersvc-postgres newsfeedersvc
docker-compose up apigatewaysvc
```

NB: 
docker-compose exposes services on ports 

* 8801 (apigatewaysvc)
* 8802 (newsfeedersvc)
* 15432 (newsfeedersvc-postgres)

Please make sure these port are not taken or change the values in the `docker-compose.yml`

## `API endpoints`

### `/newsfeeder/feeds`
Retrieve list of available feeds
```
curl --request GET --url http://localhost:8801/newsfeeder/feeds
```

```json
{
  "feeds": [
    {
      "feed_id": 1,
      "created_at": "2019-12-26T22:16:24.239Z",
      "modified_at": "2019-12-26T22:16:24.239Z",
      "category": "UK",
      "provider": "BBC",
      "url": "http://feeds.bbci.co.uk/news/uk/rss.xml"
    },
    {
      "feed_id": 2,
      "created_at": "2019-12-26T22:43:08.976Z",
      "modified_at": "2019-12-26T22:43:08.976Z",
      "category": "Technology",
      "provider": "BBC",
      "url": "http://feeds.bbci.co.uk/news/technology/rss.xml"
    },
    {
      "feed_id": 3,
      "created_at": "2019-12-26T22:44:03.088Z",
      "modified_at": "2019-12-26T22:44:03.088Z",
      "category": "UK",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    },
    {
      "feed_id": 4,
      "created_at": "2019-12-27T00:04:53.336Z",
      "modified_at": "2019-12-27T00:04:53.336Z",
      "category": "Techology",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    }
  ]
}
```

### `/newsfeeder/feeds`
Create a new feed
```
curl --request POST --url http://localhost:8801/newsfeeder/feeds --header 'content-type: application/json' --data '{"category": "politics","provider": "BBC","url": "http://feeds.bbci.co.uk/news/politics/rss.xml"}'
```

```json
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /newsfeeder/feeds/6
Date: Fri, 27 Dec 2019 16:51:59 GMT
Content-Length: 13
Connection: close

{
  "FeedID": 5
}
```

### `/newsfeeder/feeds?category=UK`
Retrieve list of available feeds filtered by category
```
curl --request GET --url 'http://localhost:8801/newsfeeder/feeds?category=UK'
```

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 27 Dec 2019 16:54:05 GMT
Content-Length: 387
Connection: close

{
  "feeds": [
    {
      "feed_id": 1,
      "created_at": "2019-12-26T22:16:24.239Z",
      "modified_at": "2019-12-26T22:16:24.239Z",
      "category": "UK",
      "provider": "BBC",
      "url": "http://feeds.bbci.co.uk/news/uk/rss.xml"
    },
    {
      "feed_id": 3,
      "created_at": "2019-12-26T22:44:03.088Z",
      "modified_at": "2019-12-26T22:44:03.088Z",
      "category": "UK",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    }
  ]
}
```



### `/newsfeeder/feeds?provider=Reuters`
Retrieve list of available feeds filtered by category
```
curl --request GET --url 'http://localhost:8801/newsfeeder/feeds?provider=Reuters'
```

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 27 Dec 2019 16:54:40 GMT
Content-Length: 417
Connection: close

{
  "feeds": [
    {
      "feed_id": 3,
      "created_at": "2019-12-26T22:44:03.088Z",
      "modified_at": "2019-12-26T22:44:03.088Z",
      "category": "UK",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    },
    {
      "feed_id": 4,
      "created_at": "2019-12-27T00:04:53.336Z",
      "modified_at": "2019-12-27T00:04:53.336Z",
      "category": "Techology",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    }
  ]
}
```

### `/newsfeeder/feeds?category=UK&provider=Reuters`
Retrieve list of available feeds filtered by category and provider
```
curl --request GET --url 'http://localhost:8801/newsfeeder/feeds?category=UK&provider=Reuters'
```

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 27 Dec 2019 16:55:13 GMT
Content-Length: 211
Connection: close

{
  "feeds": [
    {
      "feed_id": 3,
      "created_at": "2019-12-26T22:44:03.088Z",
      "modified_at": "2019-12-26T22:44:03.088Z",
      "category": "UK",
      "provider": "Reuters",
      "url": "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml"
    }
  ]
}
```

### `/newsfeeder/feeds/1/articles`
Retrieve list of articles for the feed ID
```
curl --request GET --url http://localhost:8801/newsfeeder/feeds/1/articles
```

```json
TTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 27 Dec 2019 16:55:52 GMT
Connection: close
Transfer-Encoding: chunked

{
  "articles": [
    {
      "title": "Christmas Day with a difference across England",
      "description": "Spare a thought for those sea swimming, feeding the lonely and even giving birth on Christmas Day.",
      "link": "https://www.bbc.co.uk/news/uk-england-50823553",
      "published": "2019-12-26T12:04:01Z",
      "guid": "https://www.bbc.co.uk/news/uk-england-50823553",
      "thumbnail_image_url": "https://api.urlbox.io/v1/TFb7yVm5aCzXBXhD/png?url=https%3A%2F%2Fwww.bbc.co.uk%2Fnews%2Fuk-england-50823553&thumb_width=150&ttl=86400"
    },
    ...
    {
      "title": "Darts history-maker Sherrock loses in third round",
      "description": "Fallon Sherrock's challenge at the PDC World Championship is ended in a third-round defeat by world number 22 Chris Dobey.",
      "link": "https://www.bbc.co.uk/sport/darts/50927894",
      "published": "2019-12-27T16:49:00Z",
      "guid": "https://www.bbc.co.uk/sport/darts/50927894",
      "thumbnail_image_url": "https://api.urlbox.io/v1/TFb7yVm5aCzXBXhD/png?url=https%3A%2F%2Fwww.bbc.co.uk%2Fsport%2Fdarts%2F50927894&thumb_width=150&ttl=86400"
    }
  ]
}
```

### `/newsfeeder/feeds/1/articles`
Retrieve a single article for the feed ID by article GUID
```
curl --request POST --url http://localhost:8801/newsfeeder/articles/searches --data '{"feed_id": 1,"article_guid": "https://www.bbc.co.uk/news/uk-england-50823553"}'
```

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 27 Dec 2019 16:56:55 GMT
Connection: close
Transfer-Encoding: chunked

{
  "article": {
    "title": "Christmas Day with a difference across England",
    "description": "Spare a thought for those sea swimming, feeding the lonely and even giving birth on Christmas Day.",
    "link": "https://www.bbc.co.uk/news/uk-england-50823553",
    "published": "2019-12-26T12:04:01Z",
    "guid": "https://www.bbc.co.uk/news/uk-england-50823553",
    "thumbnail_image_url": "https://api.urlbox.io/v1/TFb7yVm5aCzXBXhD/png?url=https%3A%2F%2Fwww.bbc.co.uk%2Fnews%2Fuk-england-50823553&thumb_width=150&ttl=86400",
    "html_content": "Cgo8IURPQ1RZUEUgaHRtbD4KPGh0bWwgbGFuZz0iZW4tR0IiIGlkPSJyZXNwb25zaXZlLW5ld3MiPgo8aGVhZCAgcHJlZml4PSJvZzogaHR0cDovL29ncC5tZS9ucyMiPgogICAgPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZ
    ...
    +CjwhLS0gY29tc2NvcmUgbW14IC0gZW5kIC0tPiAgICAgPC9ib2R5PiA8L2h0bWw+IAoKCgoKCgoKCgoKCg=="
  }
}

```

### Additonal Notes