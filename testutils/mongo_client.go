package testutils

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultStandaloneMongoDBServerURL = "mongodb://localhost:27027"
	DefaultShardedMongoDBServerURL    = "mongodb://localhost:27037"
	DefaultReplSetMongoDBServerURL    = "mongodb://localhost:27047"
)

// MustGetConnectedReplSetClient return mongo.Client instance connected to server started in replicaSet mode.
func MustGetConnectedReplSetClient(ctx context.Context, t *testing.T) *mongo.Client {
	opts := options.Client().
		ApplyURI(DefaultReplSetMongoDBServerURL + "/admin").
		SetReplicaSet("rs3").
		SetDirect(true).SetServerSelectionTimeout(time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		t.Fatalf("Couldn't connect to MongoDB instance, reason: %v", err)
	}

	return client
}

// MustGetConnectedMongodClient return mongo.Client instance connected to server started in single mode.
func MustGetConnectedMongodClient(ctx context.Context, t *testing.T) *mongo.Client {
	opts := options.Client().
		ApplyURI(DefaultStandaloneMongoDBServerURL + "/admin").
		SetDirect(true).SetServerSelectionTimeout(time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		t.Fatalf("Couldn't connect to MongoDB instance, reason: %v", err)
	}

	return client
}
