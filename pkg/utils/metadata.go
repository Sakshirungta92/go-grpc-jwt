package utils

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var InternalServerError = status.Error(codes.Internal, "internal server error")

// MDGet returns the metadata object present in the incoming context
func MDGet(ctx context.Context) (metadata.MD, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md, nil
	}
	return nil, InternalServerError
}

// MDGetValue find the value for which key has been passed in the arguments from the provided context
func MDGetValue(ctx context.Context, key string) (string, error) {
	md, err := MDGet(ctx)
	if err != nil {
		return "", err
	}

	if values := md.Get(key); len(values) > 0 {
		return values[0], nil
	}
	// log.Error().Msgf("error while getting %s from metadata", key)

	return "", InternalServerError
}