package vaneck

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
			args: args{body: []byte(`{"data":{"Title":"ETF Statistics","FooterText":"","AsOfDate":"02/28/2025","TooltipText":"","RowsTooltips":["","","","","",""],"Navs":[{"Key":"Total Net Assets","Value":"1,193,754,239"},{"Key":"Bitcoin per 1,000 Shares","Value":".283"},{"Key":"Bitcoin per Basket","Value":"7.071"},{"Key":"Shares Outstanding","Value":"50,225,000"},{"Key":"Bitcoin in Trust","Value":"14,204.695"},{"Key":"Indicative Bitcoin per Basket","Value":"7.071"}],"HighLow":null,"MonthEnd":null,"CssClass":null,"ContainerCssClass":null,"TextClass":null,"ValueClass":null,"AsOfText":"as of","AsAtText":"as at","TimePeriodText":"Time Period","MonthEndText":"Month End","IsNavHistory":false},"blockId":252190}`),
				search: "Bitcoin in Trust"},
			wantResult: types.Result{TotalAsset: 14204.695, Date: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "ETH",
			args: args{body: []byte(`{"data":{"Title":"ETF Statistics","FooterText":"","AsOfDate":"02/28/2025","TooltipText":"","RowsTooltips":["","","","","",""],"Navs":[{"Key":"Total Net Assets","Value":"102,933,830"},{"Key":"Ether per 1,000 Shares","Value":"14.645"},{"Key":"Ether per Basket","Value":"366.132"},{"Key":"Shares Outstanding","Value":"3,175,000"},{"Key":"Ether in Trust","Value":"46,498.750"},{"Key":"Indicative Ether per Basket","Value":"366.132"}],"HighLow":null,"MonthEnd":null,"CssClass":null,"ContainerCssClass":null,"TextClass":null,"ValueClass":null,"AsOfText":"as of","AsAtText":"as at","TimePeriodText":"Time Period","MonthEndText":"Month End","IsNavHistory":false},"blockId":280232}`),
				search: "Ether in Trust"},
			wantResult: types.Result{TotalAsset: 46498.75, Date: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)},
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
