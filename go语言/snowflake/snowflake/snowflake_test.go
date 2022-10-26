package snowflake

import (
	"testing"
)

func TestSnowflake_Gen(t *testing.T) {
	type fields struct {
		machineID int64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "one",
			fields: fields{
				machineID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New(tt.fields.machineID)
			if err != nil {
				t.Errorf("New error %v\n", err)
			}
			output := s.Gen()
			t.Logf("output = %d\n", output)
		})
	}
}
