// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.2

package types

import "k8s.io/kubernetes/third_party/golang/go/ast"

func slice3(x *ast.SliceExpr) bool {
	return false
}

func sliceMax(x *ast.SliceExpr) ast.Expr {
	return nil
}
