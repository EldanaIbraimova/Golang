package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Fact struct {
	ID primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Categories []string `bson:"categories"`
	Text string `bson:"text"`

}

type FactsFilter struct {
	Query *string `json:"query"`
}