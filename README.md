# JetStream Mirroring using Benthos

Mirroring our NATS (JetStream) cluster from different 
regions to each other by using __Benthos__.
In this repository I mirrored one stream inside a JetStream service to another stream inside a different
JetStream cluster. I used Benthos to do this message duplicating.

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

### Output

Output is our back-up nats js service, which will be mirrored
by the main service.

```yaml
output:
  label: "nats2"
  nats_jetstream:
    urls: [ "nats://0.0.0.0:4223" ]
    subject: "snapp*"
```

## Setup

Use the following command to start Benthos, two NATS clusters, bootstrap service, two consumers, and one publisher
to see the mirroring between clusters:

```shell
docker-compose up
```