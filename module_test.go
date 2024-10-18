package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/huaweicloud"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`huaweicloud {
			access_key_id thekey
			secret_access_key itsasecret
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&huaweicloud.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedAccessKeyId := "thekey"
			actualAccessKeyId := p.Provider.AccessKeyId
			if expectedAccessKeyId != actualAccessKeyId {
				t.Errorf("Expected AccessKeyId to be '%s' but got '%s'", expectedAccessKeyId, actualAccessKeyId)
			}

			expectedSecretAccessKey := "itsasecret"
			actualSecretAccessKey := p.Provider.SecretAccessKey
			if expectedSecretAccessKey != actualSecretAccessKey {
				t.Errorf("Expected SecretAccessKey to be '%s' but got '%s'", expectedSecretAccessKey, actualSecretAccessKey)
			}
		})
	}
}
