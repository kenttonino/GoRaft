## (Phase-1) Single-node KV store
- [ ] _Build in-memory KV store with GET, SET, DEL._
- [ ] _Open a raw TCP server with net.Listen._
- [ ] _Parse and handle client commands over TCP._
- [ ] _Handle concurrent connections with goroutines._
- [ ] _Write unit tests for the KV store._

<br />
<br />
<br />





## (Phase-2) Persistence with WAL
- [ ] _Design the WAL entry format (index, term, command)._
- [ ] _Write every command to the WAL file before applying it._
- [ ] _Replay WAL on startup to restore state after a crash._
- [ ] _Test crash recovery — kill the process, restart, verify data is intact._

<br />
<br />
<br //>





## (Phase-3) Raft Consensus
- [ ] _Read the Raft paper fully before writing any code._
- [ ] _Define Raft node states: follower, candidate, leader._
- [ ] _Implement leader election and RequestVote RPC._
- [ ] _Implement AppendEntries RPC for log replication._
- [ ] _Implement heartbeats to prevent unnecessary elections._
- [ ] _Implement log commitment once quorum is reached._
- [ ] _Implement snapshotting to compact the log over time._
- [ ] _Write tests for election, replication, and commit scenarios._

<br />
<br />
<br />





## (Phase-4) 3-node Local Cluster
- [ ] _Set up gRPC transport for node-to-node communication._
- [ ] _Write docker-compose.yml for 3 nodes on different ports._
- [ ] _Spin up the cluster with docker compose up._
- [ ] _Test writes — SET a key, GET it from a different node._
- [ ] _Kill the leader and watch a new one get elected._
- [ ] _Simulate a partition with docker network disconnect._





