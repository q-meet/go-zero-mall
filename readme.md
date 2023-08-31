# maill

## 运行 dtm

```shell
docker run -itd  --name dtm -p 36789:36789 -p 36790:36790  yedf/dtm:latest
```

## 日志

如果需要打印日志带trace_id 需要使用 logx.WithContext(l.ctx).Info 方法

## http 自定义错误修改模板

go-zero生成代码都是基于模板去生成的，如果生成的代码，不符合期望，你可以自行去修改模板代码

初始化模板

```shell
goctl template init
```

api handler 文件

```tpl
package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		result.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}

```
