package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldSetAndGetResourcesToAndFromAppContext(t *testing.T) {
	// Given
	resourceName := "test-resource-name"
	resource := "some-test-string-resource"
	appContext := NewApplicationContext()

	// When
	appContext.Set(resourceName, resource)
	actualResource := appContext.Get(resourceName)

	// Then
	assert.Equal(t, resource, actualResource, "should be the same resource")
}
