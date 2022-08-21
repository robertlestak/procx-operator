package controllers

import (
	"encoding/json"

	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	"github.com/robertlestak/procx-operator/drivers/activemq"
	"github.com/robertlestak/procx-operator/drivers/aws"
	"github.com/robertlestak/procx-operator/drivers/cassandra"
	"github.com/robertlestak/procx-operator/drivers/centauri"
	"github.com/robertlestak/procx-operator/drivers/elasticsearch"
	"github.com/robertlestak/procx-operator/drivers/fs"
	"github.com/robertlestak/procx-operator/drivers/gcp"
	"github.com/robertlestak/procx-operator/drivers/http"
	"github.com/robertlestak/procx-operator/drivers/kafka"
	"github.com/robertlestak/procx-operator/drivers/mongodb"
	"github.com/robertlestak/procx-operator/drivers/mysql"
	"github.com/robertlestak/procx-operator/drivers/nats"
	"github.com/robertlestak/procx-operator/drivers/nfs"
	"github.com/robertlestak/procx-operator/drivers/nsq"
	"github.com/robertlestak/procx-operator/drivers/postgres"
	"github.com/robertlestak/procx-operator/drivers/pulsar"
	"github.com/robertlestak/procx-operator/drivers/rabbitmq"
	"github.com/robertlestak/procx-operator/drivers/redis"
	"github.com/robertlestak/procx-operator/internal/driver"
	"github.com/robertlestak/procx/pkg/drivers"
)

func unmarshal(s any, dest driver.Driver) driver.Driver {
	jd, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jd, dest)
	if err != nil {
		panic(err)
	}
	return dest
}

func Driver(m *procxv1alpha1.ProcX) driver.Driver {
	switch m.Spec.DriverName {
	case drivers.ActiveMQ:
		return unmarshal(m.Spec.ActiveMQ, &activemq.ActiveMQ{})
	case drivers.AWSDynamoDB:
		return unmarshal(m.Spec.AWSDynamoDB, &aws.AWSDynamoDB{})
	case drivers.AWSSQS:
		return unmarshal(m.Spec.AWSSQS, &aws.AWSSQS{})
	case drivers.AWSS3:
		return unmarshal(m.Spec.AWSS3, &aws.AWSS3{})
	case drivers.CassandraDB:
		return unmarshal(m.Spec.Cassandra, &cassandra.Cassandra{})
	case drivers.Centauri:
		return unmarshal(m.Spec.Centauri, &centauri.Centauri{})
	case drivers.Elasticsearch:
		return unmarshal(m.Spec.Elasticsearch, &elasticsearch.Elasticsearch{})
	case drivers.FS:
		return unmarshal(m.Spec.FS, &fs.FS{})
	case drivers.GCPBQ:
		return unmarshal(m.Spec.GCPBQ, &gcp.GCPBQ{})
	case drivers.GCPGCS:
		return unmarshal(m.Spec.GCPGCS, &gcp.GCPGCS{})
	case drivers.GCPFirestore:
		return unmarshal(m.Spec.GCPFirestore, &gcp.GCPFirestore{})
	case drivers.HTTP:
		return unmarshal(m.Spec.HTTP, &http.HTTP{})
	case drivers.Kafka:
		return unmarshal(m.Spec.Kafka, &kafka.Kafka{})
	case drivers.MongoDB:
		return unmarshal(m.Spec.MongoDB, &mongodb.MongoDB{})
	case drivers.MySQL:
		return unmarshal(m.Spec.MySQL, &mysql.MySQL{})
	case drivers.Nats:
		return unmarshal(m.Spec.NATS, &nats.NATS{})
	case drivers.NFS:
		return unmarshal(m.Spec.NFS, &nfs.NFS{})
	case drivers.NSQ:
		return unmarshal(m.Spec.NSQ, &nsq.NSQ{})
	case drivers.Postgres:
		return unmarshal(m.Spec.Postgres, &postgres.Postgres{})
	case drivers.Pulsar:
		return unmarshal(m.Spec.Pulsar, &pulsar.Pulsar{})
	case drivers.Rabbit:
		return unmarshal(m.Spec.RabbitMQ, &rabbitmq.RabbitMQ{})
	case drivers.RedisList:
		return unmarshal(m.Spec.RedisList, &redis.RedisList{})
	case drivers.RedisSubscription:
		return unmarshal(m.Spec.RedisPubSub, &redis.RedisPubSub{})
	case drivers.RedisStream:
		return unmarshal(m.Spec.RedisStream, &redis.RedisStream{})
	default:
		return nil
	}
}
