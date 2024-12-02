# go-sfomuseum-instagram-publish

## Tools

### publish

This is not ideal so we might need to revisit this code to search a folder recursively for `media.json` files.

```
$> go run -mod vendor cmd/publish/main.go \
	-reader-uri fs:///usr/local/data/sfomuseum-data-socialmedia-instagram/data \
	-writer-uri fs:///usr/local/data/sfomuseum-data-socialmedia-instagram/data \
	file:///Users/asc/Desktop/instagram/sfomuseum_20201008_part_2/media.json \
	file:///Users/asc/Desktop/instagram/sfomuseum_20201008_part_3/media.json \
	file:///Users/asc/Desktop/instagram/sfomuseum_20201008_part_4/media.json \
	file:///Users/asc/Desktop/instagram/sfomuseum_20201008_part_5/media.json
```

## See also

* https://github.com/sfomuseum/go-sfomuseum-instagram