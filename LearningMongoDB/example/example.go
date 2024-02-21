package example

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a type which we can utilise to implement different abilities
type Server struct {
	client *mongo.Client
}

func NewServer(c *mongo.Client) *Server {
	return &Server{
		client: c,
	}
}

// given a server (remember we made this above in the server function) We can now add a function to act on the server
// this is a http handler
func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	// Your going to want to setup you mongoDB
	coll := s.client.Database("catfact").Collection("facts")

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

type CatFactWorker struct {
	client *mongo.Client
}

func NewCatFactWorker(c *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: c,
	}
}

func (cfw *CatFactWorker) start() error {
	// in mongo if we were to push to a collection which didnt exist yet it would be created
	// After some ersearch For our use case we would most likely need a collection for the files
	// and a collection for files, these files need a userId for querying
	// Query for both user-specific files and public files
	// ```
	// cursor, err := filesCollection.Find(context.Background(), bson.M{
	//     "$or": []bson.M{
	//         {"userId": "user123"}, // Filter for user-specific files
	//         {"isPublic": true},    // Filter for public files
	//     },
	// })
	// ```
	coll := cfw.client.Database("catfact").Collection("facts")
	ticker := time.NewTicker(2 * time.Second)

	// some boring api
	for {
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}
		var catFact bson.M // map[string]any // map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		// inserts to the collection
		_, err = coll.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}

		<-ticker.C
	}
}

func example() {
	// connect to the mongodb server!
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client)
	go worker.start()

	server := NewServer(client)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(":3000", nil)
}
