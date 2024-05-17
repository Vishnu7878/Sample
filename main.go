// Package debug provides functionality for debugging Kubernetes services by fetching logs
// from a Loki API endpoint and uploading them to an AWS S3 bucket.
//
// This package contains the FetchAndUploadLogs function, which fetches logs within a given
// date range from Kubernetes services, constructs queries, fetches logs from Loki, and uploads
// them to an S3 bucket. It also includes the fetchLogs function, which fetches logs from a Loki
// API endpoint in batches until all logs within the specified date range are obtained.
//
// Usage:
//
//	To use this package, import it in your Go code and utilize the FetchAndUploadLogs function
//	to fetch and upload logs from Kubernetes services.
//
// Note: This package requires the "a8-cli/helper/aws" and "a8-cli/helper/logger" packages.
// Ensure they are imported in your codebase for proper functionality.