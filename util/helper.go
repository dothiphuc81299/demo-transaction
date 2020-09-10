package util

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/golang/protobuf/ptypes"
)

// HelperParseStringToObjectID ...
func HelperParseStringToObjectID(val string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(val)
	return result
}

// HelperConvertTimeToTimestampProto ...
func HelperConvertTimeToTimestampProto(t time.Time) (*timestamppb.Timestamp) {
	result, _ := ptypes.TimestampProto(t)
	return result
}
