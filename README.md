### PoC (Proof Of Concept)

* it is running as a single process program
* it is possible to run multiple instances of the program
* each instance should implement the following functionality:
    * accept configuration to be able to connect to other instances
    * serve HTTP/2 to accept incoming connections from other instances
    * accept connections from the clients (separate server, TBD)
    * daemon functionality for monitoring in-memory objects and transferring objects between instances
    

### Technologies

- Programmimg language - GO Lang ?
- HTTP/2 as a transport, GRPC protocol
- Database(s) - TBD, not in scope of PoC

### Acceptance criteria
Basically, the following test case should pass:
1. Run 2 instances of the program
2. Instances should be able to connect between each other
3. Define 2 simple type of objects (not necessarily related ones)
4. Implement simple client which connects to 1 configured instance and subscribes for receiving objects of 1 pre-defined type
5. Run 2 instances of client, each client is connected to different instances of program
6. Each client should generate an object of another pre-defined type and push it into the program

Expected behaviour:
PoC accepts objects of 2 types;
PoC accepts subscription(s) from the client based on type of the object;
PoC transfers object between nodes;
Clients are getting notified with unprocessed objects which they were subscribed to
