package pkg

import "testing"

func TestParsingVersionFromOutput(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"abc", args{`{
  "terraform_version": "0.13.0",
  "terraform_revision": "",
  "provider_selections": {},
  "terraform_outdated": false
}`}, "0.13.0", false},
		{"not", args{`Terraform v0.12.29

Your version of Terraform is out of date! The latest version
is 0.13.0. You can update by downloading from https://www.terraform.io/downloads.html
        `}, "0.12.29", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsingVersionFromOutput(tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsingVersionFromOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParsingVersionFromOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
