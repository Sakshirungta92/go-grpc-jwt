package auth

import (
	"authservice/pkg/models"
	"authservice/pkg/utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)
// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var InternalServerError = status.Error(codes.Internal, "internal server error")
var InvalidCredentialsError = status.Error(codes.Unauthenticated, "the request does not have valid authentication credentials for the operation.")

// UserClaims struct as base for JWT payload
type UserClaims struct {
	ID       uint32
	Name	 string
	jwt.StandardClaims
}

// JWT implements auth.IAuth, and used for encoding and validating JWT Token.
type JWT struct{}

// EncodeToken creates a JWT token for a user argument for subsequent call verification.
func (j *JWT) EncodeToken(user *models.User) (*Token, error) {
	currentTime := time.Now()

	// session Time out 3 Days
	expireTime := currentTime.Add(time.Hour * 72).Unix()

	// Create the UserClaims
	claims := UserClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "authservice",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, errors.WithMessage(err, "unable to generate signed string from token.")
	}

	return &Token{
		AccessToken: accessToken,
	}, nil
}

// decodeJWT decodes a JWT token argument into UserClaims object
func (j *JWT) decodeJWT(tokenString string) (*UserClaims, error) {
	userClaims := UserClaims{}
	// Parse the token
	tokenType, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		// log.Error().Err(err).Str("token_string", tokenString).Msg("error while parsing jwt token")
		return nil, InternalServerError
	}

	// Validate the token and return the custom user claims
	if claims, ok := tokenType.Claims.(*UserClaims); ok && tokenType.Valid {
		return claims, nil
	}
	return nil, err
}

// VerifyCredentials called on each and every request made.
// It decodes the JWT token and puts user_id, name and email as principal in the context.
func (j *JWT) VerifyCredentials(ctx context.Context) (context.Context, error) {
	var (
		err         error
		tokenString string
		userClaims  *UserClaims
	)

	tokenString, err = grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		// log.Error().Err(err).Msg("unable to extract token from request header")
		return nil, InvalidCredentialsError
	}

	userClaims, err = j.decodeJWT(tokenString)
	if err != nil {
		// log.Error().Err(err).Msg("unable to decode JWT")
		return nil, InvalidCredentialsError
	}

	md, _ := metadata.FromIncomingContext(ctx)
	newMD := md.Copy()

	md.Set(utils.MDKeyUserID, strconv.FormatUint(uint64(userClaims.ID), 10))
	metadata.NewIncomingContext(ctx, newMD)
	return ctx, nil
}
