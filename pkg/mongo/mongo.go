package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateClient(
	hosts []string,
	user,
	pass,
	authDB,
	replicaSet string,
	monitor *event.CommandMonitor,
	timeout time.Duration,
) *mongo.Client {
	connStr := buildConnString(hosts, user, pass, authDB, replicaSet)
	logrus.Debugf("Creating new mongoDB connection on %s", hosts)

	opts := options.Client().
		ApplyURI(connStr).
		SetMonitor(monitor)

	if replicaSet != "" {
		opts = opts.SetReadPreference(readpref.SecondaryPreferred())
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logrus.Fatal(err)
	}

	return client
}

func buildConnString(hosts []string, user, pass, authDB, replicaSet string) string {
	hostsDelimited := strings.Join(hosts, ",")

	return fmt.Sprintf("mongodb://%s:%s@%s/%s?replicaSet=%s",
		user, pass, hostsDelimited, authDB, replicaSet)
}

func BsonToRegularInterface(data interface{}) (interface{}, error) {
	var err *error
	ret := bsonToRegularInterface(data, err)

	if err != nil {
		return ret, *err
	}

	return ret, nil
}

func bsonToRegularInterface(data interface{}, err *error) interface{} {
	if err != nil {
		return nil
	}

	if d, ok := data.(bson.D); ok {
		return bsonToRegularInterface(d.Map(), err)
	}

	if m, ok := data.(bson.M); ok {
		ret := make(map[string]interface{}, len(m))
		for k, v := range m {
			ret[k] = bsonToRegularInterface(v, err)
		}

		return ret
	}

	if id, ok := data.(primitive.ObjectID); ok {
		return id.String()
	}

	if dt, ok := data.(primitive.DateTime); ok {
		return dt.Time()
	}

	if orderedArray, ok := data.(primitive.A); ok {
		ret := make([]interface{}, 0, len(orderedArray))
		for _, item := range orderedArray {
			ret = append(ret, bsonToRegularInterface(item, err))
		}

		return ret
	}

	if d, ok := data.(map[string]interface{}); ok {
		ret := make(map[string]interface{}, len(d))
		for k, v := range d {
			ret[k] = bsonToRegularInterface(v, err)
		}

		return ret
	}

	if d, ok := data.([]interface{}); ok {
		ret := make([]interface{}, 0, len(d))
		for _, item := range d {
			ret = append(ret, bsonToRegularInterface(item, err))
		}

		return ret
	}

	if d, ok := data.(string); ok {
		return d
	}

	if d, ok := data.(bool); ok {
		return d
	}

	if d, ok := data.(int32); ok {
		return int(d)
	}

	if d, ok := data.(int64); ok {
		return int(d)
	}

	if d, ok := data.(float64); ok {
		return d
	}

	if data == nil {
		return nil
	}

	*err = fmt.Errorf("could not adapt %v to interface", data)

	return nil
}
