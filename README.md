CamliGIT
========

A tool to push GIT repositories into Camlistore

The git hashed object is pushed, so the sha1 blobref in camlistore
matches the GIT ref

usage:

    camligit .GIT-DIRECTORY CAMLISTORE-URL LEVELDB-DIR

e.g.

    camligit .git http://me:password@localhost:8671 /tmp/hascache.db

