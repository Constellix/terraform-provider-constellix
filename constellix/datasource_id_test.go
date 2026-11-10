package constellix

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Constellix/constellix-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestDataSourceIDs(t *testing.T) {
	tests := []struct {
		name       string
		resource   *schema.Resource
		path       string
		attributes string
		want       map[string]interface{}
	}{
		{
			name:     "tags",
			resource: datasourceConstellixTags(),
			path:     "/v2/tags",
		},
		{
			name:       "template",
			resource:   datasourceConstellixTemplate(),
			path:       "/v1/templates",
			attributes: `,"Domain":123456,"version":1000001,"hasGeoIP":true,"hasGtdRegions":false`,
			want:       map[string]interface{}{"domain": 123456, "version": 1000001, "has_geoip": true},
		},
		{
			name:       "geo_proximity",
			resource:   datasourceConstellixGeoProximity(),
			path:       "/v1/geoProximities",
			attributes: `,"country":"US","region":"VA","latitude":38.9,"longitude":-77.1`,
			want:       map[string]interface{}{"latitude": 38.9, "longitude": -77.1},
		},
		{
			name:       "cname_record",
			resource:   datasourceConstellixCNameRecord(),
			path:       "/v1/domains/123456/records/cname/",
			attributes: `,"host":"target.example.","ttl":300,"pools":[1000001],"contactId":[123]`,
			want: map[string]interface{}{
				"host":        "target.example.",
				"ttl":         300,
				"pools":       []interface{}{1000001},
				"contact_ids": []interface{}{123},
			},
		},
		{
			name:       "a_record",
			resource:   datasourceConstellixARecord(),
			path:       "/v1/domains/123456/records/a/",
			attributes: `,"ttl":300,"roundRobin":[],"roundRobinFailover":[{"value":"192.0.2.1","sortOrder":1,"disableFlag":false}]`,
			want: map[string]interface{}{
				"ttl":                                300,
				"roundrobin_failover.0.value":        "192.0.2.1",
				"roundrobin_failover.0.sort_order":   "1",
				"roundrobin_failover.0.disable_flag": "false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, id := range []string{
				"123456",
				"999999",
				"1000000",
				"1234567",
				"9007199254740993",
				"9223372036854775807",
			} {
				t.Run(id, func(t *testing.T) {
					body := `[{"id":` + id + `,"name":"example"` + tt.attributes + `}]`
					constellixClient, closeProxy := dataSourceTestClient(t, tt.path, body)
					defer closeProxy()

					config := map[string]interface{}{"name": "example"}
					if _, ok := tt.resource.Schema["domain_id"]; ok {
						config["domain_id"] = "123456"
						config["source_type"] = "domains"
					}
					d := schema.TestResourceDataRaw(t, tt.resource.Schema, config)
					if err := tt.resource.Read(d, constellixClient); err != nil {
						t.Fatal(err)
					}
					if d.Id() != id {
						t.Errorf("ID = %q, want %q", d.Id(), id)
					}
					for key, want := range tt.want {
						if got := d.Get(key); !reflect.DeepEqual(got, want) {
							t.Errorf("%s = %#v, want %#v", key, got, want)
						}
					}
				})
			}
		})
	}
}

func TestDataSourceInvalidJSON(t *testing.T) {
	for _, body := range []string{`[`, `{}`, `[{"id":1000000,"name":"example"}] trailing`, `[{"id":1000000,"name":"example"}] []`} {
		t.Run(body, func(t *testing.T) {
			constellixClient, closeProxy := dataSourceTestClient(t, "/v2/tags", body)
			defer closeProxy()
			resource := datasourceConstellixTags()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{"name": "example"})
			if err := resource.Read(d, constellixClient); err == nil {
				t.Fatal("expected a JSON decoding error")
			}
			if d.Id() != "" {
				t.Fatalf("invalid response set ID to %q", d.Id())
			}
		})
	}
}

// The API client has a fixed base URL, so serve mocked API responses through its proxy option.
func dataSourceTestClient(t *testing.T, path, body string) (*client.Client, func()) {
	t.Helper()
	var proxy *httptest.Server
	proxy = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodConnect || r.Host != "api.dns.constellix.com:443" {
			t.Errorf("unexpected proxy request: %s %s", r.Method, r.Host)
			http.Error(w, "unexpected proxy request", http.StatusBadRequest)
			return
		}
		hijacker, ok := w.(http.Hijacker)
		if !ok {
			t.Error("proxy connection cannot be hijacked")
			return
		}
		conn, _, err := hijacker.Hijack()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()
		if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
			t.Error(err)
			return
		}
		if _, err := io.WriteString(conn, "HTTP/1.1 200 Connection Established\r\n\r\n"); err != nil {
			t.Error(err)
			return
		}
		apiConn := tls.Server(conn, proxy.TLS)
		defer apiConn.Close()
		request, err := http.ReadRequest(bufio.NewReader(apiConn))
		if err != nil {
			t.Error(err)
			return
		}
		defer request.Body.Close()
		if request.Method != http.MethodGet || request.URL.Path != path {
			t.Errorf("API request = %s %s, want GET %s", request.Method, request.URL.Path, path)
		}
		response := &http.Response{
			StatusCode:    http.StatusOK,
			ProtoMajor:    1,
			ProtoMinor:    1,
			Header:        http.Header{"Content-Type": []string{"application/json"}},
			Body:          ioutil.NopCloser(strings.NewReader(body)),
			ContentLength: int64(len(body)),
			Close:         true,
		}
		if err := response.Write(apiConn); err != nil {
			t.Error(fmt.Errorf("writing mocked response: %s", err))
		}
	}))
	return client.GetClient("test-api-key", "test-secret-key",
		client.Insecure(true), client.ProxyUrl(proxy.URL)), proxy.Close
}
