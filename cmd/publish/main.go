// publish is a command-line tool to merge Instagram posts defined in one or more "media.json" files
// with the sfomuseum-data-socialmedia-instagram repository. For example:
//
//	$> ./bin/publish file:///usr/local/data/media.json
//
// Important: As of April, 2022 Instagram no longer publishes "media.json" files with the export bundles.
// Use the sfomuseum/go-sfomuseum-instagram/cmd/derive-media-json tool to create a media.json file from
// the available data.
package main

import (
	"context"
	"flag"
	_ "github.com/sfomuseum/go-sfomuseum-export/v2"
	"github.com/sfomuseum/go-sfomuseum-instagram-publish"
	"github.com/sfomuseum/go-sfomuseum-instagram/media"
	"github.com/sfomuseum/go-sfomuseum-instagram/walk"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	"github.com/whosonfirst/go-writer"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	"log"
)

func main() {

	iterator_uri := flag.String("iterator-uri", "repo://", "A valid whosonfirst/go-whosonfirst-iterate/v2 URI")
	iterator_source := flag.String("iterator-source", "/usr/local/data/sfomuseum-data-socialmedia-instagram", "...")

	reader_uri := flag.String("reader-uri", "fs:///usr/local/data/sfomuseum-data-socialmedia-instagram/data", "A valid whosonfirst/go-reader URI")
	writer_uri := flag.String("writer-uri", "fs:///usr/local/data/sfomuseum-data-socialmedia-instagram/data", "A valid whosonfirst/go-writer URI")

	media_bucket_uri := flag.String("media-bucket-uri", "", "A valid gocloud.dev/blob URI")

	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	rdr, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatalf("Failed to create reader, %v", err)
	}

	wrtr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		log.Fatalf("Failed to create writer, %v", err)
	}

	exprtr, err := export.NewExporter(ctx, "sfomuseum://")

	if err != nil {
		log.Fatalf("Failed to create new exported, %w", err)
	}

	lookup, err := publish.BuildLookup(ctx, *iterator_uri, *iterator_source)

	if err != nil {
		log.Fatalf("Failed to build lookup, %v", err)
	}

	media_bucket, err := blob.OpenBucket(ctx, *media_bucket_uri)

	if err != nil {
		log.Fatalf("Failed to open media bucket, %w", err)
	}

	publish_opts := &publish.PublishOptions{
		Lookup:      lookup,
		Reader:      rdr,
		Writer:      wrtr,
		Exporter:    exprtr,
		MediaBucket: media_bucket,
	}

	max_procs := 10
	throttle := make(chan bool, max_procs)

	for i := 0; i < max_procs; i++ {
		throttle <- true
	}

	cb := func(ctx context.Context, body []byte) error {

		<-throttle

		defer func() {
			throttle <- true
		}()

		return publish.PublishMedia(ctx, publish_opts, body)
	}

	walk_opts := &walk.WalkWithCallbackOptions{
		Callback: cb,
	}

	args := flag.Args()

	for _, media_uri := range args {

		media_fh, err := media.Open(ctx, media_uri)

		if err != nil {
			log.Fatalf("Failed to open %s, %v", media_uri, err)
		}

		defer media_fh.Close()

		err = walk.WalkMediaWithCallback(ctx, walk_opts, media_fh)

		if err != nil {
			log.Fatalf("Failed to walk media for %s, %v", media_uri, err)
		}

		log.Println(media_uri)
	}

}
