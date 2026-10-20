# Groupie Tracker Visualizations

A Go web application for browsing music artists and bands, then exploring their members, first album, formation year and concert locations with dates.

**Project type:** Individual project, developed by Hasan Abbas.

## Features

- Artist gallery populated from the Groupie Tracker API.
- A detail page for each artist with band members and album information.
- Concert locations paired with their dates through the API's relation data.
- Server-rendered HTML templates, custom CSS and dedicated error pages.

## My work

I built the Go HTTP handlers, fetched and decoded the external JSON data, connected artists to concert information, and created the gallery and detail views with HTML templates and CSS.

## Stack

Go standard library (`net/http`, `encoding/json`, `html/template`), HTML and CSS.

## Run locally

Use Go 1.22.5 or newer and an internet connection for the upstream API:

```bash
git clone https://github.com/hasanjami/groupie-tracker-visualizations.git
cd groupie-tracker-visualizations
go run .
```

Open `http://localhost:8080`.

## Structure

| File or directory | Role |
| --- | --- |
| `main.go` | Routes, gallery/detail handlers and error responses |
| `fetch_data.go` | API requests and data structures |
| `static/` | HTML templates and CSS |

## Check

```bash
go test ./...
go build ./...
```

There are currently no dedicated automated test files. The application depends on the availability of the Groupie Tracker API; artist and concert data belong to their respective sources.
