package test

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	app "mall-go/api/app/service"
	spu "mall-go/api/spu/service"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/assert"
)

var (
	appClient app.AppHTTPClient
	spuClient spu.SpuHTTPClient
)

func TestGetBannerById(t *testing.T) {

	reply, err := appClient.GetBannerById(context.Background(), &app.IdRequest{Id: -1})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Id, int64(1))

}
func TestGetBannerByName(t *testing.T) {
	reply, err := appClient.GetBannerByName(context.Background(), &app.NameRequest{Name: "b-1"})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Name, "b-1")

}
func TestGetActivtyByName(t *testing.T) {

	reply, err := appClient.GetActivityByName(context.Background(), &app.NameRequest{Name: "a-2"})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Title, "夏日好礼送不停")

}
func TestGetActivityWithCoupon(t *testing.T) {

	reply, err := appClient.GetActivityWithCoupon(context.Background(), &app.NameRequest{Name: "a-2"})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Title, "夏日好礼送不停")

}

func TestGetThemeByNames(t *testing.T) {
	reply, err := appClient.GetThemeByNames(context.Background(), &app.ThemeByNamesRequest{Names: "t-6,t-3,t-5"})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Theme[0].Title, "风袖甄选")
}

func TestGetThemeByName(t *testing.T) {
	reply, err := appClient.GetThemeByName(context.Background(), &app.NameRequest{Name: "t-1"})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)

}

func TestListCategory(t *testing.T) {
	reply, err := appClient.ListCategory(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, reply.Roots[0].Name, "鞋")
	assert.Equal(t, reply.Subs[0].Name, "平底鞋")
}

func TestListGridCategory(t *testing.T) {
	reply, err := appClient.ListGridCategory(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
	assert.Equal(t, len(reply.Category), 6)
}

func TestMain(m *testing.M) {
	var err error
	appconn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("127.0.0.1:8001"),
	)
	spuconn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithEndpoint("127.0.0.1:8002"),
	)
	if err != nil {
		panic(err)
	}
	appClient = app.NewAppHTTPClient(appconn)
	spuClient = spu.NewSpuHTTPClient(spuconn)

	m.Run()
}
