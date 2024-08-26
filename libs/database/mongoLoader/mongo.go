package mongoLoader

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Ssl      bool
}

func NewConnection(cfg MongoConfig) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s:%s`, cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	if cfg.Ssl {
		uri = fmt.Sprintf(`mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&readPreference=secondaryPreferred`, cfg.Username, cfg.Password, cfg.Host)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	clientOpts := options.Client()
	client, err := mongo.Connect(ctx, clientOpts.ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("MONGO", err)
		return nil, err
	}

	db := client.Database(cfg.DBName)
	return db, nil
}
