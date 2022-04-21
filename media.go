package publish

import (
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-instagram/media"
	"github.com/tidwall/gjson"
)

// DeriveMediaId will derive a (hopefully) persistent SFO Museum specific media ID
// from JSON properties in 'body'. This might be a `go-sfomuseum-instagram/media.Photo`
// instance or a WOF-style SFO Museum record.
func DeriveMediaId(body []byte, prefix string) (string, error) {

	path_taken := "taken_at"
	path_phash := "perceptual_hash"

	if prefix != "" {
		path_taken = fmt.Sprintf("%s,%s", prefix, path_taken)
		path_phash = fmt.Sprintf("%s,%s", prefix, path_phash)
	}

	taken_rsp := gjson.GetBytes(body, "taken_at")

	if !taken_rsp.Exists() {
		return "", fmt.Errorf("Missing '%s' property", path_taken)
	}

	phash_rsp := gjson.GetBytes(body, "perceptual_hash")

	if !phash_rsp.Exists() {
		return "", fmt.Errorf("Missing '%s' property", path_phash)
	}

	media_id := fmt.Sprintf("%d-%s", taken_rsp.Int(), phash_rsp.String())
	return media.DeriveMediaIdFromString(media_id), nil
}
