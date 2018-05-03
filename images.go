package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Image struct {
	Name string
	Tags []Tag
}

type Tag struct {
	Sha string
	Tag string
}

func ParseImageListConfig(config []byte) ([]Image, error) {
	var c []Image
	err := json.Unmarshal(config, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func RetaggedName(registry, organisation string, image Image) string {
	parts := strings.Split(image.Name, "/")

	return fmt.Sprintf("%v/%v/%v", registry, organisation, parts[len(parts)-1])
}

func ImageWithTag(image, tag string) string {
	return fmt.Sprintf("%v:%v", image, tag)
}

func ShaName(imageName, sha string) string {
	return fmt.Sprintf("%v@sha256:%v", imageName, sha)
}
