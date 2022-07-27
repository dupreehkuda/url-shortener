package memstore

import (
	"sync"
	"testing"
)

func Test_storage_Create(t *testing.T) {
	type fields struct {
		mtx       sync.RWMutex
		shortened map[string]string
	}
	type args struct {
		link string
	}
	test := struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{

		name: "Create duplicated",
		fields: fields{
			mtx: sync.RWMutex{},
			shortened: map[string]string{
				"rfBd6": "https://youtube.com/",
			},
		},
		args:    args{link: "https://youtube.com/"},
		want:    "",
		wantErr: true,
	}
	t.Run(test.name, func(t *testing.T) {
		s := storage{
			mtx:       sync.RWMutex{},
			shortened: test.fields.shortened,
		}
		got, err := s.Create(test.args.link)
		if (err != nil) != test.wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got != test.want {
			t.Errorf("Create() got = %v, want %v", got, test.want)
		}
	})
}
