package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// HelperParseStringToObjectID ...
func HelperParseStringToObjectID(val string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(val)
	return result
}

// HelperParseObjectIDToString ...
func HelperParseObjectIDToString(val primitive.ObjectID) string {
	return val.Hex()
}

// HelperConvertTimeToTimestampProto ...
func HelperConvertTimeToTimestampProto(t time.Time) *timestamppb.Timestamp {
	result, _ := ptypes.TimestampProto(t)
	return result
}
