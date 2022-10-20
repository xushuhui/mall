package test

import (
	"context"
	"log"
	"testing"

	spu "mall-go/api/spu/service"
)

func TestListSpuByIds(t *testing.T) {
	reply, err := spuClient.ListSpuByIds(context.Background(), &spu.IdsRequest{Ids: []int64{25, 28}})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(reply)
}
