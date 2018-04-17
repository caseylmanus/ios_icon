# ios_icon
This is a work in progress, designed to replace an iOS icon set from a source image.

## install
go get github.com/caseylmanus/ios_icon

## run
ios_icon --xcassets={path to image assets folder} --source={path to source icon (png only supported)}

## what it does
It will scan the xcassets folder for the Contents.json file, parse that and replace all existing icons with the source image, resized according to the specification.  

## what doesn't work
If your source png has alpha or transperancy it doesn't remove it for the app store icon (yet)

