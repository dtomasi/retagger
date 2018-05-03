package main

import (
	"testing"
)

func TestRetaggedName(t *testing.T) {
	registry := "quay.io"
	organisation := "giantswarm"

	tests := []struct {
		image        Image
		expectedName string
	}{
		{
			image: Image{
				Name: "quay.io/coreos/hyperkube",
			},
			expectedName: "quay.io/giantswarm/hyperkube",
		},
		{
			image: Image{
				Name: "prom/prometheus",
			},
			expectedName: "quay.io/giantswarm/prometheus",
		},
		{
			image: Image{
				Name: "golang",
			},
			expectedName: "quay.io/giantswarm/golang",
		},
	}

	for _, test := range tests {
		returnedName := RetaggedName(registry, organisation, test.image)
		if returnedName != test.expectedName {
			t.Fatalf("'%v' != '%v'", returnedName, test.expectedName)
		}
	}
}

func TestShaName(t *testing.T) {
	tests := []struct {
		imageName    string
		sha          string
		expectedName string
	}{
		{
			imageName:    "quay.io/coreos/hyperkube",
			sha:          "5ff22b5c65d5b93aa948b79028dc136a22cda2f049283103f10bd45650b47312",
			expectedName: "quay.io/coreos/hyperkube@sha256:5ff22b5c65d5b93aa948b79028dc136a22cda2f049283103f10bd45650b47312",
		},
	}

	for _, test := range tests {
		returnedName := ShaName(test.imageName, test.sha)
		if returnedName != test.expectedName {
			t.Fatalf("'%v' != '%v'", returnedName, test.expectedName)
		}
	}
}

func TestParseConfigFile(t *testing.T) {

	ims := `[
		{
			"Name": "coredns/coredns",
			"Tags": [
				{
					"Sha": "d291f8b87eab26845a0c4605df4194924806712c4f624b9a9ddfc9d382b3ddbd",
					"Tag": "1.0.4"
				},
				{
					"Sha": "a01b8b7465f8ce5326e1589c7bbed1b99322804c472872a03edb60fbedaaa6f6",
					"Tag": "1.0.5"
				}
			]
		}
	]`
	images, err := ParseImageListConfig([]byte(ims))
	if err != nil {
		t.Fatalf("config cannot be parsed: %v", err)
	}

	if len(images) != 1 {
		t.Fatalf("1 image expected, got %v", len(images))
	}

	if len(images[0].Tags) != 2 {
		t.Fatalf("2 tags expected, got %v", len(images[0].Tags))
	}
}
