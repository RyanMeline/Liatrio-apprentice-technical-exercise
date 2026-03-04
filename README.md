# Liatrio-apprentice-technical-exercise

## Link: https://goapp-529576149401.us-west1.run.app

## Workflow
On push or PR to main
### build

Builds Docker Image<br>
`docker build goapp/ --tag <Docker Username>/liatrio-apprentice:<commit hash>`

Logs into Docker Hub using<br>
 `docker/login-action@v3`

Pushes to Docker Hub<br>
`docker push <Docker Username>/liatrio-apprentice:<commit hash>`

### test
#### Depends on build
Pulls from Docker Hub<br>
`docker pull <Docker Username>/liatrio-apprentice:<commit hash>`

Runs Container<br>
`docker run -d -p 80:80 <Docker Username>/liatrio-apprentice:<commit hash>`<br>
`-d` is used to run it in the background, so the job doesn't get hung up on running the container

Runs Tests using<br>
`liatrio/github-actions/apprentice-action@0b41561cca6822cc8d880fe0e49e7807a41fdf91`

### deploy
#### Depends on tests
Authenticates using service agent key stored in secrets<br>

Pushes to Google Cloud Run directly from Docker Hub<br>
`gcloud run deploy goapp --image docker.io/<Docker Username>/liatrio-apprentice:<commit hash> --region us-west1 --allow-unauthenticated`<br>
`--image <url>` Allows pushing from Docker Hub to Cloud Run<br>
`--regioun us-west1` Sets the region to Oregon<br>
`--allow-unauthenticated` Allows all users to access the site<br>

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
