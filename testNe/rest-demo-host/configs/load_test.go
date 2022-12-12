package configs

import "testing"

func TestLoadConfigFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfigFromEnv(); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigFromEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
