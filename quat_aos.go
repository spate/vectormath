// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

import "fmt"

func QCopy(result, quat *Quat) {
	result.X = quat.X
	result.Y = quat.Y
	result.Z = quat.Z
	result.W = quat.W
}

func QMakeFromElems(result *Quat, x, y, z, w float32) {
	result.X = x
	result.Y = y
	result.Z = z
	result.W = w
}

func QMakeFromV3Scalar(result *Quat, xyz *Vector3, w float32) {
	result.SetXYZ(xyz)
	result.SetW(w)
}

func QMakeFromV4(result *Quat, vec *Vector4) {
	result.X = vec.X
	result.Y = vec.Y
	result.Z = vec.Z
	result.W = vec.W
}

func QMakeFromScalar(result *Quat, scalar float32) {
	result.X = scalar
	result.Y = scalar
	result.Z = scalar
	result.W = scalar
}

func QMakeIdentity(result *Quat) {
	QMakeFromElems(result, 0.0, 0.0, 0.0, 1.0)
}

func QLerp(result *Quat, t float32, quat0, quat1 *Quat) {
	var tmpQ_0, tmpQ_1 Quat
	QSub(&tmpQ_0, quat1, quat0)
	QScalarMul(&tmpQ_1, &tmpQ_0, t)
	QAdd(result, quat0, &tmpQ_1)
}

func QSlerp(result *Quat, t float32, unitQuat0, unitQuat1 *Quat) {
	var start, tmpQ_0, tmpQ_1 Quat
	var scale0, scale1 float32
	cosAngle := QDot(unitQuat0, unitQuat1)
	if cosAngle < 0.0 {
		cosAngle = -cosAngle
		QNeg(&start, unitQuat0)
	} else {
		QCopy(&start, unitQuat0)
	}
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := (1.0 / sin(angle))
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = (1.0 - t)
		scale1 = t
	}
	QScalarMul(&tmpQ_0, &start, scale0)
	QScalarMul(&tmpQ_1, unitQuat1, scale1)
	QAdd(result, &tmpQ_0, &tmpQ_1)
}

func QSquad(result *Quat, t float32, unitQuat0, unitQuat1, unitQuat2, unitQuat3 *Quat) {
	var tmp0, tmp1 Quat
	QSlerp(&tmp0, t, unitQuat0, unitQuat3)
	QSlerp(&tmp1, t, unitQuat1, unitQuat2)
	QSlerp(result, (2.0*t)*(1.0-t), &tmp0, &tmp1)
}

func (q *Quat) SetXYZ(vec *Vector3) {
	q.X = vec.X
	q.Y = vec.Y
	q.Z = vec.Z
}

func (q *Quat) SetX(x float32) {
	q.X = x
}

func (q *Quat) SetY(y float32) {
	q.Y = y
}

func (q *Quat) SetZ(z float32) {
	q.Z = z
}

func (q *Quat) SetW(w float32) {
	q.W = w
}

func (q *Quat) SetElem(index int, value float32) {
	switch index {
	case 0:
		q.X = value
	case 1:
		q.Y = value
	case 2:
		q.Z = value
	case 3:
		q.W = value
	}
}

func (q *Quat) GetElem(index int) float32 {
	switch index {
	case 0:
		return q.X
	case 1:
		return q.Y
	case 2:
		return q.Z
	case 3:
		return q.W
	}
	return 0
}

func QAdd(result, quat0, quat1 *Quat) {
	result.X = quat0.X + quat1.X
	result.Y = quat0.Y + quat1.Y
	result.Z = quat0.Z + quat1.Z
	result.W = quat0.W + quat1.W
}

func QSub(result, quat0, quat1 *Quat) {
	result.X = quat0.X - quat1.X
	result.Y = quat0.Y - quat1.Y
	result.Z = quat0.Z - quat1.Z
	result.W = quat0.W - quat1.W
}

func QScalarMul(result, quat *Quat, scalar float32) {
	result.X = quat.X * scalar
	result.Y = quat.Y * scalar
	result.Z = quat.Z * scalar
	result.W = quat.W * scalar
}

func QScalarDiv(result, quat *Quat, scalar float32) {
	result.X = quat.X / scalar
	result.Y = quat.Y / scalar
	result.Z = quat.Z / scalar
	result.W = quat.W / scalar
}

func QNeg(result, quat *Quat) {
	result.X = -quat.X
	result.Y = -quat.Y
	result.Z = -quat.Z
	result.W = -quat.W
}

func QDot(quat0, quat1 *Quat) float32 {
	result := quat0.X * quat1.X
	result += quat0.Y * quat1.Y
	result += quat0.Z * quat1.Z
	result += quat0.W * quat1.W
	return result
}

func (q *Quat) Dot(quat1 *Quat) float32 {
	result := q.X * quat1.X
	result += q.Y * quat1.Y
	result += q.Z * quat1.Z
	result += q.W * quat1.W
	return result
}

func (q *Quat) Norm() float32 {
	result := q.X * q.X
	result += q.Y * q.Y
	result += q.Z * q.Z
	result += q.W * q.W
	return result
}

func (q *Quat) Length() float32 {
	return sqrt(q.Norm())
}

func QNormalize(result, quat *Quat) {
	lenSqr := quat.Norm()
	lenInv := 1.0 / sqrt(lenSqr)
	result.X = quat.X * lenInv
	result.Y = quat.Y * lenInv
	result.Z = quat.Z * lenInv
	result.W = quat.W * lenInv
}

func QMakeRotationArc(result *Quat, unitVec0, unitVec1 *Vector3) {
	var tmpV3_0, tmpV3_1 Vector3
	cosHalfAngleX2 := sqrt((2.0 * (1.0 + V3Dot(unitVec0, unitVec1))))
	recipCosHalfAngleX2 := (1.0 / cosHalfAngleX2)
	V3Cross(&tmpV3_0, unitVec0, unitVec1)
	V3ScalarMul(&tmpV3_1, &tmpV3_0, recipCosHalfAngleX2)
	QMakeFromV3Scalar(result, &tmpV3_1, (cosHalfAngleX2 * 0.5))
}

func QMakeRotationAxis(result *Quat, radians float32, unitVec *Vector3) {
	var tmpV3_0 Vector3
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	V3ScalarMul(&tmpV3_0, unitVec, s)
	QMakeFromV3Scalar(result, &tmpV3_0, c)
}

func QMakeRotationX(result *Quat, radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	QMakeFromElems(result, s, 0.0, 0.0, c)
}

func QMakeRotationY(result *Quat, radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	QMakeFromElems(result, 0.0, s, 0.0, c)
}

func QMakeRotationZ(result *Quat, radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	QMakeFromElems(result, 0.0, 0.0, s, c)
}

func QMul(result, quat0, quat1 *Quat) {
	tmpX := (quat0.W * quat1.X) + (quat0.X * quat1.W) + (quat0.Y * quat1.Z) - (quat0.Z * quat1.Y)
	tmpY := (quat0.W * quat1.Y) + (quat0.Y * quat1.W) + (quat0.Z * quat1.X) - (quat0.X * quat1.Z)
	tmpZ := (quat0.W * quat1.Z) + (quat0.Z * quat1.W) + (quat0.X * quat1.Y) - (quat0.Y * quat1.X)
	tmpW := (quat0.W * quat1.W) - (quat0.X * quat1.X) - (quat0.Y * quat1.Y) - (quat0.Z * quat1.Z)
	QMakeFromElems(result, tmpX, tmpY, tmpZ, tmpW)
}

func QRotate(result *Vector3, quat *Quat, vec *Vector3) {
	tmpX := (quat.W * vec.X) + (quat.Y * vec.Z) - (quat.Z * vec.Y)
	tmpY := (quat.W * vec.Y) + (quat.Z * vec.X) - (quat.X * vec.Z)
	tmpZ := (quat.W * vec.Z) + (quat.X * vec.Y) - (quat.Y * vec.X)
	tmpW := (quat.X * vec.X) + (quat.Y * vec.Y) + (quat.Z * vec.Z)
	result.X = (tmpW * quat.X) + (tmpX * quat.W) - (tmpY * quat.Z) + (tmpZ * quat.Y)
	result.Y = (tmpW * quat.Y) + (tmpY * quat.W) - (tmpZ * quat.X) + (tmpX * quat.Z)
	result.Z = (tmpW * quat.Z) + (tmpZ * quat.W) - (tmpX * quat.Y) + (tmpY * quat.X)
}

func QConj(result, quat *Quat) {
	QMakeFromElems(result, -quat.X, -quat.Y, -quat.Z, quat.W)
}

func QSelect(result, quat0, quat1 *Quat, select1 int) {
	if select1 != 0 {
		result.X = quat1.X
		result.Y = quat1.Y
		result.Z = quat1.Z
		result.W = quat1.W
	} else {
		result.X = quat0.X
		result.Y = quat0.Y
		result.Z = quat0.Z
		result.W = quat0.W
	}
}

func (q *Quat) String() string {
	return fmt.Sprintf("( %f %f %f %f )\n", q.X, q.Y, q.Z, q.W)
}
