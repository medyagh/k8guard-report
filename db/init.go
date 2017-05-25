package db

import (
	"github.com/gocql/gocql"
	libs "github.com/k8guard/k8guardlibs"
	libsdb "github.com/k8guard/k8guardlibs/db"
)

// Here we'll store connection
var Sess *gocql.Session
var err error

func InitDB() {
	Sess = libsdb.Connect(libs.Cfg.CassandraHosts)

}
