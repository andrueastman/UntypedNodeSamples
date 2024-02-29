package main

import (
	"context"
	"fmt"

	client "github.com/andrueastman/untypednodesamples/goclient/goclient"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	r "github.com/microsoft/kiota-http-go"
)

func main() {
	authprovider := &auth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authprovider)
	if err != nil {
		fmt.Printf("Error creating adapter: %v\n", err)
		return
	}
	client := client.NewApiClient(adapter)
	response, err := client.MetricsJson().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting messages: %v\n", err)
		return
	}

	dataSets := response.GetDatasets()
	indent := ""
	parseUntypedNode(dataSets, &indent)
}

func parseUntypedNode(node absser.UntypedNodeable, indent *string) {
	emptyIndent := ""
	switch value := node.(type) {
	case *absser.UntypedObject:
		fmt.Printf("%vFound object value: \n", *indent)
		for key, val := range value.GetValue() {
			fmt.Printf("%vProperty Name: %v \n", *indent, key)
			emptyIndent = *indent + "  "
			parseUntypedNode(val, &emptyIndent)
		}
	case *absser.UntypedArray:
		fmt.Printf("%vFound array value: \n", *indent)
		for _, item := range value.GetValue() {
			fmt.Printf("%vNew Item: \n", *indent)
			emptyIndent = *indent + "  "
			parseUntypedNode(item, &emptyIndent)
		}
	case *absser.UntypedBoolean:
		fmt.Printf("%vFound bool value: %v \n", *indent, *value.GetValue())
	case *absser.UntypedFloat:
		fmt.Printf("%vFound float value: %v \n", *indent, *value.GetValue())
	case *absser.UntypedDouble:
		fmt.Printf("%vFound double value: %v \n", *indent, *value.GetValue())
	case *absser.UntypedInteger:
		fmt.Printf("%vFound int value: %v \n", *indent, *value.GetValue())
	case *absser.UntypedLong:
		fmt.Printf("%vFound long value: %v \n", *indent, *value.GetValue())
	case *absser.UntypedNull:
		fmt.Printf("%vFound null value: %v \n", *indent, value.GetValue())
	case *absser.UntypedString:
		fmt.Printf("%vFound string value: %v \n", *indent, *value.GetValue())
	}
}
