package client

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

var (
	session *gocql.Session
)

func CreateSession() {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "meli"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}

func CreateKeyspace(keyspace string) {
	c := gocql.NewCluster("127.0.0.1")
	c.Keyspace = "system"
	c.Timeout = 30 * time.Second
	flagRF := 1
	session, err := c.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.Query(`DROP KEYSPACE IF EXISTS ` + keyspace).Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = session.Query(fmt.Sprintf(`CREATE KEYSPACE %s
    WITH replication = {
        'class' : 'SimpleStrategy',
        'replication_factor' : %d
    }`, keyspace, flagRF)).Exec()

	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(){
	err := session.Query(`CREATE TABLE movies(id  timeUUID PRIMARY KEY, user_id bigint, title text, length int, synopsis text,image_url text)`).Exec()

	if err != nil {
		log.Fatal(err)
	}
}
