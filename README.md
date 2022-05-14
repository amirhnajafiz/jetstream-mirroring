# Jet-stream Mirroring

Mirroring our Nats Jet-stream streams from different 
clusters to each other using Benthos.

<img src="./assets/benthos.png" width="700" />

## What is this repository for?
In this repository I mirrored one stream inside a
Jet-stream service to another stream inside a different
Jet-stream cluster.

I used Benthos to do this message duplicating.

## Benthos configs
### Input
Input is our first nats js service that is the main jet-stream cluster.
We config our benthos input based of the configs in our 
project.
```yaml
input:
  label: "nats1"
  nats_jetstream:
    urls: [ "nats://0.0.0.0:4222" ]
    queue: ""
    subject: "snapp*"
    durable: ""
    stream: "snapp"
    bind: false
    deliver: all
```

Output is our back-up nats js service, which will be mirrored
by the main service.
```yaml
output:
  label: "nats2"
  nats_jetstream:
    urls: [ "nats://0.0.0.0:4223" ]
    subject: "snapp*"
```

## Set configs
Set our services configs:
```shell
cp ./configs/example-config.yaml ./config.yaml
```

You can set the project configs in this file:
```yaml
nats1: "0.0.0.0:4222" # main service
nats2: "0.0.0.0:4223" # second service
stream_name: "snapp"  # stream names that are same in both
subject: "snapp*"     # subjects are also the same
```

## Testing
Use the following command to start our nats js clusters:
```shell
make up
```

Now you have two Jet-stream clusters on:
- 0.0.0.0:4222 (Main)
- 0.0.0.0:4223 (Secondary)

Bring up the benthos:
```shell
make benthos
```

Now you have the benthos service on:
- 0.0.0.0:4195

Now you can test the service with following command:
```shell
make build
make run
```

If you check the nats containers logs, you can see the results
of the test.

You should be able to see same publications on both clusters.

If you want to shut down the clusters:
```shell
make down
```