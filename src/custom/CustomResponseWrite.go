package custom

import (
	"bytes"
	"ginPlus/src/customUtil"
	"github.com/gin-gonic/gin"
)

type CustomResponseWrite struct {
	gin.ResponseWriter
	Body *bytes.Buffer
	LogUtil *customUtil.LoggerUtil
	RequestParam map[string]interface{}
}

func (w CustomResponseWrite) Write(b []byte) (int,error){
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWrite) WriteString(s string) (int,error){
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
