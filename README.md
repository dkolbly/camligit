CamliGIT
========

A tool to push GIT repositories into Camlistore

The git hashed object is pushed, so the sha1 blobref in camlistore
matches the GIT ref

usage:

    camligit .GIT-DIRECTORY CAMLISTORE-URL LEVELDB-DIR

e.g.

    $ camligit .git http://me:password@localhost:8671 /tmp/hascache.db
    00:01:15.758 INFO     [camligit|backend.go:99] Hitting http://localhost:8671
    00:01:15.954 DEBUG    [camligit|backend.go:38] sha1-47b52495a0f7bf4baa8a53b7da6aaf6d1fc71c42 226 created=true
    00:01:15.970 DEBUG    [camligit|backend.go:38] sha1-2fc59f1a1db3099c0a4c284d11a7b8e597e426d4 1303 created=true
    00:01:15.984 DEBUG    [camligit|backend.go:38] sha1-460d3ee4da18c22101d6c3a813f9bed00f23b175 157 created=true
    00:01:15.995 DEBUG    [camligit|backend.go:38] sha1-03130ba202a65ba57df0c8eb65048bfb837805a7 2194 created=true
    00:01:16.006 DEBUG    [camligit|backend.go:38] sha1-2e2c88a8767424bd326cb51b52d160809c0fed2b 305 created=true
    00:01:16.006 INFO     [camligit|sync.go:51] Uploaded 5 out of a total of 6
    $ camget sha1-$(git rev-parse HEAD)
    commit 215tree 460d3ee4da18c22101d6c3a813f9bed00f23b175
    author Donovan Kolbly <donovan@rscheme.org> 1479794389 -0600
    committer Donovan Kolbly <donovan@rscheme.org> 1479794389 -0600
    
    Initial commit of git/camlistore sync tool
