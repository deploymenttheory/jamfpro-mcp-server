package mcp

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ResourceProvider defines an interface for providing MCP resources
type ResourceProvider interface {
	// ListResources returns a list of available resources
	ListResources() ([]Resource, error)

	// ReadResource reads the content of a resource identified by URI
	ReadResource(uri string) (*ReadResourceResult, error)
}

// FileResourceProvider provides access to file system resources
type FileResourceProvider struct {
	basePath             string
	resourceURIs         map[string]string // Maps URI to file path
	resourceDescriptions map[string]string // Maps URI to description
}

// NewFileResourceProvider creates a new file resource provider
func NewFileResourceProvider(basePath string) *FileResourceProvider {
	return &FileResourceProvider{
		basePath:             basePath,
		resourceURIs:         make(map[string]string),
		resourceDescriptions: make(map[string]string),
	}
}

// RegisterResource registers a file as a resource
func (p *FileResourceProvider) RegisterResource(uri, filePath string, description string) {
	p.resourceURIs[uri] = filePath
	// Store the description with the URI for later use
	if description != "" {
		p.resourceDescriptions[uri] = description
	}
}

// RegisterDirectory registers all files in a directory as resources
func (p *FileResourceProvider) RegisterDirectory(uriPrefix, dirPath string) error {
	fullPath := filepath.Join(p.basePath, dirPath)

	// Check if the directory exists
	info, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory does not exist: %s", fullPath)
		}
		return err
	}

	// Ensure it's a directory
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", fullPath)
	}

	return filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(fullPath, path)
			if err != nil {
				return err
			}

			// Convert backslashes to forward slashes for URI consistency
			relPath = strings.ReplaceAll(relPath, "\\", "/")

			uri := fmt.Sprintf("%s/%s", strings.TrimSuffix(uriPrefix, "/"), relPath)
			p.resourceURIs[uri] = path

			// Set a default description
			p.resourceDescriptions[uri] = fmt.Sprintf("File: %s", filepath.Base(path))
		}

		return nil
	})
}

// ListResources returns a list of registered resources
func (p *FileResourceProvider) ListResources() ([]Resource, error) {
	resources := make([]Resource, 0, len(p.resourceURIs))

	for uri, path := range p.resourceURIs {
		info, err := os.Stat(path)
		if err != nil {
			continue // Skip files that can't be accessed
		}

		mimeType := mime.TypeByExtension(filepath.Ext(path))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		description := p.resourceDescriptions[uri]
		if description == "" {
			description = fmt.Sprintf("File: %s", filepath.Base(path))
		}

		resources = append(resources, Resource{
			URI:         uri,
			Name:        filepath.Base(path),
			Description: description,
			MimeType:    mimeType,
			Meta: map[string]interface{}{
				"size": info.Size(),
				"path": path,
			},
		})
	}

	return resources, nil
}

// ReadResource reads the content of a resource
func (p *FileResourceProvider) ReadResource(uri string) (*ReadResourceResult, error) {
	path, ok := p.resourceURIs[uri]
	if !ok {
		return nil, fmt.Errorf("resource not found: %s", uri)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	mimeType := mime.TypeByExtension(filepath.Ext(path))
	if mimeType == "" {
		// Try to detect content type from first 512 bytes
		// Create a buffer to read the first 512 bytes
		var buf [512]byte
		n, _ := file.Read(buf[:])
		file.Seek(0, 0) // Reset file position to start

		// Detect content type
		contentType := http.DetectContentType(buf[:n])
		mimeType = contentType
	}

	// For text files, read as text
	if isTextMimeType(mimeType) {
		content, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		return &ReadResourceResult{
			Contents: []ResourceContent{
				{
					URI:      uri,
					MimeType: mimeType,
					Text:     string(content),
					Meta: map[string]interface{}{
						"size": info.Size(),
						"path": path,
					},
				},
			},
		}, nil
	}

	// For binary files, encode as base64
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return &ReadResourceResult{
		Contents: []ResourceContent{
			{
				URI:      uri,
				MimeType: mimeType,
				Blob:     base64.StdEncoding.EncodeToString(content),
				Meta: map[string]interface{}{
					"size": info.Size(),
					"path": path,
				},
			},
		},
	}, nil
}

// Common resource URIs
const (
	ResourceURIPrefix = "file://"
)

// BuildResourceURI builds a standard resource URI for a file
func BuildResourceURI(path string) string {
	// Normalize path by replacing backslashes with forward slashes
	normalizedPath := strings.ReplaceAll(path, "\\", "/")

	// Ensure the path starts with a slash
	if !strings.HasPrefix(normalizedPath, "/") {
		normalizedPath = "/" + normalizedPath
	}

	return ResourceURIPrefix + normalizedPath
}

// isTextMimeType checks if a MIME type represents a text file
func isTextMimeType(mimeType string) bool {
	// Common text-based MIME types
	return strings.HasPrefix(mimeType, "text/") ||
		strings.Contains(mimeType, "json") ||
		strings.Contains(mimeType, "xml") ||
		strings.Contains(mimeType, "javascript") ||
		strings.Contains(mimeType, "typescript") ||
		strings.Contains(mimeType, "yaml") ||
		strings.Contains(mimeType, "markdown") ||
		strings.Contains(mimeType, "css") ||
		strings.Contains(mimeType, "csv") ||
		strings.Contains(mimeType, "+xml") || // For types like application/atom+xml
		strings.Contains(mimeType, "+json") || // For types like application/ld+json
		mimeType == "application/x-sh" || // Shell scripts
		mimeType == "application/x-shellscript" // Shell scripts
}
