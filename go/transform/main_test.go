package main

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/text/transform"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type want struct {
		dst  []byte
		nDst int
		nSrc int
		err  error
	}

	cases := []struct {
		name  string
		dst   []byte
		src   []byte
		atEOF bool
		want  want
	}{
		{"aaaa -> AAAA", make([]byte, 4), []byte("aaaa"), false, want{[]byte("AAAA"), 4, 4, nil}},
		{"dst is nil", nil, []byte("aaaa"), false, want{nil, 0, 0, transform.ErrShortDst}},
		{"dst is short", make([]byte, 3), []byte("aaaa"), false, want{[]byte("AAA"), 3, 3, nil}},
		{"src is short", make([]byte, 4), []byte("aaa"), false, want{[]byte("AAA"), 3, 3, nil}},
		{"src is nil", make([]byte, 4), nil, false, want{[]byte{}, 0, 0, transform.ErrShortSrc}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var upper Upper
			nDst, nSrc, err := upper.Transform(tt.dst, tt.src, tt.atEOF)
			if tt.want.err == nil && err != nil {
				t.Fatal("unexpected error:", err)
			}

			if tt.want.err != nil && !errors.Is(err, tt.want.err) {
				t.Fatalf("The expected error %v did not occur", tt.want.err)
			}

			if diff := cmp.Diff(tt.want.dst, tt.dst[:nDst]); diff != "" {
				t.Error(diff)
			}

			if nDst != tt.want.nDst {
				t.Error("nDst", nDst, tt.want.nDst)
			}

			if nSrc != tt.want.nSrc {
				t.Error("nSrc", nSrc, tt.want.nSrc)
			}
		})
	}
}

func TestUpper_Transform(t *testing.T) {
	type args struct {
		dst   []byte
		src   []byte
		atEOF bool
	}
	tests := []struct {
		name     string
		args     args
		wantNDst int
		wantNSrc int
		want     []byte
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				dst:   []byte("Hello, World"),
				src:   []byte("Hello, World"),
				atEOF: false,
			},
			wantNDst: 0,
			wantNSrc: 0,
			want:     []byte("HELLO, WORLD"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Upper{}
			gotNDst, gotNSrc, err := u.Transform(tt.args.dst, tt.args.src, tt.args.atEOF)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNDst != tt.wantNDst {
				t.Errorf("Transform() gotNDst = %v, want %v", gotNDst, tt.wantNDst)
			}
			if gotNSrc != tt.wantNSrc {
				t.Errorf("Transform() gotNSrc = %v, want %v", gotNSrc, tt.wantNSrc)
			}

			if diff := cmp.Diff(tt.want, tt.args.dst); diff != "" {
				t.Errorf("Transform() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
