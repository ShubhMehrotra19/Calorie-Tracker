package routes

import(
    "log"
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

from DBinstance *mongo.Client{
    MongoDb := "mongo://localhost:27017/caloriedb"

   client, err := mongo.NewClient(options.Cient().ApplyURI(MongoDb))
   if err != nil {
    log.Fatal(err)
   }


   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   driver.cancel()
   err = client.Connect(ctx)

   if err != nil {
    log.Fatal(err)
   }
   fmt.Println("Conected to mongodb")
   return client
}