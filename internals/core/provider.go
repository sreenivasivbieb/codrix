package core


input {
	"context"
	"encoding/json"
}

type Provider interface{
	complete(ctx context.Context, request CompleteRequest)(CompletionRequest error)

}

type ToolExecutor interface{
    Defination 

}