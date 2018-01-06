package routers

import (

	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInitRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
