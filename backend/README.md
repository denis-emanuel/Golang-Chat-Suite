# BACKEND DOCUMENTATION

The backend project is split into microservices

## Dealing with a new GO microservice

### Create microservice folder inside `/backend`

```bash
    mkdir $newService
    cd $newService
    go mod init github.com/denis-emanuel/Golang-Chat-Suite/$newService
```

### Add new service to the GO workspace

Find the `go.work` file and add the service path

```go
    use (
        ...
        ./$newService
    )
```

Then run the following commands:

```bash
    go work user ./$newService
    go work sync
```

### Verify the new workspace setup

Run the following command and check if the new microservice is in the list

```bash
    go list -m all
```

### Setting up `cobra` and `viper`

#### Install `cobra-cli`

```bash
    go install github.com/spf13/cobra-cli@latest
```

Make sure that you have `$PATH = $GOPATH/bin` inside your `~/.bashrc` or `~/.zshrc` configs.

#### Generate microservice cobra setup

Change directory to `cmd/` where `main.go` is and run

```bash
    cobra-cli add $newService
```

You will see a new file called `$newService` appear inside `cmd/cmd/` folder.
This is where application specific the setup code will lie, for example:

- Environment variable setup
- Kafka connection setup
- Database connection setup
- Running the code

#### Check the new setup

You can add this line to the new file inside the `Run` callback function.

```go
    fmt.Println("New service is starting")
```

Run the new microservice:

```bash
    go run main.go $newService
```
