package controllers

import (
	"encoding/json"

	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	"github.com/robertlestak/procx-operator/drivers/aws"
	"github.com/robertlestak/procx-operator/drivers/cassandra"
	"github.com/robertlestak/procx-operator/drivers/centauri"
	"github.com/robertlestak/procx-operator/drivers/elasticsearch"
	"github.com/robertlestak/procx-operator/drivers/gcp"
	"github.com/robertlestak/procx-operator/drivers/kafka"
	"github.com/robertlestak/procx-operator/drivers/mongodb"
	"github.com/robertlestak/procx-operator/drivers/mysql"
	"github.com/robertlestak/procx-operator/drivers/nats"
	"github.com/robertlestak/procx-operator/drivers/nfs"
	"github.com/robertlestak/procx-operator/drivers/nsq"
	"github.com/robertlestak/procx-operator/drivers/postgres"
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
	case drivers.DriverAWSDynamoDB:
		return unmarshal(m.Spec.AWSDynamoDB, &aws.AWSDynamoDB{})
	case drivers.DriverAWSSQS:
		return unmarshal(m.Spec.AWSSQS, &aws.AWSSQS{})
	case drivers.DriverAWSS3:
		return unmarshal(m.Spec.AWSS3, &aws.AWSS3{})
	case drivers.DriverCassandraDB:
		return unmarshal(m.Spec.Cassandra, &cassandra.Cassandra{})
	case drivers.DriverCentauriNet:
		return unmarshal(m.Spec.Centauri, &centauri.Centauri{})
	case drivers.DriverElasticsearch:
		return unmarshal(m.Spec.Elasticsearch, &elasticsearch.Elasticsearch{})
	case drivers.DriverGCPBQ:
		return unmarshal(m.Spec.GCPBQ, &gcp.GCPBQ{})
	case drivers.DriverGCPGCS:
		return unmarshal(m.Spec.GCPGCS, &gcp.GCPGCS{})
	case drivers.DriverKafka:
		return unmarshal(m.Spec.Kafka, &kafka.Kafka{})
	case drivers.DriverMongoDB:
		return unmarshal(m.Spec.MongoDB, &mongodb.MongoDB{})
	case drivers.DriverMySQL:
		return unmarshal(m.Spec.MySQL, &mysql.MySQL{})
	case drivers.DriverNats:
		return unmarshal(m.Spec.NATS, &nats.NATS{})
	case drivers.DriverNFS:
		return unmarshal(m.Spec.NFS, &nfs.NFS{})
	case drivers.DriverNSQ:
		return unmarshal(m.Spec.NSQ, &nsq.NSQ{})
	case drivers.DriverPostgres:
		return unmarshal(m.Spec.Postgres, &postgres.Postgres{})
	case drivers.DriverRabbit:
		return unmarshal(m.Spec.RabbitMQ, &rabbitmq.RabbitMQ{})
	case drivers.DriverRedisList:
		return unmarshal(m.Spec.RedisList, &redis.RedisList{})
	case drivers.DriverRedisSubscription:
		return unmarshal(m.Spec.RedisPubSub, &redis.RedisPubSub{})
	case drivers.DriverRedisStream:
		return unmarshal(m.Spec.RedisStream, &redis.RedisStream{})
	default:
		return nil
	}
}
