/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package runtime

import (
	"strings"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go"
	cesql "github.com/cloudevents/sdk-go/sql/v2"
	"github.com/cloudevents/sdk-go/sql/v2/function"
)

func Test_functionTable_AddFunction(t *testing.T) {

	var HasPrefixFunction cesql.Function = function.NewFunction(
		"HASPREFIX",
		[]cesql.Type{cesql.StringType, cesql.StringType},
		nil,
		func(event cloudevents.Event, i []interface{}) (interface{}, error) {
			str := i[0].(string)
			prefix := i[1].(string)

			return strings.HasPrefix(str, prefix), nil
		},
	)

	AddFunction(HasPrefixFunction)

	type args struct {
		function cesql.Function
	}
	tests := []struct {
		name    string
		table   functionTable
		args    args
		wantErr bool
	}{
		{
			name:  "Add HASPREFIX",
			table: globalFunctionTable,
			args: args{
				function: HasPrefixFunction,
			},
			wantErr: false,
		},
		{
			name:  "Fail add, HASPREFIX exists",
			table: globalFunctionTable,
			args: args{
				function: HasPrefixFunction,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.table.AddFunction(tt.args.function); (err != nil) != tt.wantErr {
				t.Errorf("functionTable.AddFunction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddFunction(t *testing.T) {
	type args struct {
		fn cesql.Function
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddFunction(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("AddFunction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
