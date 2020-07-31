// https://github.com/tencentyun/cos-go-sdk-v5
package cos

import (
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"io/ioutil"
	"context"
	"net/http"
	"net/url"
	"os"
	"time"
//	"bytes"
    "strings"
//	"encoding/json"
)

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
        fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}



func conn()(*cos.Client){
	//var	bucket = "ttt-1252957949"
	//var	region = "hongkong"
	cos_uri:="https://ttt-1252957949.cos.ap-hongkong.myqcloud.com"
	u, _ := url.Parse(cos_uri)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID : os.Getenv("TENCENTCLOUD_SECRET_ID"),
			SecretKey : os.Getenv("TENCENTCLOUD_SECRET_KEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	return c
}

func Download(name string)(string){
    c:=conn()
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s:=string(bs)
	fmt.Printf("%s\n", s)
	return s
}


func Upload(name string ,s string){
    c:=conn()
	f := strings.NewReader(s)
	_, err := c.Object.Put(context.Background(), name, f, nil)
	log_status(err)
}


// n2->n1
func Upload1(n1 string,n2 string){
    c:=conn()
	_, err := c.Object.PutFromFile(context.Background(), n1,n2, nil)
	log_status(err)
}

func Upload2( name string, s string, ContentType string){
    c:=conn()
	f:= strings.NewReader(s)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: ContentType,
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "public-read",
			//XCosACL: "private",
		},
	}
	_, err := c.Object.Put(context.Background(), name, f, opt)
	log_status(err)
}

func Del(n string){
    c:=conn()
	_, err := c.Object.Delete(context.Background(), n, nil)
	log_status(err)
}

func test() {
	p:="/tmp/1593505522.sql"
	p1:="/tmp/1.sql"
	s:="test xxx"

	Upload1(p,p)
	Download(p)
	Del(p)

	Upload2(p1,s,"text/html")
	Download(p1)
	Del(p1)
}

func main(){
	//test()
}
