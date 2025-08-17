package mcp2openai

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/shared"
)

func Convert(tools []*mcp.Tool) []openai.ChatCompletionToolUnionParam {
	openaiTools := make([]openai.ChatCompletionToolUnionParam, 0, len(tools))

	for _, tool := range tools {
		openaiTool := openai.ChatCompletionToolUnionParam{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Type: "function",
				Function: shared.FunctionDefinitionParam{
					Name:        tool.Name,
					Description: openai.String(tool.Description),
					Parameters: shared.FunctionParameters{
						"Type":       "object",
						"Properties": tool.InputSchema.Properties,
						"Required":   tool.InputSchema.Required,
					},
				},
			},
		}

		if len(tool.InputSchema.Properties) == 0 {
			openaiTool.OfFunction.Function.Parameters = nil
		}

		openaiTools = append(openaiTools, openaiTool)
	}

	return openaiTools
}
