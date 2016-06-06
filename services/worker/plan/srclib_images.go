package plan

import (
	"fmt"
	"strings"
)

// Drone Docker images for running each toolchain. Remove the sha256 version when
// developing to make it easier to test out changes to a given toolchain. E.g.,
// `droneSrclibGoImage = "sourcegraph/srclib-go"`.
var (
	droneSrclibGoImage         = "sourcegraph/srclib-go@sha256:06a824e0d8663ad0dbd893849432f9f3d2d91f25b2bf3a4dcb02e96d0d26fb04"
	droneSrclibJavaScriptImage = "sourcegraph/srclib-javascript@sha256:d09a1d9cb01f27fefccb13532df09e581e3ed23d924cc50e8d281dd4ad47e275"
	droneSrclibJavaImage       = "sourcegraph/srclib-java@sha256:4f3ffd9dd8b4b0f38c9c2b381dd5dc2103b5052e0d1524f7a0c7883b122e5056"
	droneSrclibTypeScriptImage = "sourcegraph/srclib-typescript@sha256:39adfea4bdaea50be63431fe8c85c174a6a83d34db1196ac0bb171cb79cc88d6"
	droneSrclibCSharpImage     = "sourcegraph/srclib-csharp@sha256:e5c112fc5ccb0551a09289cd732b00d038c4bc366f853e552826e36bcd903507"
	droneSrclibCSSImage        = "sourcegraph/srclib-css@sha256:5dea2ffe7183d2fb3f3f3d6a43790189f1c9a82ef46d8d605d87a99830ca9fbd"
	droneSrclibPythonImage     = "sourcegraph/srclib-python@sha256:ce9625fc4648cc2200237dc95c9b14b0eeb01621a55a53f5123509e84da069be"
)

func versionHash(image string) (string, error) {
	split := strings.Split(image, "@sha256:")
	if len(split) != 2 {
		return "", fmt.Errorf("cannot parse version hash from toolchain image %s", image)
	}

	return split[1], nil
}

func SrclibVersion(lang string) (string, error) {
	switch lang {
	case "Bash":
		// TODO(mate): create a droneSrclibBashImage before enabling in prod
		return "", nil
	case "Go":
		return versionHash(droneSrclibGoImage)
	case "JavaScript":
		return versionHash(droneSrclibJavaScriptImage)
	case "Java":
		return versionHash(droneSrclibJavaImage)
	case "TypeScript":
		return versionHash(droneSrclibTypeScriptImage)
	case "C#":
		return versionHash(droneSrclibCSharpImage)
	case "CSS":
		return versionHash(droneSrclibCSSImage)
	case "Python":
		return versionHash(droneSrclibPythonImage)
	}

	return "", fmt.Errorf("no srclib image found for %s", lang)
}
