package test

import (
	"context"
	"fmt"
	"log"
	"mall-go/api/app"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/assert"
)

var appClient app.AppHTTPClient

func TestGetBannerById(t *testing.T) {

	reply, err := appClient.GetBannerById(context.Background(), &app.IdRequest{Id: -1})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Id, int64(1))

}
func TestGetBannerByName(t *testing.T) {
	reply, err := appClient.GetBannerByName(context.Background(), &app.NameRequest{Name: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	assert.Equal(t, reply.Name, "b-1")

}
func TestGetActivtyByName(t *testing.T) {

	reply, err := appClient.GetActivityByName(context.Background(), &app.NameRequest{Name: "a-2"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	assert.Equal(t, reply.Title, "夏日好礼送不停")

}
func TestGetActivityWithCoupon(t *testing.T) {

	reply, err := appClient.GetActivityWithCoupon(context.Background(), &app.NameRequest{Name: "a-2"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	assert.Equal(t, reply.Title, "夏日好礼送不停")

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
