// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

type Vector3 struct {
	X, Y, Z, _ float32
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Point3 struct {
	X, Y, Z, _ float32
}

type Quat struct {
	X, Y, Z, W float32
}

type Matrix3 struct {
	col0, col1, col2 Vector3
}

type Matrix4 struct {
	col0, col1, col2, col3 Vector4
}

type Transform3 struct {
	col0, col1, col2, col3 Vector3
}
