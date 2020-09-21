The design is displayed in the "system design.jpg" which is next to this file.

There are two designs.
- The first one is a complete design with high scalability support.
- The second one is a simpler one which is just for easy implementation in a limited short period.

I'm here to introduce the complete design.
There are five components in the design. They are:

- **Client**: This is the end user who may use a phone or web browser to use our system.
- **LB**: This is a load balancer. It can be NGINX, LVS or HAProxy, etc. It's used to hide all the internal Gate Server's information.
- **Gate Cluster**: This is used to assign a Server to Client. It's designed to be scalable.
- **Server**: This is the component which serves the Client finally. It's designed to be scalable too.
- **Redis Cluster**: This is used for Gate Cluster and Servers to exchange data.

Let's dive into these components deeper and see how they cooperate to provide a high performance and scalable chat system.
- **Server**: This is the component to provide the chat logic. And it has a room-based design. The reason for this is: In order to scale horizontally.
Server is a stateful process, because it has to maintain all the connected clients and the chat history posted by end users. It's very difficult to scale horizontally. In order to scale horizontally, we need to restrict the state to a certain degree, such as a room. 
One room can scale vertically, but can't scale horizontally. In order to scale the whole system horizontally, each server can host a bunch of rooms, and we can increase or decrease servers. When the number of active users increases which is beyond a single server's capacity, we can add more servers. When it decreases, we can reduce the server's number.
- **Gate Cluster**: This is the component to assign a Server to Client. When a Client sends a request to Gate, it calculates and chooses an appropriate Server and sends back to the Client.
It's designed to be scalable. So I make it stateless. All the data it needs can get from Redis Cluster.
- **Redis Cluster**: This is the component for data sharing. Server stores data into Redis Cluster and Gate fetches data from it.
- **LB**: It's very common to introduce LB into a system, because we can scale the system easily without chaning any endpoint which will affect the end user.