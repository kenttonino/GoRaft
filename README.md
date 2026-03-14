## Description

> - `GoRaft` is a distributed key-value store built in Go, powered by the Raft consensus algorithm.
> - This is a personal project exploring distributed systems fundamentals.

<br />

`High Level Architecture`

```mermaid
graph TD
    Client["Client<br/><small>CLI / Go SDK</small>"]

    Client -->|SET / GET| Leader

    subgraph Cluster ["3-node cluster (localhost)"]
        Leader["Leader node<br/><small>accepts all writes</small>"]
        F1["Follower 1"]
        F2["Follower 2"]

        Leader -->|AppendEntries| F1
        Leader -->|AppendEntries| F2
        F1 -->|ACK| Leader
        F2 -->|ACK| Leader
    end

    subgraph Inside ["Inside each node"]
        Raft["Raft engine<br/><small>consensus</small>"]
        WAL["WAL<br/><small>durability</small>"]
        KV["KV store<br/><small>GET / SET / DEL</small>"]

        Raft --> WAL
        WAL --> KV
    end
```

<br />
<br />
<br />

