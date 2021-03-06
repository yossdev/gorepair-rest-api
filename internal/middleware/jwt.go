package middleware

import (
	"fmt"
	"gorepair-rest-api/config"
	"gorepair-rest-api/internal/web"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

/*
** Package for handling Auth Middleware using JWT
 */
func JwtVerifyToken(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if jwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]

		if tokenType != "access_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	if err != nil || !token.Valid {
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	return ctx.Next()
}

func JwtVerifyRefresh(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if jwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		fmt.Println(tokenType != "refresh_token")
		if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	if err != nil || !token.Valid {
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	rawID := token.Claims.(jwt.MapClaims)["id"].(float64)

	if rawID == 0 {
		res := web.Response{
			Code:    401,
			Message: "Token not found",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}
	rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
	exp := int64(rawExp)

	if exp < time.Now().Unix() {
		res := web.Response{
			Code:    401,
			Message: "Refresh Token was expired",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	id := strconv.Itoa(int(rawID))
	ctx.Context().Request.Header.Set("ID", id)

	return ctx.Next()
}