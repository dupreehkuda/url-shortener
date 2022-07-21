package storage

import (
	"sync"
	"testing"
)

func Test_storage_Create(t *testing.T) {
	type fields struct {
		mtx       *sync.RWMutex
		shortened map[string]string
	}
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Create duplicated",
			fields: fields{
				mtx: &sync.RWMutex{},
				shortened: map[string]string{
					"rfBd6": "https://youtube.com/",
				},
			},
			args:    args{link: "https://youtube.com/"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				mtx:       tt.fields.mtx,
				shortened: tt.fields.shortened,
			}
			got, err := s.Create(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
