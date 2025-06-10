package mcp

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFileResourceProvider tests the FileResourceProvider functionality
func TestFileResourceProvider(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "resource-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create some test files
	testFiles := map[string]string{
		"test1.txt":           "This is test file 1",
		"test2.json":          "{\"name\": \"Test 2\", \"value\": 42}",
		"subdir/test3.txt":    "This is test file 3 in a subdirectory",
		"subdir/test4.binary": "\x00\x01\x02\x03\x04",
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(tempDir, path)

		// Create the directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		require.NoError(t, err)

		// Write the file
		err = os.WriteFile(fullPath, []byte(content), 0644)
		require.NoError(t, err)
	}

	// Create a resource provider with the temp directory as base path
	provider := NewFileResourceProvider(tempDir)

	// Test registering individual resources
	t.Run("RegisterResource", func(t *testing.T) {
		provider.RegisterResource("file:///test1.txt", filepath.Join(tempDir, "test1.txt"), "Test file 1")

		// List resources
		resources, err := provider.ListResources()
		require.NoError(t, err)

		// Should have 1 resource
		assert.Len(t, resources, 1)
		assert.Equal(t, "file:///test1.txt", resources[0].URI)
		assert.Equal(t, "test1.txt", resources[0].Name)
	})

	// Test registering a directory of resources
	t.Run("RegisterDirectory", func(t *testing.T) {
		// Clear previous registrations
		provider = NewFileResourceProvider(tempDir)

		// Register the subdirectory
		err := provider.RegisterDirectory("file:///subdir", "subdir")
		require.NoError(t, err)

		// List resources
		resources, err := provider.ListResources()
		require.NoError(t, err)

		// Should have 2 resources (from the subdirectory)
		assert.Len(t, resources, 2)

		// Resources should be sorted by name, but the order isn't guaranteed
		resourceURIs := make([]string, len(resources))
		for i, res := range resources {
			resourceURIs[i] = res.URI
		}

		assert.Contains(t, resourceURIs, "file:///subdir/test3.txt")
		assert.Contains(t, resourceURIs, "file:///subdir/test4.binary")
	})

	// Test reading resources
	t.Run("ReadResource", func(t *testing.T) {
		// Clear previous registrations
		provider = NewFileResourceProvider(tempDir)

		// Register all files
		provider.RegisterResource("file:///test1.txt", filepath.Join(tempDir, "test1.txt"), "Test file 1")
		provider.RegisterResource("file:///test2.json", filepath.Join(tempDir, "test2.json"), "Test file 2")

		// Test reading a text file
		t.Run("TextFile", func(t *testing.T) {
			result, err := provider.ReadResource("file:///test1.txt")
			require.NoError(t, err)

			assert.Len(t, result.Contents, 1)
			assert.Equal(t, "file:///test1.txt", result.Contents[0].URI)
			assert.Contains(t, result.Contents[0].MimeType, "text/plain")
			assert.Equal(t, "This is test file 1", result.Contents[0].Text)
			assert.Empty(t, result.Contents[0].Blob) // Should not have binary data
		})

		// Test reading a JSON file
		t.Run("JSONFile", func(t *testing.T) {
			result, err := provider.ReadResource("file:///test2.json")
			require.NoError(t, err)

			assert.Len(t, result.Contents, 1)
			assert.Equal(t, "file:///test2.json", result.Contents[0].URI)
			assert.Equal(t, "application/json", result.Contents[0].MimeType)
			assert.Equal(t, "{\"name\": \"Test 2\", \"value\": 42}", result.Contents[0].Text)
		})

		// Test reading a non-existent resource
		t.Run("NonExistentResource", func(t *testing.T) {
			_, err := provider.ReadResource("file:///non-existent.txt")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "resource not found")
		})
	})
}

// TestBuildResourceURI tests the BuildResourceURI function
func TestBuildResourceURI(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Simple path",
			path:     "test.txt",
			expected: "file:///test.txt",
		},
		{
			name:     "Path with directory",
			path:     "path/to/test.txt",
			expected: "file:///path/to/test.txt",
		},
		{
			name:     "Absolute path",
			path:     "/path/to/test.txt",
			expected: "file:///path/to/test.txt",
		},
		{
			name:     "Path with backslashes",
			path:     "path\\to\\test.txt",
			expected: "file:///path/to/test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildResourceURI(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}
