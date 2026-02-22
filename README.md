# Liatrio-apprentice-technical-exercise


## Process

### Go app

Init Go app
```
go mod init github.com/RyanMeline/Liatrio-apprentice-technical-exercise
```

Install Fiber
```
go get github.com/gofiber/fiber/v3
```

Returns JSON object with a string and the current UNIX time
```
{
    "message": "My name is Ryan Meline",
    "timestamp": 12341234
}
```

### Dockerfile

Build image
```
docker build -t goapp .
```

Run image
```
docker run -p 3000:3000 goapp
```
