package mongodb

import (
	"context"
	"fmt"
	"github.com/SananGuliyev/gossignment/domain/entity"
	"github.com/SananGuliyev/gossignment/domain/io"
	"github.com/SananGuliyev/gossignment/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type recordRepository struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewRecordRepository(database *mongo.Database) repository.RecordRepository {
	return &recordRepository{
		database:   database,
		collection: database.Collection("records"),
	}
}

func (r *recordRepository) FilterByTimeAndAmount(startDate, endDate io.Date, minCount, maxCount int64) ([]*entity.Record, error) {
	var err error

	matchDateStage := bson.D{
		{
			"$match",
			bson.D{
				{
					"$and",
					bson.A{
						bson.D{
							{
								"createdAt",
								bson.D{
									{
										"$gte",
										startDate.Time,
									},
								},
							},
							{
								"createdAt",
								bson.D{
									{
										"$lte",
										endDate.Time,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	addFieldStage := bson.D{
		{
			"$addFields",
			bson.D{
				{"totalCount", bson.D{{"$sum", "$counts"}}},
			},
		},
	}

	matchCountStage := bson.D{
		{
			"$match",
			bson.D{
				{
					"$and",
					bson.A{
						bson.D{
							{
								"totalCount",
								bson.D{
									{
										"$gte",
										minCount,
									},
								},
							},
							{
								"totalCount",
								bson.D{
									{
										"$lte",
										maxCount,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	filterCursor, err := r.collection.Aggregate(
		context.TODO(),
		bson.A{matchDateStage, addFieldStage, matchCountStage},
	)
	if err != nil {
		return nil, err
	}

	var records []*entity.Record
	if err = filterCursor.All(context.TODO(), &records); err != nil {
		return nil, err
	}

	fmt.Println(len(records))

	return records, nil
}
