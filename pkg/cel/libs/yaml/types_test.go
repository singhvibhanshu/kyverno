package yaml

import (
	"reflect"
	"testing"
)

func TestYamlType(t *testing.T) {
	if name := YamlType.TypeName(); name != "yaml.Yaml" {
		t.Errorf("YamlType.TypeName() = %q, want %q", name, "yaml.Yaml")
	}
}

func TestYamlImpl_Parse(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		want    any
		wantErr bool
	}{
		{
			name:    "valid yaml",
			content: []byte("key: value"),
			want:    map[string]interface{}{"key": "value"},
			wantErr: false,
		},
		{
			name:    "valid list",
			content: []byte("- item1\n- item2"),
			want:    []interface{}{"item1", "item2"},
			wantErr: false,
		},
		{
			name:    "invalid yaml",
			content: []byte(": - :"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty yaml",
			content: []byte(""),
			want:    nil,
			wantErr: false,
		},
	}

	impl := &YamlImpl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := impl.Parse(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("YamlImpl.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlImpl.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlStruct(t *testing.T) {
	impl := &YamlImpl{}
	y := Yaml{YamlIface: impl}

	if y.YamlIface == nil {
		t.Error("Yaml struct failed to embed YamlIface")
	}
}
