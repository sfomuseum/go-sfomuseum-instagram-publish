// one-off backfill script but leaving around for historical purposes.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-instagram/media"
	sfom_writer "github.com/sfomuseum/go-sfomuseum-writer"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"github.com/whosonfirst/go-writer"
	"io"
	"log"
	_ "log"
	"time"
)

func fix_date(str_t string) time.Time {

	loc, _ := time.LoadLocation("America/Los_Angeles")

	t0, _ := time.Parse(time.RFC3339, str_t)

	t0, _ = time.Parse(media.TIME_FORMAT, t0.Format(media.TIME_FORMAT))

	t0 = t0.In(loc)

	if t0.IsDST() {
		// Go from -0700 to -0800
		loc = time.FixedZone("America/Los_Angeles", -28800)
		t0 = t0.In(loc)
	}

	t3, _ := time.Parse(fmt.Sprintf("%s UTC", media.TIME_FORMAT), fmt.Sprintf("%s UTC", t0.Format(media.TIME_FORMAT)))
	return t3
}

func main() {

	iterator_uri := flag.String("iterator-uri", "repo://", "...")
	iterator_source := flag.String("iterator-source", "/usr/local/data/sfomuseum-data-socialmedia-instagram", "...")

	writer_uri := flag.String("writer-uri", "repo:///usr/local/data/sfomuseum-data-socialmedia-instagram", "...")

	flag.Parse()

	ctx := context.Background()

	wr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		log.Fatalf("Failed to create new writer, %v", err)
	}

	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {

		body, err := io.ReadAll(r)

		if err != nil {
			return fmt.Errorf("Failed to read %s, %w", path, err)
		}

		t_rsp := gjson.GetBytes(body, "properties.instagram:post.taken_at")
		new_t := fix_date(t_rsp.String())

		log.Println(t_rsp.String(), new_t.Format(media.TIME_FORMAT))

		body, err = sjson.SetBytes(body, "properties.instagram:post.taken_at", new_t.Format(media.TIME_FORMAT))

		if err != nil {
			return fmt.Errorf("Failed to assign taken_at to %s, %w", path, err)
		}

		body, err = sjson.SetBytes(body, "properties.instagram:post.taken", new_t.Unix())

		if err != nil {
			return fmt.Errorf("Failed to assign taken to %s, %w", path, err)
		}

		_, err = sfom_writer.WriteFeatureBytes(ctx, wr, body)

		if err != nil {
			return fmt.Errorf("Failed to write %s, %w", path, err)
		}

		return nil
	}

	iter, err := iterator.NewIterator(ctx, *iterator_uri, iter_cb)

	if err != nil {
		log.Fatalf("Failed to create new iterator, %v", err)
	}

	err = iter.IterateURIs(ctx, *iterator_source)

	if err != nil {
		log.Fatalf("Failed to iterate, %v", err)
	}

}
