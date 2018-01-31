package tag

import (
	"encoding/hex"
	"unicode/utf16"
	"time"
	"gopkg.in/mgo.v2/bson"
)


var (
	AsfContextDescriptionObject = []byte{0x33, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11, 0xA6, 0xD9, 0x00, 0xAA, 0x00, 0x62, 0xCE, 0x6C}
	AsfExtendedContentDescriptionObject = []byte{0x40, 0xA4, 0xD0, 0xD2, 0x07, 0xE3, 0xD2, 0x11, 0x97, 0xF0, 0x00, 0xA0, 0xC9, 0x5E, 0xA8, 0x50}
	StreamBitratePropertiesObject = []byte{0x91, 0x07, 0xDC, 0xB7, 0xB7, 0xA9, 0xCF, 0x11, 0x8E, 0xE6, 0x00, 0x00, 0xC0, 0x0C, 0x20, 0x53, 0x65}
	AsfFilePropertyObject = []byte{0xA1, 0xDC, 0xAB, 0x8C, 0x47, 0xA9, 0xCF, 0x11, 0x8E, 0xE4, 0x00, 0xC0, 0x0C, 0x20, 0x53, 0x65}
	AsfStreamPropertiesObject = []byte{0x91, 0x07, 0xDC, 0xB7, 0xB7, 0xA9, 0xCF, 0x11, 0x8E, 0xE6, 0x00, 0xC0, 0x0C, 0x20, 0x53, 0x65}
	StreamTypeAsfAudioMedia = []byte{0x40, 0x9E, 0x69, 0xF8, 0x4D, 0x5B, 0xCF, 0xaa, 0xA8, 0xFD, 0x00, 0x80, 0x5F, 0x5C, 0x44, 0x2B}

)

type metadataASF struct {
	fileType FileType
	data map[string]interface{}
}

var asfCodecTypes = map[string]string{
	"0x0161": "Windows Media Audio 7,8,9",
	"0x0162": "Windows Media Audio 9 Professional",
	"0x0163": "Windows Media Audio 9 Lossless",
	"0x7a21": "GSM-AMR(No SID)",
	"0x7a22": "GSM-AMR(Exist SID)",
}

func ReadASFTags() {}

func (m metadataASF) getString(s []uint16) string {
	return string(utf16.Decode(s))
}

func (m metadataASF) Format() Format {

}

func (m metadataASF) FileType() FileType {

}

func (m metadataASF) Title() string {

}

func (m metadataASF) Artist() string {

}

func (m metadataASF) Album() string {

}

func (m metadataASF) AlbumArtist() string {

}

func (m metadataASF) Composer() string {

}

func (m metadataASF) Performer() string {

}

func (m metadataASF) Genre() string {

}

func (m metadataASF) Year() int {

}

func (m metadataASF) Track() (int, int) {

}

func (m metadataASF) Disc() (int, int) {

}

func (m metadataASF) Picture() *Picture {

}

// Lyrics returns the lyrics, or an empty string if unavailable.
func (m metadataASF) Lyrics() string {

}

func (m metadataASF) Duration() time.Duration {

}

func (m metadataASF) Comment() string {

}

func (m metadataASF) EncodedBy() string {

}

func (m metadataASF) EncoderSettings() string {

}

func (m	metadataASF) Encoding() FileType {

}

func (m metadataASF) EncodingBitRate() int {

}

// Raw returns the raw mapping of retrieved tag names and associated values.
// NB: tag/atom names are not standardised between formats.
func (m metadataASF) Raw() map[string]interface{} {

}
