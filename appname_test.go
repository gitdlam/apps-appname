package appname

import "testing"

func TestFolder(t *testing.T) {
	tests := []struct {
		name  string
		wantF string
		wantN string
	}{{"x", "x", "x"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotN := Folder, Name
			t.Log("Please manual check:", gotF, gotN)

		})
	}
}
