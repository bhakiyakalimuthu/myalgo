package stringsalgo

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single word",
			args: args{
				words: "nospace",
			},
			want: 1,
		},
		{
			name: "multiple word",
			args: args{
				words: "multiple word without space in the end",
			},
			want: 7,
		},
		{
			name: "multiple word with space at the end",
			args: args{
				words: "multiple word without space in the end ",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCounter1(tt.args.words); got != tt.want {
				t.Errorf("WordCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestWordOccCounter(t *testing.T) {
// 	type args struct {
// 		words string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[string]int
// 	}{
// 		{
// 			name: "",
// 			args: args{
// 				words: "multiple word without space in the end ",
// 			},
// 			want: nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := WordOccCounter(tt.args.words); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("WordOccCounter() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
