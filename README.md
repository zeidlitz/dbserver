## dbserver (Database Server)

The goal of dbserver is to be a very simple & lightweight HTTP interface to your database. The aim is to have support for commonly used databases where It should be easy for the user to interface towards multiple different types of connections. 

Example of usage:

- You have some database or persistent storage running and you want to interface to it with HTTP.

- You have some database or persistent storage running and you want a middleware layer between them and some other service. 

- You want to add unnecessary abstraction to your cloud to make a already over-engineer solution even more over-engineered.

### Installation & usage

TODO

### Configuration

TODO

### Testing

The goal is to add some test for each kind of supported database. (Not sure as of yet how realistic that would be). Below shows a example of testing the connection to my trashdatabase. The idea is that as long as the connection method return non nil we consider the test complete. This test might not be feasable to have for all supported database types since it would need a actual database connection to test. But things like queries should be rather doable to mock up and ensure that a given query returns the expected behaviour. 

```go
package connection

import(
  "testing"
  "github.com/zeidlitz/dbserver/internal/trashdatabase"
)

func TestTrashDBConnection(t *testing.T){
  var db trashdatabase.TrashDB
  err := db.Connect("trashconnection")
  if err != nil {
    t.Fatal("Failed to connect to TrashDB")
  }

}
```

This test can be executed from the package itself. (I still don't understand goalngs testing framework well enough to know how to run it from the root directory of the project.) Expected output:

```bash
go tests
2024/05/19 14:11:15 INFO TrashDB connected connection=trashconnection
PASS
ok      github.com/zeidlitz/dbserver/internal/trashdatabase
0.009s
```

### Supported databases

- trashdatabase

- sqlite

- redis
