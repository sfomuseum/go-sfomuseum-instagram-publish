package publish

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"

	"github.com/tidwall/gjson"
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
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

			m, err := DeriveMediaId(body, "properties.instagram:post")

			if err != nil {
				return fmt.Errorf("Failed to derive media ID for %s, %w", path, err)
			}

			media_id = m

		} else {

			log.Printf("%s is missing hash\n", path)
			return nil
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
