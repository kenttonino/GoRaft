# Core Concepts
> - __Distributed System__
>   - Multiple computers working together that appear as one system to the outside world.
> - __Consensus__
>   - The process by which all nodes in a cluster agree on the same value or state.
> - __Raft__
>   - A consensus algorithm that elects one leader to manage all writes and keep nodes in sync.
> - __CAP Theorem__
>   - A rule that says a distributed system can only guarantee two of three: consistency, availability, partition tolerance.
> - __Fault Tolerance__
>   - The ability of a system to keep working even when somes nodes crash or disconnect.
> - __Quorum__
>   - The minimum number of nodes that must agree before a write is considered successful (majority - 2 of 3).
> - __Linearizability__
>   - A consistency guarantee, once a write succeeds, every subsequent read anywhere returns that value.

<br />
<br />
<br />



# Cluster Roles
> - __Leader__
>   - The one nodes elected to accept all writes and replicate them to followers.
> - __Follower__
>   - A passive node that receives log entries from the leader and applies them to its own state.
> - __Leader Election__
>   - The process where nodes vote to pick a new leader when the current one becomes unavailable.
> - __Term__
>   - A logical clock in Raft, increments every election, used to detect state messages.

<br />
<br />
<br />



# Store
> - __KV__ (Key-Value) __Store__
>   - A simple database that a store data as key-value pairs, like a dictionary on disk.
> - __Write-Ahead Log__ (WAL)
>   - A file where every change is recorded before it's applied - so data survives crashes.
> - __Log Replication__
>   - The process of the leader sending its log entries to followers so all nodes have the same data.
> - __Snapshot__
>   - A compressed point-in-time copy of the entire state - so the log doesn't grow forever.
> - __State Machine__
>   - The component that applies committed log entries to the actual data - in GoRaft, this is the KV store.

<br />
<br />
<br />



# Networking
> - __Transmission Control Protocol__(TCP)
>   - A reliable network protocol that guarantees data arrives in order and without loss.
> - __Google Remote Procedural Call__ (gRPC)
>   - A modern framework for making remote function calls between services over the network.
> - __Protobuf__
>   - A compact binary format for serializing structured data - smaller and faster than JSON.
> - __AppendEntries RPC__
>   - The Raft message the leader sends to followers to replicate log entries and prove it's still alive.
> - __Heartbeat__
>   - A periodic empty AppendEntries message the leader sends to stop followers from starting an election.

<br />
<br />
<br />



# Go Concurrency
> - __Goroutine__
>   - A lightweight thread managed by Go — you can run thousands of them concurrently with minimal cost.
> - __Channel__
>   - A Go pipe for safely passing data between goroutines without shared memory.
> - __Mutex__
>   - A lock that ensures only one goroutine can access shared data at a time.
> - __Context__
>   - A Go object for passing cancellation signals and deadlines across goroutines and function calls.
