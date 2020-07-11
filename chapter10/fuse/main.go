package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/billziss-gh/cgofuse/fuse"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

type CloudFileSystem struct {
	fuse.FileSystemBase
	bucket *blob.Bucket
}

func (cf *CloudFileSystem) Getattr(path string, stat *fuse.Stat_t, fh uint64) (errc int) {
	if path == "/" {
		stat.Mode = fuse.S_IFDIR | 0555
		return 0
	}
	ctx := context.Background()
	name := strings.TrimLeft(path, "/")
	a, err := cf.bucket.Attributes(ctx, name)
	if err != nil {
		_, err := cf.bucket.Attributes(ctx, name+"/")
		if err != nil {
			return -fuse.ENOENT
		}
		stat.Mode = fuse.S_IFDIR | 0555
	} else {
		stat.Mode = fuse.S_IFREG | 0444
		stat.Size = a.Size
		stat.Mtim = fuse.NewTimespec(a.ModTime)
	}
	stat.Nlink = 1
	return 0
}

func (cf *CloudFileSystem) Readdir(path string, fill func(name string, stat *fuse.Stat_t, ofst int64) bool,
ofst int64, fh uint64) (errc int)  {
	ctx := context.Background()
	fill(".", nil, 0)
	fill("..", nil, 0)
	fmt.Println(path)
	prefix := strings.TrimLeft(path, "/")
	if prefix != "" {
		prefix = prefix + "/"
	}
	i := cf.bucket.List(&blob.ListOptions{
		Prefix: prefix,
		Delimiter: "/",
	})
	for {
		o, err := i.Next(ctx)
		if err != nil {
			break
		}
		key := o.Key[len(prefix):]
		if len(key) == 0 {
			continue
		}
		fill(strings.TrimRight(key, "/"), nil, 0)
	}
	return 0
}

func (cf *CloudFileSystem) Read(path string, buff []byte, ofst int64, fh uint64) (n int) {
	name := strings.TrimLeft(path, "/")
	ctx := context.Background()
	reader, err := cf.bucket.NewRangeReader(ctx, name, ofst, int64(len(buff)), nil)
	if err != nil {
		return
	}
	defer reader.Close()
	n, _ = reader.Read(buff)
	return
}


// Open opens a file.
// The flags are a combination of the fuse.O_* constants.
// The FileSystemBase implementation returns -ENOSYS.
func (cf *CloudFileSystem) Open(path string, flags int) (n int, u uint64) {
	fmt.Println(path)
	fmt.Println(flags)
	// return -ENOSYS, ^uint64(0)
	return
}

func main()  {
	ctx := context.Background()
	if len(os.Args) < 3 {
		fmt.Printf("%s [bucket-path] [mount-point] etc...", os.Args[0])
		os.Exit(1)
	}
	b, err := blob.OpenBucket(ctx, os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b.Close()
	cf := &CloudFileSystem{bucket: b}
	host := fuse.NewFileSystemHost(cf)
	host.Mount(os.Args[2], os.Args[3:])
}