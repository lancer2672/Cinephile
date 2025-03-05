package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/lancer2672/cinephile/main/config"
	"github.com/lancer2672/cinephile/main/log"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoStore struct {
	client     *mongo.Client
	db         *mongo.Database
	session    mongo.Session
	sessionCtx mongo.SessionContext
}

// NewMongoStore initializes and returns a new MongoStore instance
func NewMongoStore(cfg *config.Config) (*MongoStore, error) {
	// Context with timeout for DB connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create MongoDB client options with logging capabilities
	clientOptions := buildClientOptions(cfg)

	// Initialize MongoDB client
	client, err := connectToMongoDB(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// // Create index
	// err = createUniqueIndex(cfg, client.Database(cfg.DBName))
	// if err != nil {
	// 	return nil, errors.New("failed to create unique index: " + err.Error())
	// }

	// Ensure connection is alive
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("failed to ping MongoDB: " + err.Error())
	}

	log.Logger.Info("✅ Connected to MongoDB!")

	db := client.Database(cfg.DBName)

	return &MongoStore{
		client: client,
		db:     db,
	}, nil
}

// buildClientOptions sets up MongoDB client options including logging
func buildClientOptions(cfg *config.Config) *options.ClientOptions {
	cmdMonitor := &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			log.Logger.Debug("Command started",
				zap.String("command", evt.CommandName),
				zap.Any("details", evt.Command),
			)
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			log.Logger.Debug("Command succeeded",
				zap.String("command", evt.CommandName),
				zap.Duration("duration", evt.Duration),
			)
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			log.Logger.Error("Command failed",
				zap.String("command", evt.CommandName),
				zap.Duration("duration", evt.Duration),
			)
			log.Logger.Error("Failure details",
				zap.Any("failure", evt.Failure),
			)
		},
	}

	return options.Client().
		ApplyURI(cfg.DBSource).
		SetMonitor(cmdMonitor)
}

// connectToMongoDB connects to the MongoDB client and handles errors
func connectToMongoDB(ctx context.Context, opts *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.New("failed to connect to MongoDB: " + err.Error())
	}
	return client, nil
}

func (m *MongoStore) Disconnect() error {
	err := m.client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	log.Logger.Info("Connection to MongoDB closed.")
	return nil
}
