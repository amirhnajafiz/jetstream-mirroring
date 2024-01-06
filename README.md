# JetStream Mirroring

Mirroring our NATS (JetStream) cluster from different 
regions to each other, using __Benthos__.
In this project, I provided an example in which I mirrored one stream inside a JetStream cluster to another stream
inside a different JetStream cluster.
In order to do this, I used __Benthos__ as a third party service between the two clusters.

## Benthos configs

Benthos is a high performance and resilient stream processor, able to connect
various sources and sinks in a range of brokering patterns and perform hydration, enrichments,
transformations and filters on payloads.
Read more about [Benthos](https://github.com/benthosdev/benthos).

### Input

Input is our first NATS (JetStream) service that is the main cluster.

```yaml
input:
  label: "nats1"
  nats_jetstream:
    urls: [ "nats://0.0.0.0:4222" ]
    queue: ""
    subject: "snapp*" # subject
    durable: ""
    stream: "snapp" # stream
    bind: false
    deliver: all
```

### Output

Output is the other NATS (JetStream) service, which will be mirrored
by the main service.

```yaml
output:
  label: "nats2"
  nats_jetstream:
    urls: [ "nats://0.0.0.0:4223" ]
    subject: "snapp*"
```

## Example Setup

Use the following command to start Benthos, two NATS clusters, bootstrap service, two consumers, and one publisher
to see the mirroring between clusters:

```shell
docker-compose up
```
