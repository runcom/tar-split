package verify

// CheckType is how the on disk attributes will be verified against the
// recorded header information
type CheckType int

// Check types for customizing how fuzzy or strict on-disk verification will be
// handled
const (
	CheckFileSize CheckType = iota
	CheckMode
	CheckDigest
)
