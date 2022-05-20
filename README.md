### Feature
- 拆分多个子订单
- 退款
- 积分
- 钱包充值和余额支付
- 商品收藏
- 商品关联推荐


## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## 微服务
### mall interface
外网接口，BFF
### app service
主业务服务，banner/theme/activity
### sku service
商品服务，sku/spu商品查询，创建，更新，商品推荐
### user service
用户服务，积分，用户优惠券，用户信息，用户账户，实名认证，地址，收藏
### order service
订单服务，退款，下单，售后，退换货
### payment service
支付服务，扫码支付，小程序支付，app支付，查询账单，退款