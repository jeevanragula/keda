// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDescribeStream struct {
}

func (*validateOpDescribeStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRecords struct {
}

func (*validateOpGetRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetShardIterator struct {
}

func (*validateOpGetShardIterator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetShardIterator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetShardIteratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetShardIteratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDescribeStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeStream{}, middleware.After)
}

func addOpGetRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRecords{}, middleware.After)
}

func addOpGetShardIteratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetShardIterator{}, middleware.After)
}

func validateOpDescribeStreamInput(v *DescribeStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStreamInput"}
	if v.StreamArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRecordsInput(v *GetRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRecordsInput"}
	if v.ShardIterator == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardIterator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetShardIteratorInput(v *GetShardIteratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetShardIteratorInput"}
	if v.StreamArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StreamArn"))
	}
	if v.ShardId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardId"))
	}
	if len(v.ShardIteratorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ShardIteratorType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
