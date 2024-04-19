package ip

import (
	"fmt"
	"testing"
)

func TestGetIP(t *testing.T) {
	ip, _ := GetIP()
	fmt.Println(ip)
}
