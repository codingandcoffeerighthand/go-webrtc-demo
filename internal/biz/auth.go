package biz

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v5"
)

type authBiz struct {
	secretKey string
}

type Claims struct {
	WalletAddress string `json:"wallet_address"`
	jwt.RegisteredClaims
}

func NewAuthBiz(secretKey string) *authBiz {
	return &authBiz{
		secretKey: secretKey,
	}
}

func (a *authBiz) GenerateJWT(walletAddress string) (string, error) {
	claims := &Claims{
		WalletAddress: walletAddress,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *authBiz) ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(a.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

type LoginData struct {
	WalletAddress string `json:"wallet_address"`
	Signature     string `json:"signature"`
	Nonce         string `json:"nonce"`
}

func (a *authBiz) Login(data LoginData) (string, error) {
	if !common.IsHexAddress(data.WalletAddress) {
		return "", errors.New("invalid wallet address")
	}

	address := common.HexToAddress(data.WalletAddress)
	sig := common.FromHex(data.Signature)
	if len(sig) != 65 {
		return "", errors.New("invalid signature length")
	}

	if sig[64] >= 27 {
		sig[64] -= 27
	}

	nonce := data.Nonce
	// nextjs using personal.sign: it add "\x19Ethereum Signed Message:\n" + len(nonce) + nonce
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	hash := crypto.Keccak256Hash([]byte(msg))
	// hash := crypto.Keccak256Hash([]byte(nonce))

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)

	if err != nil {
		return "", err
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	if recoveredAddress != address {
		return "", errors.New("signature verification failed")
	}
	token, err := a.GenerateJWT(data.WalletAddress)
	if err != nil {
		return "", err
	}
	return token, nil
}
