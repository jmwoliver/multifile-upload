
# Multiple File Uploads with Go

This repo is a small demo to showcase uploading multiple files in one request with `multipart/form-data`.

## How to Use

In a terminal window, run:

```bash
go run main.go --entry=server
```

In a seperate terminal window, run the client and pass in the files you want to upload to the server:

```bash
go run main.go --entry=client --files='client/files/file1.txt, client/files/file2.txt, client/files/file3.txt'
```

You should see the files appear in the `./server/files/` folder and see the server terminal window output that it was successful:

```bash
Uploading 'file1.txt' to server...
Uploading 'file2.txt' to server...
Uploading 'file3.txt' to server...
Successfully uploaded file(s) to server!
```