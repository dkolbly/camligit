package main

import (
	"bytes"
	"camlistore.org/pkg/blob"
	"camlistore.org/pkg/client"
	"camlistore.org/pkg/search"
	"context"
	//"camlistore.org/pkg/index"
	"github.com/syndtr/goleveldb/leveldb"
	"net/url"
	"strings"
)

type Backend struct {
	camli *client.Client
	//index index.Interface
	cache *leveldb.DB
}

func (be *Backend) Put(ref string, value []byte) error {
	bref, ok := blob.Parse("sha1-" + ref)
	if !ok {
		panic("rats")
	}

	up := &client.UploadHandle{
		BlobRef:  bref,
		Contents: bytes.NewReader(value),
		Size:     uint32(len(value)),
	}

	put, err := be.camli.Upload(up)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Debug("%s %d created=%t", put.BlobRef, put.Size, !put.Skipped)
	err = be.cache.Put([]byte(ref), []byte{'2'}, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (be *Backend) Has(ref string) bool {
	ok, err := be.cache.Has([]byte(ref), nil)
	if err != nil {
		log.Error("Looking up %s: %s", ref, err)
	}

	if ok {
		return true
	}

	// not very efficient, but at least not hacky... how to get access to
	// the a remote index.Interface ?

	bref, ok := blob.Parse("sha1-" + ref)
	if !ok {
		return false
	}

	q := &search.DescribeRequest{
		BlobRefs: []blob.Ref{
			bref,
		},
	}
	desc, err := be.camli.Describe(context.Background(), q)
	if err != nil {
		log.Error("Could not describe %s: %s", ref, err)
		return false
	}
	ok = desc.Meta.Get(bref) != nil
	if ok {
		err = be.cache.Put([]byte(ref), []byte{'1'}, nil)
		if err != nil {
			log.Error("Failed to put: %s", err)
		}
	}
	return ok
}

func NewBackend(uri, cachefile string) (*Backend, error) {

	var c *client.Client

	if strings.IndexByte(uri, ':') < 0 {
		// looks like a camlistore server reference
		c = client.New(uri)
	} else {
		u, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		creds := u.User
		u.User = nil
		log.Info("Hitting %s", u.String())
		c = client.New(u.String())
		err = c.SetupAuthFromString("userpass:" + creds.String())
		if err != nil {
			return nil, err
		}
	}
	ldb, err := leveldb.OpenFile(cachefile, nil)
	if err != nil {
		c.Close()
		return nil, err
	}

	return &Backend{
		camli: c,
		cache: ldb,
	}, nil
}
