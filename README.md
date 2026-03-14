## Description

> - `GoRaft` is a distributed key-value store built in Go, powered by the Raft consensus algorithm.
> - This is a personal project exploring distributed systems fundamentals.
> - Check the `/docs/*` for further information.

<br />
<br />
<br />



## Local Setup

> - Run the following commands.

| `Script` | `Description` |
| -------- | ------------- |
| (req) __make run-server__ | Start the GoRaft server. |
| (req) __make run-client__ | Connect as a client to the server. |

<br />

> - Sample input.

```sh
# Save the key-value pair in Store.
SET project GoRaft

# Get the value based on the key.
GET project

# Delete the value based on the key.
DEL project

# Check if the value is already deleted using the key.
GET project
```
