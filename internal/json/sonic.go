// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import "github.com/bytedance/sonic"

var (
	json = sonic.Config{
		SortMapKeys:      true,  // Critical for gzip efficiency
		CompactMarshaler: true,  // Avoids extra whitespace
		EscapeHTML:       false, // Optional
	}
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)
