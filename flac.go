// Copyright 2015, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"errors"
	"io"
	"os"
	"time"
	"strconv"
)

// blockType is a type which represents an enumeration of valid FLAC blocks
type blockType byte

// FLAC block types.
const (
	streamInfoBlock    blockType = 0
	paddingBlock                 = 1
	applicationBlock             = 2
	seektableBlock               = 3
	vorbisCommentBlock           = 4 // Supported
	cueSheetBlock                = 5
	pictureBlock                 = 6 // Supported
)

// ReadFLACTags reads FLAC metadata from the io.ReadSeeker, returning the resulting
// metadata in a Metadata implementation, or non-nil error if there was a problem.
func ReadFLACTags(r io.ReadSeeker) (Metadata, error) {
	flac, err := readString(r, 4)
	if err != nil {
		return nil, err
	}
	if flac != "fLaC" {
		return nil, errors.New("expected 'fLaC'")
	}

	m := &metadataFLAC{
		newMetadataVorbis(),
	}

	for {
		last, err := m.readFLACMetadataBlock(r)
		if err != nil {
			return nil, err
		}

		if last {
			break
		}
	}
	return m, nil
}

type metadataFLAC struct {
	*metadataVorbis
}

func (m *metadataFLAC) readFLACMetadataBlock(r io.ReadSeeker) (last bool, err error) {
	blockHeader, err := readBytes(r, 1)
	if err != nil {
		return
	}

	if getBit(blockHeader[0], 7) {
		blockHeader[0] ^= (1 << 7)
		last = true
	}

	blockLen, err := readInt(r, 3)
	if err != nil {
		return
	}

	switch blockType(blockHeader[0]) {

	case streamInfoBlock:
		err = m.readStreamInfo(r)

	case vorbisCommentBlock:
		err = m.readVorbisComment(r)

	case pictureBlock:
		err = m.readPictureBlock(r)

	default:
		_, err = r.Seek(int64(blockLen), io.SeekCurrent)
	}
	return
}

func (m *metadataFLAC) readStramInfo(r io.Reader) error {
	streamInfoLen, err := readInt32LittleEndian(r)
	if err != nil {
		return err
	}
	for i := 0; i < streamInfoLen; i++ {
	metaHeader :=

	}

	m.c["duration"] = strconv.Itoa(duration)
	m.c["channels"] = strconv.Itoa(numChannels)
	m.c["bitrate"] = strconv.Itoa(bitrate)
	return nil
}

func (m *metadataFLAC) FileType() FileType {
	return FLAC
}

func (m *metadataFLAC) Duration() time.Duration {

}

func (m *metadataFLAC) Encoding() FileType {

}

func (m *metadataFLAC) EncodedBy() string {

}

func (m *metadataFLAC) EncoderSettings() string {

}

func (m *metadataFLAC) EncodingBitRate() int {

}

func (m *metadataFLAC) duration() int {

}

func (m *metadataFLAC) bitrate() int {

}
