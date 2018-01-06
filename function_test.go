package apis

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetDataList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetDataList(tt.args.c)
		})
	}
}

func TestPageNextData(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PageNextData(tt.args.c)
		})
	}
}

func TestAddPersonApi(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddPersonApi(tt.args.c)
		})
	}
}

func TestEditHtml(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EditHtml(tt.args.c)
		})
	}
}

func TestEditPersonApi(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EditPersonApi(tt.args.c)
		})
	}
}

func TestDeletePersonApi(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeletePersonApi(tt.args.c)
		})
	}
}
