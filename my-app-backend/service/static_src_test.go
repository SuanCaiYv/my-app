package service

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestStaticSrcApiHandler_ADownloadFile(t *testing.T) {
	fmt.Println(filepath.Ext("/aaa/bbb/ccc.eee")[1:])
}
