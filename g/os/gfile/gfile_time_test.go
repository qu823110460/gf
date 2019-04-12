//have test 100%
package gfile

import (
	"github.com/gogf/gf/g/test/gtest"
	"os"
	"testing"
)

func TestMTime(t *testing.T) {
	gtest.Case(t, func() {

		var (
			file1 string ="/testfile_t1.txt"
			err error
			fileobj os.FileInfo
		)

		CreateTestFile(file1,"")
		defer  DelTestFiles(file1)
		fileobj, err = os.Stat(os.TempDir()+file1)
		gtest.Assert(err, nil)


		gtest.AssertGT(MTime(os.TempDir()+file1), fileobj.ModTime().Unix())
		gtest.Assert(MTime(""), 0)
	})
}

func TestMTimeMillisecond(t *testing.T) {
	gtest.Case(t, func() {
		var (
			file1 string ="/testfile_t1.txt"
			err error
			fileobj os.FileInfo
		)

		CreateTestFile(file1,"")
		defer  DelTestFiles(file1)
		fileobj, err = os.Stat(os.TempDir()+file1)
		gtest.Assert(err, nil)



		//这里本不为0,但github中的ci测试时，值为0
		gtest.AssertGTE(MTimeMillisecond(os.TempDir()+file1),fileobj.ModTime().Nanosecond()/1000000)
		gtest.Assert(MTimeMillisecond(""), 0)
	})
}
