#mgowrapper

Go package to wrap mongo database connection and queries

##Usage

Implement interfaces for your structs to use the wrapper for querying, inserting, updating and deleting

```go

package main

import (
  "github.com/clanbeat/mgowrapper"
  "log"
)


func main() {
  var mgo *mongo.Mongo
  var err error
  //set up mongo connection
  mgo, err = mongo.New(os.Getenv("MONGODB_URI"))
  defer mgo.Close()
  if err != nil {
    log.Fatal(err)
  }
  if err := mgo.EnsureIndexes(getMyStructIndexes()); err != nil {
    log.Fatal(err)
  }
}

```
