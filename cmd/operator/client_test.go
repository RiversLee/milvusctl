package operator

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBuildClientConfig(t *testing.T) {
	config,err := BuildClientConfig("","")
	//assert.Equal(t,"",config)
	fmt.Println(config)
	assert.Equal(t,nil,err)
}
func TestGetClientSet(t *testing.T) {
	config,_ := BuildClientConfig("","")
	clientset, err := GetClientSet(config)
	fmt.Println(clientset)
	assert.Equal(t,nil,err)
}