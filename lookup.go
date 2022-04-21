package publish

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"io"
	_ "log"
	"sync"
	"sync/atomic"
)

func BuildLookup(ctx context.Context, indexer_uri string, indexer_path string) (*sync.Map, error) {

	lookup := new(sync.Map)
	count := int32(0)

	indexer_cb := func(ctx context.Context, path string, fh io.ReadSeeker, args ...interface{}) error {

		body, err := io.ReadAll(fh)

		if err != nil {
			return err
		}

		wof_rsp := gjson.GetBytes(body, "properties.wof:id")

		if !wof_rsp.Exists() {
			return fmt.Errorf("Missing WOF ID")
		}

		wof_id := wof_rsp.Int()

		// See notes about lookup_keys (and media_id) in publish.go

		var media_id string

		phash_rsp := gjson.GetBytes(body, "properties.instagram:post.perceptual_hash")

		if phash_rsp.Exists() {

			id_rsp := gjson.GetBytes(body, "properties.instagram:post.media_id")

			if !id_rsp.Exists() {
				return fmt.Errorf("Record (%d) has perceptual hash but no media ID", wof_id)
			}

			media_id = id_rsp.String()

		} else {

			// I AM HERE
			// Need to fetch and hash image from sfomuseum-media bucket here...
			return fmt.Errorf("Not implemented (yet)")

		}

		lookup.Store(media_id, wof_id)

		atomic.AddInt32(&count, 1)
		return nil
	}

	iter, err := iterator.NewIterator(ctx, indexer_uri, indexer_cb)

	if err != nil {
		return nil, err
	}

	err = iter.IterateURIs(ctx, indexer_path)

	if err != nil {
		return nil, err
	}

	return lookup, nil
}
