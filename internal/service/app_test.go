package service

import (
	"context"
	"mall-go/api/mall"
	"mall-go/internal/biz"
	"reflect"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)



func TestAppService_GetBannerById(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.BannerByIdRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mall.Banner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			got, err := s.GetBannerById(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetBannerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppService.GetBannerById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppService_GetBannerByName(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.BannerByNameRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.Banner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetBannerByName(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetBannerByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetBannerByName() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetThemeByNames(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.ThemeByNamesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.Themes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetThemeByNames(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetThemeByNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetThemeByNames() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetThemeWithSpu(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.ThemeWithSpuRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.ThemeSpu
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetThemeWithSpu(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetThemeWithSpu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetThemeWithSpu() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetActivityByName(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.ActivityByNameRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.Activity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetActivityByName(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetActivityByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetActivityByName() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetActivityWithCoupon(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.ActivityWithCouponRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.ActivityCoupon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetActivityWithCoupon(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetActivityWithCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetActivityWithCoupon() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetAllCategory(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *emptypb.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.AllCategory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetAllCategory(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetAllCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetAllCategory() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetGridCategory(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *emptypb.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.GridCategories
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetGridCategory(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetGridCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetGridCategory() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestAppService_GetTagByType(t *testing.T) {
	type fields struct {
		UnimplementedAppServer mall.UnimplementedAppServer
		bu                     *biz.BannerUsecase
		tu                     *biz.ThemeUsecase
		au                     *biz.ActivityUsecase
		log                    *log.Helper
	}
	type args struct {
		ctx context.Context
		in  *mall.TagByTypeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *mall.Tags
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AppService{
				UnimplementedAppServer: tt.fields.UnimplementedAppServer,
				bu:                     tt.fields.bu,
				tu:                     tt.fields.tu,
				au:                     tt.fields.au,
				log:                    tt.fields.log,
			}
			gotOut, err := s.GetTagByType(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppService.GetTagByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("AppService.GetTagByType() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
