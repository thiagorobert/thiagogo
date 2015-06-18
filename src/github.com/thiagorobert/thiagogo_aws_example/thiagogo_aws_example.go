package main
 
import (
    "fmt"
    "os"
    "path/filepath"
 
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)
 
var (
    Bucket         = "thiagogo-s3-bucket" // Download from this bucket
    Prefix         = ""    // Using this key prefix
    LocalDirectory = "/tmp/s3data"   // Into this directory
)
 
func main() {
    fmt.Printf("starting main \n")
    manager := s3manager.NewDownloader(nil)
    d := downloader{bucket: Bucket, dir: LocalDirectory, Downloader: manager}
 
    client := s3.New(nil)
    params := &s3.ListObjectsInput{Bucket: &Bucket, Prefix: &Prefix}
    list := client.ListObjectsPages(params, d.eachPage)
    fmt.Printf("%v", list)
    fmt.Printf("ending main \n")
}
 
type downloader struct {
    *s3manager.Downloader
    bucket, dir string
}
 
func (d *downloader) eachPage(page *s3.ListObjectsOutput, more bool) bool {
    for _, obj := range page.Contents {
        d.downloadToFile(*obj.Key)
    }
 
    return true
}
 
func (d *downloader) downloadToFile(key string) {
    // Create the directories in the path
    file := filepath.Join(d.dir, key)
    if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
        panic(err)
    }
 
    // Setup the local file
    fd, err := os.Create(file)
    if err != nil {
        panic(err)
    }

    defer fd.Close()

    // Download the file using the AWS SDK
    fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, key, file)
    params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &key}
    d.Download(fd, params)
}

