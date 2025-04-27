package shop

import (
	"context"
	"encoding/json"
	"fmt"
	"httpproj1/initializers"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetProductService(title string) (bson.M, error) {
	var result bson.M

	//bson is Binary JSON, MongoDB's internal format.
	//bson.D  ---->   Ordered list of key-value pairs to form MongoDB queries
	//bson.M  ---->   is just a shorthand for map[string]interface{}
	//bson.M  ---->   keys are string and values are interface{} (could be string, int, bool, array, object, anything)
	err := initializers.ProductCollection.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product with title '%s' not found", title)
		}
		return nil, err
	}

	return result, nil
}

func ListProductService(ctx context.Context) ([]bson.M, error) {
	filter := bson.D{{}}

	cursor, err := initializers.ProductCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // good practice!

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
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
	_, err := initializers.ProductCollection.InsertOne(context.TODO(), d)
	if err != nil {
		panic(err)
	}
	return err

}
