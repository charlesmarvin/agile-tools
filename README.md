# agile-tools

### Build and run locally

Creates a database called `agile-tools` in the mongodb specified in the `MONGO_URL` env

```terminal
% cd app
% go build
% MONGO_URL=mongodb://localhost:27017 ./agile-tools
```

### Build and run Docker image
```terminal
% docker build -t agile-tools .
% docker run -it -p 8080:8080 -e MONGO_URL=mongodb://host.docker.internal:27017  agile-tools
```

### Live Reload
Note: There is a bug in Modd v8.0 so we need to install from `HEAD` until the next release see https://github.com/cortesi/modd/issues/69
```zsh
% brew install --HEAD modd
% modd
```