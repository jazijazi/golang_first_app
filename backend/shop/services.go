package shop

import (
	"context"
	"encoding/json"
	"fmt"
	"httpproj1/initializers"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetProductService(title string) []byte {
	var result bson.M

	err := initializers.ProductCollection.FindOne(context.TODO(), bson.D{{"title", title}}).
		Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No Document was found")
	}

	jsonData, _ := json.MarshalIndent(result, "", "    ")
	return jsonData
}

func ListProductService() []interface{} {
	filter := bson.D{{}}
	cursor, err := initializers.ProductCollection.Find(context.TODO(), filter)

	var results []interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results

}

func CreateProductService(request *ProductRequest) error { //*ProductRequest
	_, err := initializers.ProductCollection.InsertOne(context.TODO(), request)
	return err
}

func ConsumeProductService(request []byte) error {
	var d interface{}
	if err := json.Unmarshal(request, &d); err != nil {
		panic(err)
	}
	// CreateProductService(d)
	// fmt.Println(request)
	_, err := initializers.ProductCollection.InsertOne(context.TODO(), d)
	if err != nil {
		panic(err)
	}
	return err

}
