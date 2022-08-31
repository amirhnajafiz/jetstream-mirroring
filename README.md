<p align="center">
    <img src="./assets/benthos.png" width="318" alt="benthos-pic" />
</p>

<h1 align="center">
Jet-stream Mirroring
</h1>

Mirroring our Nats Jet-stream streams from different 
clusters to each other using Benthos. In this repository I mirrored one stream inside a
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
nats:
  nats1_url: "0.0.0.0:4222" # main service
  nats2_url: "0.0.0.0:4223" # second service
stream:
  stream_name: "snapp" # stream names that are same in both
  subject: "snapp*"    # subjects are also the same
  subject_name: "snapp"
number_of_tests: 20 # number of tests
```

## Set up servers
Use the following command to start our nats js clusters:
```shell
make up
```

Now you have two Jet-stream clusters on:
- 0.0.0.0:4222 (Main)
- 0.0.0.0:4223 (Secondary)

Bring up the benthos on docker:
```shell
make benthos-run-docker
```

If you have benthos installed on your system, you can use
the following command instead, to run benthos on your local system:
```shell
make benthos-run
```

Now you have the benthos service on:
- 0.0.0.0:4195

Now set the streams configs for the service with following command:
```shell
make build
make run
```

## Testing
Test the mirroring by the following command:
```shell
make tests
```

You should be able to see same publications on both clusters.

If you want to shut down the clusters:
```shell
make down
```