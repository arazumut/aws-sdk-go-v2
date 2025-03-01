// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

func ExampleLakeFormationScopeUnion_outputUsage() {
	var union types.LakeFormationScopeUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.LakeFormationScopeUnionMemberLakeFormationQuery:
		_ = v.Value // Value is types.LakeFormationQuery

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LakeFormationQuery

func ExampleNamespaceIdentifierUnion_outputUsage() {
	var union types.NamespaceIdentifierUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.NamespaceIdentifierUnionMemberProvisionedIdentifier:
		_ = v.Value // Value is types.ProvisionedIdentifier

	case *types.NamespaceIdentifierUnionMemberServerlessIdentifier:
		_ = v.Value // Value is types.ServerlessIdentifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ProvisionedIdentifier
var _ *types.ServerlessIdentifier

func ExampleS3AccessGrantsScopeUnion_outputUsage() {
	var union types.S3AccessGrantsScopeUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.S3AccessGrantsScopeUnionMemberReadWriteAccess:
		_ = v.Value // Value is types.ReadWriteAccess

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ReadWriteAccess

func ExampleServiceIntegrationsUnion_outputUsage() {
	var union types.ServiceIntegrationsUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ServiceIntegrationsUnionMemberLakeFormation:
		_ = v.Value // Value is []types.LakeFormationScopeUnion

	case *types.ServiceIntegrationsUnionMemberS3AccessGrants:
		_ = v.Value // Value is []types.S3AccessGrantsScopeUnion

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.LakeFormationScopeUnion
var _ []types.S3AccessGrantsScopeUnion
