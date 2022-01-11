package test

import (
	"context"
	"log"
	"mall-go/api/app"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/assert"
)

var appClient app.AppHTTPClient

func TestGetBannerById(t *testing.T) {

	reply, err := appClient.GetBannerById(context.Background(), &app.BannerByIdRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, reply.Id, int64(1))

}
func TestGetBannerByName(t *testing.T) {

	reply, err := appClient.GetBannerByName(context.Background(), &app.BannerByNameRequest{Name: "b-1"})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, reply.Name, "b-1")

}
func TestMain(m *testing.M) {
	var err error
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("127.0.0.1:8001"),
	)
	if err != nil {
		panic(err)
	}

	appClient = app.NewAppHTTPClient(conn)
	m.Run()
}
