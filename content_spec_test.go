package main

import (
	"strings"
	"testing"
)

func TestGetImageSizeFloat(t *testing.T) {
	i := &ImageSpec{
		Size: "30x30",
	}
	expected := float64(30)
	actual := i.GetSizeValue()
	if expected != actual {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}
func TestGetImageScaleFloat(t *testing.T) {
	i := &ImageSpec{
		Scale: "30x",
	}
	expected := float64(30)
	actual := i.GetScaleValue()
	if expected != actual {
		t.Fatalf("Expected %v got %v", expected, actual)
	}
}

func TestParse(t *testing.T) {
	data := `
	{
		"images" : [
		  {
			"size" : "20x20",
			"idiom" : "iphone",
			"filename" : "XOEye_AppIcon-29@2x-1.png",
			"scale" : "2x"
		  },
		  {
			"size" : "20x20",
			"idiom" : "iphone",
			"filename" : "XOEye_AppIcon-29@3x-1.png",
			"scale" : "3x"
		  }
		]
	}
`
	reader := strings.NewReader(data)
	c := &Manifest{}
	err := c.Parse(reader)
	if err != nil {
		t.Fatal(err)
	}
	expected := 2
	got := len(c.Images)
	if expected != got {
		t.Fatalf("Expected %v images but got %v", expected, got)
	}
}
