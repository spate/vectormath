// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

import "fmt"

func QCopy(result, quat *Quat) {
	result.x = quat.x
	result.y = quat.y
	result.z = quat.z
	result.w = quat.w
}

func QMakeFromElems(result *Quat, x, y, z, w float32) {
	result.x = x
	result.y = y
	result.z = z
	result.w = w
}

func QMakeFromV3Scalar(result *Quat, xyz *Vector3, w float32) {
	result.SetXYZ(xyz)
	result.SetW(w)
}

func QMakeFromV4(result *Quat, vec *Vector4) {
	result.x = vec.x
	result.y = vec.y
	result.z = vec.z
	result.w = vec.w
}

func QMakeFromScalar(result *Quat, scalar float32) {
	result.x = scalar
	result.y = scalar
	result.z = scalar
	result.w = scalar
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
	q.x = vec.x
	q.y = vec.y
	q.z = vec.z
}

func (q *Quat) SetX(x float32) {
	q.x = x
}

func (q *Quat) X() float32 {
	return q.x
}

func (q *Quat) SetY(y float32) {
	q.y = y
}

func (q *Quat) Y() float32 {
	return q.y
}

func (q *Quat) SetZ(z float32) {
	q.z = z
}

func (q *Quat) Z() float32 {
	return q.z
}

func (q *Quat) SetW(w float32) {
	q.w = w
}

func (q *Quat) W() float32 {
	return q.w
}

func (q *Quat) SetElem(index int, value float32) {
	switch index {
	case 0:
		q.x = value
	case 1:
		q.y = value
	case 2:
		q.z = value
	case 3:
		q.w = value
	}
}

func (q *Quat) GetElem(index int) float32 {
	switch index {
	case 0:
		return q.x
	case 1:
		return q.y
	case 2:
		return q.z
	case 3:
		return q.w
	}
	return 0
}

func QAdd(result, quat0, quat1 *Quat) {
	result.x = quat0.x + quat1.x
	result.y = quat0.y + quat1.y
	result.z = quat0.z + quat1.z
	result.w = quat0.w + quat1.w
}

func QSub(result, quat0, quat1 *Quat) {
	result.x = quat0.x - quat1.x
	result.y = quat0.y - quat1.y
	result.z = quat0.z - quat1.z
	result.w = quat0.w - quat1.w
}

func QScalarMul(result, quat *Quat, scalar float32) {
	result.x = quat.x * scalar
	result.y = quat.y * scalar
	result.z = quat.z * scalar
	result.w = quat.w * scalar
}

func QScalarDiv(result, quat *Quat, scalar float32) {
	result.x = quat.x / scalar
	result.y = quat.y / scalar
	result.z = quat.z / scalar
	result.w = quat.w / scalar
}

func QNeg(result, quat *Quat) {
	result.x = -quat.x
	result.y = -quat.y
	result.z = -quat.z
	result.w = -quat.w
}

func QDot(quat0, quat1 *Quat) float32 {
	result := quat0.x * quat1.x
	result += quat0.y * quat1.y
	result += quat0.z * quat1.z
	result += quat0.w * quat1.w
	return result
}

func (q *Quat) Dot(quat1 *Quat) float32 {
	result := q.x * quat1.x
	result += q.y * quat1.y
	result += q.z * quat1.z
	result += q.w * quat1.w
	return result
}

func (q *Quat) Norm() float32 {
	result := q.x * q.x
	result += q.y * q.y
	result += q.z * q.z
	result += q.w * q.w
	return result
}

func (q *Quat) Length() float32 {
	return sqrt(q.Norm())
}

func QNormalize(result, quat *Quat) {
	lenSqr := quat.Norm()
	lenInv := 1.0 / sqrt(lenSqr)
	result.x = quat.x * lenInv
	result.y = quat.y * lenInv
	result.z = quat.z * lenInv
	result.w = quat.w * lenInv
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
	tmpX := (quat0.w * quat1.x) + (quat0.x * quat1.w) + (quat0.y * quat1.z) - (quat0.z * quat1.y)
	tmpY := (quat0.w * quat1.y) + (quat0.y * quat1.w) + (quat0.z * quat1.x) - (quat0.x * quat1.z)
	tmpZ := (quat0.w * quat1.z) + (quat0.z * quat1.w) + (quat0.x * quat1.y) - (quat0.y * quat1.x)
	tmpW := (quat0.w * quat1.w) - (quat0.x * quat1.x) - (quat0.y * quat1.y) - (quat0.z * quat1.z)
	QMakeFromElems(result, tmpX, tmpY, tmpZ, tmpW)
}

func QRotate(result *Vector3, quat *Quat, vec *Vector3) {
	tmpX := (quat.w * vec.x) + (quat.y * vec.z) - (quat.z * vec.y)
	tmpY := (quat.w * vec.y) + (quat.z * vec.x) - (quat.x * vec.z)
	tmpZ := (quat.w * vec.z) + (quat.x * vec.y) - (quat.y * vec.x)
	tmpW := (quat.x * vec.x) + (quat.y * vec.y) + (quat.z * vec.z)
	result.x = (tmpW * quat.x) + (tmpX * quat.w) - (tmpY * quat.z) + (tmpZ * quat.y)
	result.y = (tmpW * quat.y) + (tmpY * quat.w) - (tmpZ * quat.x) + (tmpX * quat.z)
	result.z = (tmpW * quat.z) + (tmpZ * quat.w) - (tmpX * quat.y) + (tmpY * quat.x)
}

func QConj(result, quat *Quat) {
	QMakeFromElems(result, -quat.x, -quat.y, -quat.z, quat.w)
}

func QSelect(result, quat0, quat1 *Quat, select1 int) {
	if select1 != 0 {
		result.x = quat1.x
		result.y = quat1.y
		result.z = quat1.z
		result.w = quat1.w
	} else {
		result.x = quat0.x
		result.y = quat0.y
		result.z = quat0.z
		result.w = quat0.w
	}
}

func (q *Quat) String() string {
	return fmt.Sprintf("( %f %f %f %f )\n", q.x, q.y, q.z, q.w)
}
