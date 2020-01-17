# 工作流引擎
本项目是根据github.com/antlinker/flow演变而来，原来的项目长时间不更新，并且代码qlang部分无法跑通，所以新开了这个项目进行开发

如果您有什么疑问，请联系wx：debugall，注明【go-workflow】，谢谢

目前这个项目支持如下类型：

1. 开始/结束/终止事件
2. 人工任务
3. 排他网关/并行网关

注：暂时不支持子流程，后续有支持计划

## 工作流设计器

- [Camunda下载地址](https://camunda.com/download/modeler/)
- [文档参考](https://docs.awspaas.com/reference-guide/aws-paas-process-reference-guide/process_structure/activities.html)

## 获取项目

```bash
go get github.com/jmniu/workflow
```

## 使用

### 1. 初始化工作流引擎

```go
package main
import (
    "github.com/jmniu/workflow"
    "github.com/jmniu/workflow/service/db"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
    workflow.Init(
		db.SetDSN("root:123456@tcp(127.0.0.1:3306)/flows?charset=utf8"),
		db.SetTrace(true),
	)
}

```

### 2. 加载工作流文件

```go
	err := workflow.LoadFile("leave.bpmn")
	if err != nil {
		// 处理错误
	}
```

### 3. 发起流程

```go
  input := map[string]interface{}{
	"day": 1,
  }

	result, err := workflow.StartFlow("流程编号", "开始节点编号", "流程发起人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 4. 查询待办流程列表

```go
	todos, err := workflow.QueryTodoFlows("流程编号", "流程待办人ID")
	if err != nil {
		// 处理错误
	}
```

### 5. 处理流程

```go
  input := map[string]interface{}{
	"action": "pass",
  }

  result, err = workflow.HandleFlow("待办流程节点实例ID", "流程处理人ID", input)
	if err != nil {
		// 处理错误
	}
```

### 6. 停止流程

```go
	err := workflow.StopFlow("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```

### 7. 接入WEB流程管理

```go
func main() {
serverOptions := []flow.ServerOption{
	    workflow.ServerStaticRootOption("./web"),
	    workflow.ServerPrefixOption("/flow/"),
	    workflow.ServerMiddlewareOption(filter),
	}

	http.Handle("/flow/", workflow.StartServer(serverOptions...))
}

func filter(ctx *gear.Context) error {
	fmt.Printf("请求参数：%s - %s \n", ctx.Path, ctx.Method)
	return nil
}
```

### 8. 查询流程待办数据

```go
	result,err := workflow.QueryTodoFlows("流程编号","流程处理人ID")
	if err != nil {
		// 处理错误
	}
```

### 9. 查询流程历史数据

```go
result,err := workflow.QueryFlowHistory("待办流程实例ID")
if err != nil {
	// 处理错误
}
```

### 10. 查询已办理的流程实例ID列表

```go
ids,err := workflow.QueryDoneFlowIDs("流程编号","流程处理人ID")
if err != nil {
	// 处理错误
}
```

### 11. 查询节点实例的候选人ID列表

```go
ids,err := workflow.QueryNodeCandidates("待办流程节点实例ID")
if err != nil {
	// 处理错误
}
```

### 12. 停止流程实例

```go
	err := workflow.StopFlowInstance("待办流程节点实例ID", func(flowInstance *schema.FlowInstance) bool {
		return flowInstance.Launcher == "XXX"
	})
	if err != nil {
		// 处理错误
	}
```

![流程管理](example/screenshots/QQ20180123-175942@2x.png)
![流程设计器](example/screenshots/QQ20180123-180022@2x.png)
