package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindDomainByUrl(t *testing.T) {
	t.Run("parser is in list", func(t *testing.T) {
		posturl := "https://www.qq.com"
		jsonfile, err := reloadParser("./testdata/test.json")
		got, err := findDomainByUrl(posturl, jsonfile)
		want := &Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		require.Equal(t, want, got, "the two json should be same")
	})
}
