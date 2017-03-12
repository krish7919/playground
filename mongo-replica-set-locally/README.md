## Create a MongoDB Replica Set Locally

### Project Goal

*  Aid testing
*  Understand cluster deployments
*  Replicate easily on a cluster management systems

### Procedure

#### Step 1: Build the Latest Code

`make` from the root of this project.


#### Step 2: Start the 3 Instances of MongoDB

Let us assume out replica set is named "test-repl-set".

Since, we are running locally, we need to supply different port
numbers to the mongo containers.

```
docker run \
--name=mdb1 \
--publish=17017:17017 \
--rm=true \
krish7919/mdb \
--replica-set-name test-repl-set \
--port 17017
```

```
docker run \
--name=mdb2 \
--publish=27017:27017 \
--rm=true \
krish7919/mdb \
-replica-set-name test-repl-set \
-port 27017
```

```
docker run \
--name=mdb3 \
--publish=37017:37017 \
--rm=true \
krish7919/mdb \
-replica-set-name test-repl-set \
-port 37017
```

#### Step 3: Initialize the Replica Set

Login to one of the MongoDB containers, say mdb1:

`docker exec -it mdb1 bash`

Start the `mongo` shell:

`mongo --port 17017`


Run the rs.initiate() command:
```
rs.initiate({ 
  _id : "<replica-set-name", members: [
  { 
    _id : 0,
    host : "<eth0 ip of this container>:27017"
  } ]
})
```

For example:

```
rs.initiate({ _id : "test-repl-set", members: [ { _id : 0, host :
"172.17.0.2:17017" } ] })
```

You should also see changes in the mongo shell prompt from `>` to
`test-repl-set:OTHER>` to `test-repl-set:SECONDARY>` to finally
`test-repl-set:PRIMARY>`.
If this instance is not the PRIMARY, you can use the `rs.status()` command to
find out the PRIMARY.

#### Step 4: Add members to the Replica Set

We can only add members to a replica set from the PRIMARY instance.
Login to the PRIMARY and open a `mongo` shell.

Run the rs.add() command with the ip and port number of the other
containers/instances:
```
rs.add("<host>:<port>")
```
For example:
Add mdb2 to replica set from mdb1:
```
rs.add("172.17.0.3:27017")
```
Add mdb3 to replica set from mdb1:
```
rs.add("172.17.0.4:37017")
```

