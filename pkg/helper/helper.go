package helper

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	cfg "github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/helper/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	config cfg.Config
}

func NewHelper(cfg cfg.Config) interfaces.Helper {
	return &helper{
		config: cfg,
	}
}

func (h *helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (h *helper) Copy(a interface{}, b interface{}) (interface{}, error) {
	err := copier.Copy(a, b)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (helper *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error) {
	accessTokenClaims := &models.AuthCustomClaims{
		Id:        admin.ID,
		Email:     admin.Email,
		Previlege: admin.Previlege,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := &models.AuthCustomClaims{
		Id:        admin.ID,
		Email:     admin.Email,
		Previlege: admin.Previlege,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("accesssecret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("refreshsecret"))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (h *helper) CreateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) UploadToS3(file *multipart.FileHeader) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		fmt.Println("configuration error:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("procasto"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		fmt.Println("uploading error:", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}

func (helper *helper) GenerateTokenProvider(details domain.Provider) (string, error) {
	accessTokenClaims := &models.AuthCustomClaims{
		Id:        details.ID,
		Email:     details.Email,
		Previlege: "Pro",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("accesssecret"))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
