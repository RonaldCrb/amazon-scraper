# SRE Code Challenge - JustWatch, Berlin

This Application has been created by Ronald Alonzo with the purpose of solving the Code Challenge for the company JustWatch and succesfully apply for an Site Reliability Engineer position there. The requirements have been designed as per the guidelines in the `SRE Challenge.pdf` document present on this repository.

## Access
As per recommended guidelines from JustWatch, this repository is private and has been shared with the accounts `falschparker82` and `zwopir`

## Instructions

This is an HTTP server that exposes 2 endpoints.

1. HealthCheck `GET => http://localhost:8080/healtz`
```json
{
  "status": "amazon-scraper is Healthy!",
  "timestamp": "2020-08-27 13:02:05"
}
```

2. Amazon Movie data scrape `GET => http://localhost:8080/movie/amazon/{amazon_id}`
```json
{
  "title": "Der Biber",
  "release_year": 2011,
  "actors": [
    "Mel Gibson",
    "Jodie Foster",
    "Anton Yelchin"
  ],
  "poster": "https://images-na.ssl-images-amazon.com/images/I/81jpeIbNEbL._SX300_.jpg",
  "similar_ids": [
    "B00HDXNJJ6",
    "B00FCLTFCG",
    "B00FCLKWVO",
    "B083QMSYH5",
    "B00OL2XW94",
    "B00FYWWYXU",
    "B00OL3T7X8",
    "B00PWYZVO4",
    "B00IF47I62",
    "B07RDMM46C",
    "B00FXKK0VQ",
    "B00WBZTX50",
    "B00FCLE2SI",
    "B01N55HQIE",
    "B00NP9U97C",
    "B00TP12LW0",
    "B081X9PR31",
    "B075ZBNCK7",
    "B00FXL57YA",
    "B0789T3NRL"
  ]
}
```
