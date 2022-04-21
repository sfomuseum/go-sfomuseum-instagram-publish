package publish

import (
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-instagram/media"
	"github.com/tidwall/gjson"
	_ "log"
)

// DeriveMediaId will derive a (hopefully) persistent SFO Museum specific media ID
// from JSON properties in 'body'. This might be a `go-sfomuseum-instagram/media.Photo`
// instance or a WOF-style SFO Museum record.
func DeriveMediaId(body []byte, prefix string) (string, error) {

	path_taken := "taken"
	path_phash := "perceptual_hash"

	if prefix != "" {
		path_taken = fmt.Sprintf("%s.%s", prefix, path_taken)
		path_phash = fmt.Sprintf("%s.%s", prefix, path_phash)
	}

	taken_rsp := gjson.GetBytes(body, path_taken)

	if !taken_rsp.Exists() {
		return "", fmt.Errorf("Missing '%s' property", path_taken)
	}

	phash_rsp := gjson.GetBytes(body, path_phash)

	if !phash_rsp.Exists() {
		return "", fmt.Errorf("Missing '%s' property", path_phash)
	}

	// log.Printf("%s %d\n", path_taken, taken_rsp.Int())
	// log.Printf("%s %s\n", path_phash, phash_rsp.String())

	media_id := fmt.Sprintf("%d-%s", taken_rsp.Int(), phash_rsp.String())
	return media.DeriveMediaIdFromString(media_id), nil
}
