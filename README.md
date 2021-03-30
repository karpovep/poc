# PoC

The aim of this project is to build easy-to-scale distributed processing cloud with abilities to support custom types of objects as well as custom implementations of the processing services

## Table Of Contents
0. [Definitions](#Definitions)
1. [PoC - General Idea](#General-Idea)
2. [Technologies](#Technologies)
3. [Acceptance Criteria](#Acceptance-Criteria)
4. [Main Advantages](#Main-Advantages)
4. [Open Questions](#Open-Questions)

### Definitions

- **PoC** - Proof Of Concept
- **Node** - single instance of program (service)


### General Idea

Each Node should implement the following functionality:
* accept configuration to be able to connect to other Nodes
* serve HTTP/2 to accept incoming connections from other Nodes
* accept connections from the clients (same server VS another one - ? TBD)
* subscriptions from clients to objects by type
* daemon functionality for monitoring in-memory objects and transferring objects between instances


### Technologies

- Programming language - GO Lang VS Java ?
- HTTP/2 as a transport, grpc VS capn proto
- Database(s) - TBD, not in scope of PoC


### Acceptance Criteria
Basically, the following test case should pass:
1. Run 2 Nodes
2. Nodes should be able to connect between each other
3. Define 2 simple type of objects (references to each other are not required)
4. Implement simple client which can connect to a Node and subscribe for receiving objects of 1 pre-defined type
5. Run 2 clients: each client connects to different Node
6. Each client should generate an object of another pre-defined type and push it into the program

#### Expected behaviour
* PoC accepts objects of 2 types;
* PoC accepts subscription(s) from the client based on type of the object;
* PoC transfers object(s) between Nodes based on rule - Node does not have registered subscriptions to process object(s);
* Clients are getting notified with unprocessed objects which they were subscribed to


### Main Advantages
- easy to scale
- no synchronisation between nodes, data transferring on-demand
- supporting of custom types of objects to be processed
- failover is handled by each node separately, failing of any Node does not affect entire network of Node(s)
- agnostic implementation - can be used for different processing purposes based on business needs


### Open Questions
- allow to define custom types to clients
- internal objects schemas
- objects state management
