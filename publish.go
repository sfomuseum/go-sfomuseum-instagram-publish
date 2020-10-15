package publish

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "fmt"
	sfom_reader "github.com/sfomuseum/go-sfomuseum-reader"
	sfom_writer "github.com/sfomuseum/go-sfomuseum-writer"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-export/exporter"
	"github.com/whosonfirst/go-writer"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type PublishOptions struct {
	Lookup   *sync.Map
	Reader   reader.Reader
	Writer   writer.Writer
	Exporter exporter.Exporter
}

func PublishMedia(ctx context.Context, opts *PublishOptions, body []byte) error {

	select {
	case <-ctx.Done():
		return nil
	default:
		// pass
	}

	path_rsp := gjson.GetBytes(body, "path")

	if !path_rsp.Exists() {
		return errors.New("Missing 'path' property")
	}

	caption_rsp := gjson.GetBytes(body, "caption")

	if !caption_rsp.Exists() {
		return errors.New("Missing 'caption' property")
	}

	path := path_rsp.String()
	// caption := caption_rsp.String()

	fname := filepath.Base(path)
	ext := filepath.Ext(path)

	media_id := strings.Replace(fname, ext, "", 1)

	pointer, ok := opts.Lookup.Load(media_id)

	var wof_record []byte

	if ok {

		wof_id := pointer.(int64)

		wof_body, err := sfom_reader.LoadBytesFromID(ctx, opts.Reader, wof_id)

		if err != nil {
			return err
		}

		wof_record = wof_body

	} else {

		new_record, err := newWOFRecord(ctx)

		if err != nil {
			return err
		}

		wof_record = new_record

		taken_rsp := gjson.GetBytes(body, "taken_at")

		if !taken_rsp.Exists() {
			return errors.New("Missing created timestamp")
		}

		taken_at := taken_rsp.String()

		taken_t, err := time.Parse(time.RFC3339, taken_at)

		if err != nil {
			return err
		}

		wof_record, err = sjson.SetBytes(wof_record, "properties.wof:created", taken_t.Unix())

		if err != nil {
			return err
		}

		wof_record, err = sjson.SetBytes(wof_record, "properties.edtf:inception", taken_at)

		if err != nil {
			return err
		}

		wof_record, err = sjson.SetBytes(wof_record, "properties.edtf:cessation", taken_at)

		if err != nil {
			return err
		}
	}

	wof_name := fmt.Sprintf("Instagram message #%d", media_id)
	wof_record, err := sjson.SetBytes(wof_record, "properties.wof:name", wof_name)

	if err != nil {
		return err
	}

	// why does this keep getting set incorrectly...

	/*
		wof_record, err = sjson.SetBytes(wof_record, "properties.wof:concordances.twitter:id", tweet_id)

		if err != nil {
			return err
		}
	*/

	var post interface{}

	err = json.Unmarshal(body, &post)

	if err != nil {
		return err
	}

	wof_record, err = sjson.SetBytes(wof_record, "properties.instagram:post", post)

	if err != nil {
		return err
	}

	wof_record, err = opts.Exporter.Export(wof_record)

	if err != nil {
		return err
	}

	id, err := sfom_writer.WriteFeatureBytes(ctx, opts.Writer, wof_record)

	if err != nil {
		return err
	}

	log.Printf("Wrote %d\n", id)
	return nil
}

func newWOFRecord(ctx context.Context) ([]byte, error) {

	// Null Terminal - please read these details from source...
	// https://raw.githubusercontent.com/sfomuseum-data/sfomuseum-data-architecture/master/data/115/916/086/9/1159160869.geojson

	parent_id := 1159160869

	lat := 37.616356
	lon := -122.386166

	hier := []map[string]interface{}{
		{
			"building_id":      1159160869,
			"campus_id":        102527513,
			"continent_id":     102191575,
			"country_id":       85633793,
			"county_id":        102087579,
			"locality_id":      85922583,
			"neighbourhood_id": -1,
			"region_id":        85688637,
		},
	}

	geom := map[string]interface{}{
		"type":        "Point",
		"coordinates": [2]float64{lon, lat},
	}

	feature := map[string]interface{}{
		"type": "Feature",
		"properties": map[string]interface{}{
			"sfomuseum:placetype": "instagram",
			"src:geom":            "sfomuseum",
			"wof:country":         "US",
			"wof:parent_id":       parent_id,
			"wof:placetype":       "custom",
			"wof:repo":            "sfomuseum-data-socialmedia-instagram",
			"wof:hierarchy":       hier,
		},
		"geometry": geom,
	}

	return json.Marshal(feature)
}
