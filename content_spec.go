package main

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type Manifest struct {
	Images []ImageSpec `json:"images"`
}

func (m *Manifest) Parse(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(m)
}

type ImageSpec struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`
}

func (i *ImageSpec) GetSizeValue() float64 {
	parts := strings.Split(i.Size, "x")
	f, _ := strconv.ParseFloat(parts[0], 64)
	return f
}

func (i *ImageSpec) GetScaleValue() float64 {
	parts := strings.Split(i.Scale, "x")
	f, _ := strconv.ParseFloat(parts[0], 64)
	return f
}
