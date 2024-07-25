package cmeny

import (
	"testing"
	"time"
)

func TestReferenceRate_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Value float64
		Date  time.Time
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1",
			fields:  fields{Value: 3476.55, Date: time.Date(2024, 7, 23, 20, 0, 0, 0, time.UTC)},
			args:    args{data: []byte(`{"value": "3476.55", "date": "2024-07-23 20:00:00"}`)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &ReferenceRate{
				Value: tt.fields.Value,
				Date:  tt.fields.Date,
			}
			if err := rr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ReferenceRate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if fields are correctly populated
			if !rr.Date.Equal(tt.fields.Date) || rr.Value != tt.fields.Value {
				t.Errorf("UnmarshalJSON() got = %+v, want %+v", rr, tt.fields)
			}
		})
	}
}
