package phone

import "testing"

func TestNormalizeID(t *testing.T) {
	type args struct {
		phoneNumber string
		countryCode int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// ID region
		{
			name: "invalid phone number without plus sign and country code",
			args: args{
				phoneNumber: "6281001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "invalid phone number without plus sign",
			args: args{
				phoneNumber: "6281001099979",
				countryCode: 62,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid phone number with plus sign and no country code",
			args: args{
				phoneNumber: "+6281001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid phone number with plus sign and country code",
			args: args{
				phoneNumber: "+6281001099979",
				countryCode: 62,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid local phone number without plus sign and no country code",
			args: args{
				phoneNumber: "081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid local phone number without plus sign and country code",
			args: args{
				phoneNumber: "081001099979",
				countryCode: 62,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid local phone number with plus sign and no country code",
			args: args{
				phoneNumber: "+081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "valid local phone number with plus sign and country code",
			args: args{
				phoneNumber: "+081001099979",
				countryCode: 62,
			},
			wantRes: "6281001099979",
		},
		// weird numbers
		{
			name: "weird number 6208",
			args: args{
				phoneNumber: "62081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "weird number +6208",
			args: args{
				phoneNumber: "+62081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "weird number 08",
			args: args{
				phoneNumber: "081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "weird number +08",
			args: args{
				phoneNumber: "+081001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "weird number +0628",
			args: args{
				phoneNumber: "+06281001099979",
				countryCode: 0,
			},
			wantRes: "6281001099979",
		},
		{
			name: "weird number +028",
			args: args{
				phoneNumber: "+0281001099979",
				countryCode: 0,
			},
			wantRes: "62281001099979",
		},
		// ID region ends

		// US regions
		{
			name: "valid us number with no plus and no country code",
			args: args{
				phoneNumber: "13153553191",
				countryCode: 0,
			},
			wantRes: "13153553191",
		},
		{
			name: "valid us number with no plus and country code",
			args: args{
				phoneNumber: "13153553191",
				countryCode: 1,
			},
			wantRes: "13153553191",
		},
		{
			name: "valid us number with plus and no country code",
			args: args{
				phoneNumber: "+13153553191",
				countryCode: 1,
			},
			wantRes: "13153553191",
		},
		{
			name: "valid us number with plus and country code",
			args: args{
				phoneNumber: "+13153553191",
				countryCode: 0,
			},
			wantRes: "13153553191",
		},
		// weird numbers
		{
			name: "weird numbers +01",
			args: args{
				phoneNumber: "+013153553191",
				countryCode: 1,
			},
			wantRes: "13153553191",
		},
		// 	US region ends
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := NormalizeID(tt.args.phoneNumber, tt.args.countryCode); gotRes != tt.wantRes {
				t.Errorf("NormalizeID() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
