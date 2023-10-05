package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/riyan-eng/golang-boilerplate-one/env"
	"github.com/riyan-eng/golang-boilerplate-one/infrastructure"
	"golang.org/x/sync/errgroup"
)

type CachedToken struct {
	AccessUID   string `json:"access"`
	RefreshUID  string `json:"refresh"`
	ResetPwdUID string `json:"reset_pwd"`
}

type CustomClaim struct {
	UserID   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

type tokenResult struct {
	token string
	expat *jwt.NumericDate
	uid   string
}

type JwtResult struct {
	AccessToken   string
	RefreshToken  string
	ResetPwdToken string
	ExpiredAt     *jwt.NumericDate
}

func GenerateJwt(userID, roleCode, issuer string) JwtResult {
	access := createToken(userID, roleCode, issuer, env.JWT_SECRET_ACCESS, env.JWT_EXPIRED_ACCESS)
	refresh := createToken(userID, roleCode, issuer, env.JWT_SECRET_REFRESH, env.JWT_EXPIRED_REFRESH)
	cachedJson, err := json.Marshal(CachedToken{
		AccessUID:  access.uid,
		RefreshUID: refresh.uid,
	})
	PanicIfNeeded(err)
	ctx := context.Background()
	infrastructure.Redis.Set(ctx, fmt.Sprintf("token-%s", userID), string(cachedJson), time.Minute*env.JWT_EXPIRED_LOGOFF)
	return JwtResult{
		AccessToken:  access.token,
		RefreshToken: refresh.token,
		ExpiredAt:    access.expat,
	}
}

func GenerateTokenResetPassword(userID, roleCode, issuer string) JwtResult {
	tokenResetPwd := createToken(userID, roleCode, issuer, env.JWT_SECRET_RESET, env.JWT_EXPIRED_RESET)
	cachedJson, err := json.Marshal(CachedToken{
		ResetPwdUID: tokenResetPwd.uid,
	})
	PanicIfNeeded(err)
	ctx := context.Background()
	infrastructure.Redis.Set(ctx, fmt.Sprintf("token-%s", userID), string(cachedJson), time.Minute*env.JWT_EXPIRED_LOGOFF)
	return JwtResult{
		ResetPwdToken: tokenResetPwd.token,
		ExpiredAt:     tokenResetPwd.expat,
	}
}

func createToken(userID, roleCode, issuer, secret string, expMinute time.Duration) tokenResult {
	uid := uuid.NewString()
	expat := jwt.NewNumericDate(time.Now().Add(expMinute * time.Minute))
	claims := CustomClaim{
		userID,
		roleCode,
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: expat,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}
	mySigningKey := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	PanicIfNeeded(err)
	return tokenResult{
		token: ss,
		expat: expat,
		uid:   uid,
	}
}

func ParseToken(tokenString string, secret string) (claims *CustomClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	claims = token.Claims.(*CustomClaim)
	return
}

func ValidateToken(claims *CustomClaim, tokenType string) (err error) {
	ctx := context.Background()
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, _ := infrastructure.Redis.Get(ctx, fmt.Sprintf("token-%s", claims.UserID)).Result()
		cachedTokens := new(CachedToken)
		err := json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string
		if tokenType == "access" {
			tokenUID = cachedTokens.AccessUID
		} else if tokenType == "refresh" {
			tokenUID = cachedTokens.RefreshUID
		} else {
			tokenUID = cachedTokens.ResetPwdUID
		}
		if err != nil || tokenUID != claims.UID {
			return errors.New("token not found")
		}
		return nil
	})

	err = g.Wait()
	return
}
