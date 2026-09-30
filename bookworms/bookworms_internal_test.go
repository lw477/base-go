package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		filePath string // 路径
		want     []Bookworm
		wantErr  bool
	}

	tests := map[string]testCase{
		"file exists": {
			filePath: "testdata/bookworms.json",
			want: []Bookworm{
				{
					Name:  "Fadi",
					Books: []Book{handmaidsTale, theBellJar},
				},
				{
					Name:  "Peggy",
					Books: []Book{oryxAndCrake, handmaidsTale, janeEyre},
				},
			},
			wantErr: false,
		},
		"file not exists": {
			filePath: "testdata/not_exist.json",
			want:     nil,
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := LoadBookworms(tc.filePath)

			if (err != nil) != tc.wantErr {
				t.Fatalf("error mismatch:got err = %v, wantErr=%v", err, tc.wantErr)
			}
			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("result not match:\ngot: %+v\nwant: %+v", got, tc.want)
			}

		})
	}
}

func equalBookworms(t *testing.T, a, b []Bookworm) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name {
			return false
		}
		if !equalBooks(t, a[i].Books, b[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(t *testing.T, a, b []Book) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
