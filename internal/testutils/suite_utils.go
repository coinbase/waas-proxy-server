package testutils

import (
	"context"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

// TestSuite is the base test suite for tests.
type TestSuite struct {
	suite.Suite
}

// ProtoEqual asserts two protobuf messages match.
// The opts are cmp.Options to control the diffing,
// eg: protocmp.IgnoreField(&foo.FooMsg{}, "created_at").
func (t *TestSuite) ProtoEqual(expected, actual interface{}, opts ...cmp.Option) bool {
	diff := t.ProtoDiff(expected, actual, opts...)

	return t.True(
		diff == "",
		fmt.Sprintf("Protos differ (-want +got):\n%s", diff))
}

// ProtoDiff returns the difference between two protos.
// The opts are cmp.Options to control the diffing,
// eg: protocmp.IgnoreField(&foo.FooMsg{}, "created_at").
func (t *TestSuite) ProtoDiff(expected, actual interface{}, opts ...cmp.Option) string {
	allOpts := append([]cmp.Option{protocmp.Transform()}, opts...)
	diff := cmp.Diff(expected, actual, allOpts...)

	return diff
}

// StatusCodeEqual is a helper function for asserting that the specified error maps to
// the expected gRPC status code.
func (t *TestSuite) StatusCodeEqual(expected codes.Code, err error) {
	t.Require().Error(err, "must return an error")

	t.Equal(
		expected.String(),
		status.Code(err).String(),
		fmt.Sprintf("grpc status should be %s for error: %q", expected.String(), err.Error()),
	)
}

// matchAnyContext is a matcher for context.Context types.
func (t *TestSuite) MatchAnyContext(_ context.Context) bool { return true }
