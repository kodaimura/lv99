package auth

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	jwtpackage "github.com/golang-jwt/jwt/v5"

	"lv99/config"
	"lv99/internal/core"
)

// JwtAuth implements the AuthI interface using JWT for authentication.
type JwtAuth struct {
	revokedRefreshTokens map[string]time.Time
	mutex                sync.Mutex
}

func NewJwtAuth() core.AuthI {
	return &JwtAuth{
		revokedRefreshTokens: make(map[string]time.Time),
	}
}

type jwtPayload struct {
	jwtpackage.RegisteredClaims
	core.AuthPayload
}

// CreateAccessToken creates a signed JWT containing the given AuthPayload.
func (j *JwtAuth) CreateAccessToken(payload core.AuthPayload) (string, error) {
	return j.createToken(
		payload,
		config.AccessTokenSecret,
		time.Second*time.Duration(config.AccessTokenExpiresSeconds))
}

// CreateRefreshToken creates a signed JWT containing the given AuthPayload.
func (j *JwtAuth) CreateRefreshToken(payload core.AuthPayload) (string, error) {
	return j.createToken(
		payload,
		config.RefreshTokenSecret,
		time.Second*time.Duration(config.RefreshTokenExpiresSeconds),
	)
}

// Common function to generate a JWT token (access or refresh)
func (j *JwtAuth) createToken(payload core.AuthPayload, secretKey string, expiresIn time.Duration) (string, error) {
	now := time.Now()
	tokenId, err := randomTokenId()
	if err != nil {
		return "", err
	}

	jp := jwtPayload{
		AuthPayload: payload,
		RegisteredClaims: jwtpackage.RegisteredClaims{
			Subject:   strconv.Itoa(payload.AccountId),
			ID:        tokenId,
			IssuedAt:  jwtpackage.NewNumericDate(now),
			NotBefore: jwtpackage.NewNumericDate(now),
			ExpiresAt: jwtpackage.NewNumericDate(now.Add(expiresIn)),
		},
	}

	token := jwtpackage.NewWithClaims(jwtpackage.SigningMethodHS256, jp)
	return token.SignedString([]byte(secretKey))
}

// VerifyAccessToken verifies the given JWT and extracts the AuthPayload.
func (j *JwtAuth) VerifyAccessToken(token string) (core.AuthPayload, error) {
	jp, err := j.verifyToken(token, config.AccessTokenSecret)
	if err != nil {
		return core.AuthPayload{}, err
	}
	return jp.AuthPayload, nil
}

// VerifyRefreshToken verifies the given refresh token and extracts the AuthPayload.
func (j *JwtAuth) VerifyRefreshToken(token string) (core.AuthPayload, error) {
	jp, err := j.verifyToken(token, config.RefreshTokenSecret)
	if err != nil {
		return core.AuthPayload{}, err
	}
	if j.isRefreshTokenRevoked(jp.ID) {
		return core.AuthPayload{}, errors.New("refresh token revoked")
	}
	return jp.AuthPayload, nil
}

// Common function to validate a JWT token (access or refresh)
func (j *JwtAuth) verifyToken(token string, secretKey string) (jwtPayload, error) {
	parsedToken, err := jwtpackage.Parse(token, func(t *jwtpackage.Token) (any, error) {
		if _, ok := t.Method.(*jwtpackage.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil || !parsedToken.Valid {
		return jwtPayload{}, err
	}

	return tokenToJwtPayload(parsedToken)
}

func tokenToJwtPayload(token *jwtpackage.Token) (jwtPayload, error) {
	var jp jwtPayload

	claimsMap, ok := token.Claims.(jwtpackage.MapClaims)
	if !ok {
		return jwtPayload{}, errors.New("invalid claims format")
	}

	jsonBytes, err := json.Marshal(claimsMap)
	if err != nil {
		return jwtPayload{}, err
	}

	if err := json.Unmarshal(jsonBytes, &jp); err != nil {
		return jwtPayload{}, err
	}

	return jp, nil
}

func (j *JwtAuth) RevokeRefreshToken(token string) error {
	jp, err := j.verifyToken(token, config.RefreshTokenSecret)
	if err != nil {
		return nil
	}
	if jp.ID == "" || jp.ExpiresAt == nil {
		return nil
	}

	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.purgeExpiredRevocations(time.Now())
	j.revokedRefreshTokens[jp.ID] = jp.ExpiresAt.Time
	return nil
}

func (j *JwtAuth) isRefreshTokenRevoked(tokenId string) bool {
	if tokenId == "" {
		return true
	}

	now := time.Now()
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.purgeExpiredRevocations(now)

	expiresAt, ok := j.revokedRefreshTokens[tokenId]
	return ok && expiresAt.After(now)
}

func (j *JwtAuth) purgeExpiredRevocations(now time.Time) {
	for tokenId, expiresAt := range j.revokedRefreshTokens {
		if !expiresAt.After(now) {
			delete(j.revokedRefreshTokens, tokenId)
		}
	}
}

func randomTokenId() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
