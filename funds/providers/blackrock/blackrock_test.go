package blackrock

import (
	"reflect"
	"testing"

	"github.com/jyap808/cryptoEtfScrape/types"
)

func Test_parseJSON(t *testing.T) {
	type args struct {
		bodyStr string
		ticker  string
	}
	tests := []struct {
		name       string
		args       args
		wantResult types.Result
		wantErr    bool
	}{
		{
			name: "ETHA",
			args: args{bodyStr: `{"aaData":[["ETH","ETHER","Alternative",{"display":"$2,861,350,156.52","raw":2861350156.52},{"display":"100.00","raw":99.9998},{"display":"2,861,350,156.52","raw":2861350156.52},{"display":"1,292,488.65380","raw":1292488.6538}],["USD","USD CASH","Cash",{"display":"$6,017.57","raw":6017.57},{"display":"0.00","raw":0.0002},{"display":"6,017.57","raw":6017.57},{"display":"6,017.57000","raw":6017.57}]]}`,
				ticker: "ETH"},
			wantResult: types.Result{TotalAsset: 1292488.6538},
			wantErr:    false,
		},
		{
			name: "IBIT",
			args: args{bodyStr: `{"aaData":[["BTC","BITCOIN","Alternative",{"display":"$48,166,141,929.58","raw":48166141929.58},{"display":"100.00","raw":99.9999},{"display":"48,166,141,929.58","raw":48166141929.58},{"display":"573,135.98590","raw":573135.9859}],["USD","USD CASH","Cash",{"display":"$45,600.64","raw":45600.64},{"display":"0.00","raw":0.0001},{"display":"45,600.64","raw":45600.64},{"display":"45,600.64000","raw":45600.64}]]}`,
				ticker: "BTC"},
			wantResult: types.Result{TotalAsset: 573135.9859},
			wantErr:    false,
		},
		{
			name: "ETHA - Bounds error",
			args: args{bodyStr: `{"aaData":[["ETH","ETHER","Alternative",{"display":"$2,861,350,156.52","raw":2861350156.52},{"display":"100.00","raw":99.9998},{"display":"2,861,350,156.52","raw":2861350156.52},{"display":"1,292,488.65380","raw":2957069172.8}],["USD","USD CASH","Cash",{"display":"$6,017.57","raw":6017.57},{"display":"0.00","raw":0.0002},{"display":"6,017.57","raw":6017.57},{"display":"6,017.57000","raw":6017.57}]]}`,
				ticker: "ETH"},
			wantResult: types.Result{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := parseJSON(tt.args.bodyStr, tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("parseJSON() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
