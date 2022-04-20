package main

import (
	_ "gocloud.dev/blob/fileblob"
	_ "image/jpeg"
)

import (
	"bytes"
	"image"
	"crypto/sha256"
	"context"
	"flag"
	"fmt"
	"io"
	"github.com/sfomuseum/go-sfomuseum-instagram/media"
	"github.com/sfomuseum/go-sfomuseum-instagram/walk"
	"log"
	"github.com/tidwall/gjson"
	"gocloud.dev/blob"
	"github.com/corona10/goimagehash"
	"path/filepath"
)

func main() {

	media_bucket_uri := flag.String("media-bucket-uri", "file:///Volumes/Museum/_Public/_Social_Media/SM downloads/2022/sfomuseum_20220418/", "...")
	
	flag.Parse()

	ctx := context.Background()

	media_bucket, err := blob.OpenBucket(ctx, *media_bucket_uri)

	if err != nil {
		log.Fatalf("Failed to open media bucket, %v", err)
	}
	
	cb := func(ctx context.Context, body []byte) error {

		path_rsp := gjson.GetBytes(body, "path")

		if !path_rsp.Exists() {
			return nil
		}

		rel_path := path_rsp.String()

		img_fh, err := media_bucket.NewReader(ctx, rel_path, nil)

		if err != nil {
			return fmt.Errorf("Failed to open %s, %w", rel_path, err)
		}

		defer img_fh.Close()

		img_body, err := io.ReadAll(img_fh)

		if err != nil {
			return fmt.Errorf("Failed to read %s, %w", rel_path, err)
		}

		img_hash := fmt.Sprintf("%x", sha256.Sum256(img_body))
		p_hash := ""
		
		if filepath.Ext(rel_path) != ".mp4" {

			im_r := bytes.NewReader(img_body)
			im, _, err := image.Decode(im_r)

			if err != nil {
				return fmt.Errorf("Failed to decode %s, %w", rel_path, err)
			}

			avg_hash, err := goimagehash.AverageHash(im)

			if err != nil {
				return fmt.Errorf("Failed to derive average hash for %s, %w", rel_path, err)
			}

			p_hash = avg_hash.ToString()
		}
		
		log.Println(rel_path, img_hash, p_hash)
		return nil
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
