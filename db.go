package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert_PDVs(pdvs []PDV) error {
	if mongo_URL == "" {
		return errors.New("CARBOFRA_MONGO_URL is not set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_URL))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("carbofra").Collection("pdvs")
	docs := make([]interface{}, len(pdvs))
	for i, pdv := range pdvs {
		docs[i] = pdv
	}
	err = collection.Drop(context.TODO())
	if err != nil {
		return err
	}
	result, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}
	
	fmt.Printf("Inserted %d documents\n", len(result.InsertedIDs))
	return nil
}
