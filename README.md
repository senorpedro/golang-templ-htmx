# Golang + Templ + HTMLX

Small dummy app for testing those technologies in combination

**Features**

- Go HTTP Server
- Templ templates
- HTMLX
- auto-rebuild via air (server is rebuild on file changes,
  browser still needs to be manually reloaded via CTRL-r)

## Usage

```
go install github.com/cosmtrek/air@latest
go mod tidy
air
```

Then point your browser to http://localhost:8000
