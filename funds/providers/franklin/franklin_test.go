package franklin

import (
	"reflect"
	"testing"
	"time"

	"github.com/jyap808/cryptoEtfScrape/types"
)

func Test_parseJSON(t *testing.T) {
	type args struct {
		body   []byte
		search string
	}
	tests := []struct {
		name       string
		args       args
		wantResult types.Result
		wantErr    bool
	}{
		{
			name: "BTC",
			args: args{body: []byte(`{"data":{"Portfolio":{"portfolio":{"dailyholdings":[{"asofdate":"02/28/2025","secname":"BITCOIN","quantityshrpar":"6,317.79"},{"asofdate":"02/28/2025","secname":"Net Current Assets","quantityshrpar":"-235,740.21"}]}}}}`),
				search: "BITCOIN"},
			wantResult: types.Result{TotalAsset: 6317.79, Date: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)},
			wantErr:    false,
		},
		{
			name: "ETH",
			args: args{body: []byte(`{"data":{"Portfolio":{"portfolio":{"dailyholdings":[{"asofdate":"02/28/2025","secname":"Net Current Assets","quantityshrpar":"-5,400.06"},{"asofdate":"02/28/2025","secname":"ETHEREUM","quantityshrpar":"12,540.00"}]}}}}`),
				search: "ETHEREUM"},
			wantResult: types.Result{TotalAsset: 12540, Date: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := parseJSON(tt.args.body, tt.args.search)
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
