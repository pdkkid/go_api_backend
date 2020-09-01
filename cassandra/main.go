package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session : Cassandra session
var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("db1.infra.cogrammer.com", "db2.infra.cogrammer.com", "db3.infra.cogrammer.com")
	cluster.Keyspace = "demo"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra Initiated")

}
