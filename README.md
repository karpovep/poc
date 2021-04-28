# [Distributed Processing Cloud - PoC](https://distributed-processing-cloud.firebaseapp.com/)

![10%](https://progress-bar.dev/10)

The aim of this project is to build easy-to-scale distributed processing cloud with abilities to support custom types of objects as well as custom implementations of the processing services to be able to meet all of the required business needs

## Table Of Contents
0. [Definitions](#Definitions)
1. [General Idea](#General-Idea)
2. [Technologies](#Technologies)
3. [Acceptance Criteria](#Acceptance-Criteria)
4. [Main Advantages](#Main-Advantages)
5. [Open Questions](#Open-Questions)
6. [Roadmap](#Roadmap)
7. [Documentation](docs/index.md)
8. [Diagrams](#Diagrams)

### Definitions
- **PoC** - Proof Of Concept
- **Node** - single instance of program (service)


### General Idea
Each Node should implement the following functionality:
* accept configuration to be able to connect to other Nodes
* serve HTTP/2 to accept incoming connections from other Nodes
* accept connections from the clients - separate server
* subscriptions from clients to objects by type
* daemon functionality for monitoring in-memory objects and transferring objects between instances
* repository - to save data to
* transactional processing


### Technologies
- Programming language - GO Lang
- HTTP/2 as a transport, grpc
- Database(s) - Cassandra as main storage, Postgres as an option for transactional processing - TBD


### Acceptance Criteria
:eight_spoked_asterisk: Basically, the following test case should pass:
1. Run 2 Nodes
2. Nodes has to be able to connect between each other
3. Define 2 simple type of objects (references to each other are not required)
4. Implement simple client which can connect to a Node and subscribe for receiving objects of 1 pre-defined type
5. Run 2 clients: each client connects to different Node
6. Each client should generate an object of another pre-defined type and push it into the cloud

#### Expected behaviour
* :white_check_mark: PoC accepts objects of 2 types;
* :white_check_mark: PoC accepts subscription(s) from the client based on type of the object;
* :white_check_mark: PoC transfers object(s) between Nodes based on rule - Node does not have registered subscriptions to process object(s);
* :white_check_mark: Clients are getting notified with unprocessed objects which they were subscribed to


### Main Advantages
- easy to scale
- no synchronisation between nodes, data transferring on-demand
- supporting of custom types of objects to be processed
- failover is handled by each node separately, failing of any Node does not affect entire network of Node(s)
- agnostic implementation - can be used for different processing purposes based on business needs


### Open Questions
- :large_orange_diamond: internal objects schemas
- :large_orange_diamond: objects state management
- :x: allow to define custom types to clients
- :x: example to be implemented for the demo purpose

### Roadmap
- :white_check_mark: Implement simplified PoC (1 server node and 1 client)
- :white_check_mark: Implement PoC (2 nodes + 2 clients)
- :large_orange_diamond: introduce data storage (store data + querying data)
- :x: Failover handling
- :x: Logging and monitoring
- :x: Transactional processing
- :x: Reporting & Analytics
- ...
- :x: Centralised cloud network

### Diagrams
- [Basic Architecture View](https://drive.google.com/file/d/1ukPn3U78vHxhr7BJNcWFetokQS_1pMXa/view)
- [Server Node Components](https://drive.google.com/file/d/1JG-yAHjmxeNS6PgxwnjE62t4KoFMdgH5/view)
- [Basic Sequence Diagram](https://drive.google.com/file/d/1AGZXQFtNuUlxJsOziDhPfv7i8YBmQfeR/view)
- [Data Structures](https://drive.google.com/file/d/1juhmjO4wt3YYu_EDF68281vRpm4O9MzJ/view)