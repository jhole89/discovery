package connectors

import (
	"database/sql"
	"log"
	"strings"
)

type Driver interface {
	Query(query string) (*sql.Rows, error)
	Index() ([]*Node, error)
}

type Indexer interface {
	getDatabases() ([]string, error)
	getTables(database string) ([]string, error)
	describeTables(database, table string) ([]Column, error)
}

type Node struct {
	Name       string
	Context    string
	Properties map[string]string
	Children   []*Node
}

func GetDriver(name string, address string) Driver {

	var supportedConnectors = map[string]func(dsn string) (Driver, error){
		"awsathena": NewAwsAthena,
	}

	c, ok := supportedConnectors[strings.ToLower(name)]

	if ok {
		d, _ := c(address)
		return d
	} else {
		keys := make([]string, len(supportedConnectors))
		for k := range supportedConnectors {
			keys = append(keys, k)
		}
		log.Printf("Connecting to %s is not supported. Please specifiy a supported connector in your config.yaml.\nValid connectors's: %s", name, keys)
		return nil
	}
}
