// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

import "fmt"

const g_SLERP_TOL = 0.999

func V3Copy(result *Vector3, vec *Vector3) {
	result.x = vec.x
	result.y = vec.y
	result.z = vec.z
}

func V3MakeFromElems(result *Vector3, x, y, z float32) {
	result.x = x
	result.y = y
	result.z = z
}

func V3MakeFromP3(result *Vector3, pnt *Point3) {
	result.x = pnt.x
	result.y = pnt.y
	result.z = pnt.z
}

func V3MakeFromScalar(result *Vector3, scalar float32) {
	result.x = scalar
	result.y = scalar
	result.z = scalar
}

func V3MakeXAxis(result *Vector3) {
	V3MakeFromElems(result, 1.0, 0.0, 0.0)
}

func V3MakeYAxis(result *Vector3) {
	V3MakeFromElems(result, 0.0, 1.0, 0.0)
}

func V3MakeZAxis(result *Vector3) {
	V3MakeFromElems(result, 1.0, 0.0, 1.0)
}

func V3Lerp(result *Vector3, t float32, vec0, vec1 *Vector3) {
	var tmpV3_0, tmpV3_1 Vector3
	V3Sub(&tmpV3_0, vec1, vec0)
	V3ScalarMul(&tmpV3_1, &tmpV3_0, t)
	V3Add(result, vec0, &tmpV3_1)
}

func V3Slerp(result *Vector3, t float32, unitVec0, unitVec1 *Vector3) {
	var tmpV3_0, tmpV3_1 Vector3
	var scale0, scale1 float32
	cosAngle := V3Dot(unitVec0, unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := 1.0 / sin(angle)
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = 1.0 - t
		scale1 = t
	}
	V3ScalarMul(&tmpV3_0, unitVec0, scale0)
	V3ScalarMul(&tmpV3_1, unitVec1, scale1)
	V3Add(result, &tmpV3_0, &tmpV3_1)
}

func (v *Vector3) SetX(x float32) {
	v.x = x
}

func (v *Vector3) X() float32 {
	return v.x
}

func (v *Vector3) SetY(y float32) {
	v.y = y
}

func (v *Vector3) Y() float32 {
	return v.y
}

func (v *Vector3) SetZ(z float32) {
	v.z = z
}

func (v *Vector3) Z() float32 {
	return v.z
}

func (v *Vector3) SetElem(index int, value float32) {
	switch index {
	case 0:
		v.x = value
	case 1:
		v.y = value
	case 2:
		v.z = value
	}
}

func (v *Vector3) GetElem(index int) float32 {
	switch index {
	case 0:
		return v.x
	case 1:
		return v.y
	case 2:
		return v.z
	}
	return 0
}

func V3Add(result, vec0, vec1 *Vector3) {
	result.x = vec0.x + vec1.x
	result.y = vec0.y + vec1.y
	result.z = vec0.z + vec1.z
}

func V3Sub(result, vec0, vec1 *Vector3) {
	result.x = vec0.x - vec1.x
	result.y = vec0.y - vec1.y
	result.z = vec0.z - vec1.z
}

func V3AddP3(result, vec0 *Vector3, pnt1 *Point3) {
	result.x = vec0.x + pnt1.x
	result.y = vec0.y + pnt1.y
	result.z = vec0.z + pnt1.z
}

func V3ScalarMul(result, vec *Vector3, scalar float32) {
	result.x = vec.x * scalar
	result.y = vec.y * scalar
	result.z = vec.z * scalar
}

func V3ScalarDiv(result, vec *Vector3, scalar float32) {
	result.x = vec.x / scalar
	result.y = vec.y / scalar
	result.z = vec.z / scalar
}

func V3Neg(result, vec *Vector3) {
	result.x = -vec.x
	result.y = -vec.y
	result.z = -vec.z
}

func V3MulPerElem(result, vec0, vec1 *Vector3) {
	result.x = vec0.x * vec1.x
	result.y = vec0.y * vec1.y
	result.z = vec0.z * vec1.z
}

func V3DivPerElem(result, vec0, vec1 *Vector3) {
	result.x = vec0.x / vec1.x
	result.y = vec0.y / vec1.y
	result.z = vec0.z / vec1.z
}

func V3RecipPerElem(result, vec *Vector3) {
	result.x = 1.0 / vec.x
	result.y = 1.0 / vec.y
	result.z = 1.0 / vec.z
}

func V3SqrtPerElem(result, vec *Vector3) {
	result.x = sqrt(vec.x)
	result.y = sqrt(vec.y)
	result.z = sqrt(vec.z)
}

func V3RsqrtPerElem(result, vec *Vector3) {
	result.x = 1.0 / sqrt(vec.x)
	result.y = 1.0 / sqrt(vec.y)
	result.z = 1.0 / sqrt(vec.z)
}

func V3AbsPerElem(result, vec *Vector3) {
	result.x = abs(vec.x)
	result.y = abs(vec.y)
	result.z = abs(vec.z)
}

func V3CopySignPerElem(result, vec0, vec1 *Vector3) {
	if vec1.x < 0.0 {
		result.x = -abs(vec0.x)
	} else {
		result.x = abs(vec0.x)
	}
	if vec1.y < 0.0 {
		result.y = -abs(vec0.y)
	} else {
		result.y = abs(vec0.y)
	}
	if vec1.z < 0.0 {
		result.z = -abs(vec0.z)
	} else {
		result.z = abs(vec0.z)
	}
}

func V3MaxPerElem(result, vec0, vec1 *Vector3) {
	result.x = max(vec0.x, vec1.x)
	result.y = max(vec0.y, vec1.y)
	result.z = max(vec0.z, vec1.z)
}

func (v *Vector3) MaxElem() float32 {
	var result float32
	result = max(v.x, v.y)
	result = max(v.z, result)
	return result
}

func V3MinPerElem(result, vec0, vec1 *Vector3) {
	result.x = min(vec0.x, vec1.x)
	result.y = min(vec0.y, vec1.y)
	result.z = min(vec0.z, vec1.z)
}

func (v *Vector3) MinElem() float32 {
	var result float32
	result = min(v.x, v.y)
	result = min(v.z, result)
	return result
}

func (v *Vector3) Sum() float32 {
	var result float32
	result = v.x + v.y + v.z
	return result
}

func V3Dot(vec0, vec1 *Vector3) float32 {
	result := vec0.x * vec1.x
	result += vec0.y * vec1.y
	result += vec0.z * vec1.z
	return result
}

func (v *Vector3) Dot(vec1 *Vector3) float32 {
	result := v.x * vec1.x
	result += v.y * vec1.y
	result += v.z * vec1.z
	return result
}

func (v *Vector3) LengthSqr() float32 {
	result := v.x * v.x
	result += v.y * v.y
	result += v.z * v.z
	return result
}

func (v *Vector3) Length() float32 {
	return sqrt(v.LengthSqr())
}

func V3Normalize(result, vec *Vector3) {
	lenSqr := vec.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result.x = vec.x * lenInv
	result.y = vec.y * lenInv
	result.z = vec.z * lenInv
}

func V3Cross(result, vec0, vec1 *Vector3) {
	tmpX := vec0.y*vec1.z - vec0.z*vec1.y
	tmpY := vec0.z*vec1.x - vec0.x*vec1.z
	tmpZ := vec0.x*vec1.y - vec0.y*vec1.x
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func V3Select(result, vec0, vec1 *Vector3, select1 int) {
	if select1 != 0 {
		result.x = vec1.x
		result.y = vec1.y
		result.z = vec1.z
	} else {
		result.x = vec0.x
		result.y = vec0.y
		result.z = vec0.z
	}
}

func (v *Vector3) String() string {
	return fmt.Sprintf("( %f %f %f )\n", v.x, v.y, v.z)
}

/*******/

func V4Copy(result, vec *Vector4) {
	result.x = vec.x
	result.y = vec.y
	result.z = vec.z
	result.w = vec.w
}

func V4MakeFromElems(result *Vector4, x, y, z, w float32) {
	result.x = x
	result.y = y
	result.z = z
	result.w = w
}

func V4MakeFromV3Scalar(result *Vector4, xyz *Vector3, w float32) {
	result.SetXYZ(xyz)
	result.SetW(w)
}

func V4MakeFromV3(result *Vector4, vec *Vector3) {
	result.x = vec.x
	result.y = vec.y
	result.z = vec.z
	result.w = 0.0
}

func V4MakeFromP3(result *Vector4, pnt *Point3) {
	result.x = pnt.x
	result.y = pnt.y
	result.z = pnt.z
	result.w = 1.0
}

func V4MakeFromQ(result *Vector4, quat *Quat) {
	result.x = quat.x
	result.y = quat.y
	result.z = quat.z
	result.w = quat.w
}

func V4MakeFromScalar(result *Vector4, scalar float32) {
	result.x = scalar
	result.y = scalar
	result.z = scalar
	result.w = scalar
}

func V4MakeXAxis(result *Vector4) {
	V4MakeFromElems(result, 1.0, 0.0, 0.0, 0.0)
}

func V4MakeYAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 1.0, 0.0, 0.0)
}

func V4MakeZAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 0.0, 1.0, 0.0)
}

func V4MakeWAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 0.0, 0.0, 1.0)
}

func V4Lerp(result *Vector4, t float32, vec0, vec1 *Vector4) {
	var tmpV4_0, tmpV4_1 Vector4
	V4Sub(&tmpV4_0, vec1, vec0)
	V4ScalarMul(&tmpV4_1, &tmpV4_0, t)
	V4Add(result, vec0, &tmpV4_1)
}

func V4Slerp(result *Vector4, t float32, unitVec0, unitVec1 *Vector4) {
	var tmpV4_0, tmpV4_1 Vector4
	var scale0, scale1 float32
	cosAngle := V4Dot(unitVec0, unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := (1.0 / sin(angle))
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = (1.0 - t)
		scale1 = t
	}
	V4ScalarMul(&tmpV4_0, unitVec0, scale0)
	V4ScalarMul(&tmpV4_1, unitVec1, scale1)
	V4Add(result, &tmpV4_0, &tmpV4_1)
}

func (v *Vector4) SetXYZ(vec *Vector3) {
	v.x = vec.x
	v.y = vec.y
	v.z = vec.z
}

func V4GetXYZ(result *Vector3, vec *Vector4) {
	V3MakeFromElems(result, vec.x, vec.y, vec.z)
}

func (v *Vector4) SetX(x float32) {
	v.x = x
}

func (v *Vector4) X() float32 {
	return v.x
}

func (v *Vector4) SetY(y float32) {
	v.y = y
}

func (v *Vector4) Y() float32 {
	return v.y
}

func (v *Vector4) SetZ(z float32) {
	v.z = z
}

func (v *Vector4) Z() float32 {
	return v.z
}

func (v *Vector4) SetW(w float32) {
	v.w = w
}

func (v *Vector4) W() float32 {
	return v.w
}


func (v *Vector4) SetElem(index int, value float32) {
	switch index {
	case 0:
		v.x = value
	case 1:
		v.y = value
	case 2:
		v.z = value
	case 3:
		v.w = value
	}
}

func (v *Vector4) GetElem(index int) float32 {
	switch index {
	case 0:
		return v.x
	case 1:
		return v.y
	case 2:
		return v.z
	case 3:
		return v.w
	}
	return 0
}

func V4Add(result, vec0, vec1 *Vector4) {
	result.x = vec0.x + vec1.x
	result.y = vec0.y + vec1.y
	result.z = vec0.z + vec1.z
	result.w = vec0.w + vec1.w
}

func V4Sub(result, vec0, vec1 *Vector4) {
	result.x = vec0.x - vec1.x
	result.y = vec0.y - vec1.y
	result.z = vec0.z - vec1.z
	result.w = vec0.w - vec1.w
}

func V4ScalarMul(result, vec *Vector4, scalar float32) {
	result.x = vec.x * scalar
	result.y = vec.y * scalar
	result.z = vec.z * scalar
	result.w = vec.w * scalar
}

func V4ScalarDiv(result, vec *Vector4, scalar float32) {
	result.x = vec.x / scalar
	result.y = vec.y / scalar
	result.z = vec.z / scalar
	result.w = vec.w / scalar
}

func V4Neg(result, vec *Vector4) {
	result.x = -vec.x
	result.y = -vec.y
	result.z = -vec.z
	result.w = -vec.w
}

func V4MulPerElem(result, vec0, vec1 *Vector4) {
	result.x = vec0.x * vec1.x
	result.y = vec0.y * vec1.y
	result.z = vec0.z * vec1.z
	result.w = vec0.w * vec1.w
}

func V4DivPerElem(result, vec0, vec1 *Vector4) {
	result.x = vec0.x / vec1.x
	result.y = vec0.y / vec1.y
	result.z = vec0.z / vec1.z
	result.w = vec0.w / vec1.w
}

func V4RecipPerElem(result, vec *Vector4) {
	result.x = 1.0 / vec.x
	result.y = 1.0 / vec.y
	result.z = 1.0 / vec.z
	result.w = 1.0 / vec.w
}

func V4SqrtPerElem(result, vec *Vector4) {
	result.x = sqrt(vec.x)
	result.y = sqrt(vec.y)
	result.z = sqrt(vec.z)
	result.w = sqrt(vec.w)
}

func V4RsqrtPerElem(result, vec *Vector4) {
	result.x = 1.0 / sqrt(vec.x)
	result.y = 1.0 / sqrt(vec.y)
	result.z = 1.0 / sqrt(vec.z)
	result.w = 1.0 / sqrt(vec.w)
}

func V4AbsPerElem(result, vec *Vector4) {
	result.x = abs(vec.x)
	result.y = abs(vec.y)
	result.z = abs(vec.z)
	result.w = abs(vec.w)
}

func V4CopySignPerElem(result, vec0, vec1 *Vector4) {
	if vec1.x < 0.0 {
		result.x = -abs(vec0.x)
	} else {
		result.x = abs(vec0.x)
	}
	if vec1.y < 0.0 {
		result.y = -abs(vec0.y)
	} else {
		result.y = abs(vec0.y)
	}
	if vec1.z < 0.0 {
		result.z = -abs(vec0.z)
	} else {
		result.z = abs(vec0.z)
	}
	if vec1.w < 0.0 {
		result.w = -abs(vec0.w)
	} else {
		result.w = abs(vec0.w)
	}
}

func V4MaxPerElem(result, vec0, vec1 *Vector4) {
	result.x = max(vec0.x, vec1.x)
	result.y = max(vec0.y, vec1.y)
	result.z = max(vec0.z, vec1.z)
	result.w = max(vec0.w, vec1.w)
}

func (v *Vector4) MaxElem() float32 {
	var result float32
	result = max(v.x, v.y)
	result = max(v.z, result)
	result = max(v.w, result)
	return result
}

func V4MinPerElem(result, vec0, vec1 *Vector4) {
	result.x = min(vec0.x, vec1.x)
	result.y = min(vec0.y, vec1.y)
	result.z = min(vec0.z, vec1.z)
	result.w = min(vec0.w, vec1.w)
}

func (v *Vector4) MinElem() float32 {
	var result float32
	result = min(v.x, v.y)
	result = min(v.z, result)
	result = min(v.w, result)
	return result
}

func (v *Vector4) Sum() float32 {
	var result float32
	result = v.x + v.y + v.z + v.w
	return result
}

func V4Dot(vec0, vec1 *Vector4) float32 {
	result := vec0.x * vec1.x
	result += vec0.y * vec1.y
	result += vec0.z * vec1.z
	result += vec0.w * vec1.w
	return result
}

func (v *Vector4) Dot(vec1 *Vector4) float32 {
	result := v.x * vec1.x
	result += v.y * vec1.y
	result += v.z * vec1.z
	result += v.w * vec1.w
	return result
}

func (v *Vector4) LengthSqr() float32 {
	result := v.x * v.x
	result += v.y * v.y
	result += v.z * v.z
	result += v.w * v.w
	return result
}

func (v *Vector4) Length() float32 {
	return sqrt(v.LengthSqr())
}

func V4Normalize(result, vec *Vector4) {
	lenSqr := vec.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result.x = vec.x * lenInv
	result.y = vec.y * lenInv
	result.z = vec.z * lenInv
	result.w = vec.w * lenInv
}

func V4Select(result, vec0, vec1 *Vector4, select1 int) {
	if select1 != 0 {
		result.x = vec1.x
		result.y = vec1.y
		result.z = vec1.z
		result.w = vec1.w
	} else {
		result.x = vec0.x
		result.y = vec0.y
		result.z = vec0.z
		result.w = vec0.w
	}
}

func (v *Vector4) String() string {
	return fmt.Sprintf("( %f %f %f %f )", v.x, v.y, v.z, v.w)
}

/*******/

func P3Copy(result, pnt *Point3) {
	result.x = pnt.x
	result.y = pnt.y
	result.z = pnt.z
}

func P3MakeFromElems(result *Point3, x, y, z float32) {
	result.x = x
	result.y = y
	result.z = z
}

func P3MakeFromV3(result *Point3, vec *Vector3) {
	result.x = vec.x
	result.y = vec.y
	result.z = vec.z
}

func P3MakeFromScalar(result *Point3, scalar float32) {
	result.x = scalar
	result.y = scalar
	result.z = scalar
}

func P3Lerp(result *Point3, t float32, pnt0, pnt1 *Point3) {
	var tmpV3_0, tmpV3_1 Vector3
	P3Sub(&tmpV3_0, pnt1, pnt0)
	V3ScalarMul(&tmpV3_1, &tmpV3_0, t)
	P3AddV3(result, pnt0, &tmpV3_1)
}

func (p *Point3) SetX(x float32) {
	p.x = x
}

func (p *Point3) X() float32 {
	return p.x
}

func (p *Point3) SetY(y float32) {
	p.y = y
}

func (p *Point3) Y() float32 {
	return p.y
}

func (p *Point3) SetZ(z float32) {
	p.z = z
}

func (p *Point3) Z() float32 {
	return p.z
}

func (p *Point3) SetElem(index int, value float32) {
	switch index {
	case 0:
		p.x = value
	case 1:
		p.y = value
	case 2:
		p.z = value
	}
}

func (p *Point3) GetElem(index int) float32 {
	switch index {
	case 0:
		return p.x
	case 1:
		return p.y
	case 2:
		return p.z
	}
	return 0
}

func P3Sub(result *Vector3, pnt0, pnt1 *Point3) {
	result.x = pnt0.x - pnt1.x
	result.y = pnt0.y - pnt1.y
	result.z = pnt0.z - pnt1.z
}

func P3AddV3(result, pnt0 *Point3, vec1 *Vector3) {
	result.x = pnt0.x + vec1.x
	result.y = pnt0.y + vec1.y
	result.z = pnt0.z + vec1.z
}

func P3SubV3(result, pnt0 *Point3, vec1 *Vector3) {
	result.x = pnt0.x - vec1.x
	result.y = pnt0.y - vec1.y
	result.z = pnt0.z - vec1.z
}

func P3MulPerElem(result, pnt0, pnt1 *Point3) {
	result.x = pnt0.x * pnt1.x
	result.y = pnt0.y * pnt1.y
	result.z = pnt0.z * pnt1.z
}

func P3DivPerElem(result, pnt0, pnt1 *Point3) {
	result.x = pnt0.x / pnt1.x
	result.y = pnt0.y / pnt1.y
	result.z = pnt0.z / pnt1.z
}

func P3RecipPerElem(result, pnt *Point3) {
	result.x = 1.0 / pnt.x
	result.y = 1.0 / pnt.y
	result.z = 1.0 / pnt.z
}

func P3SqrtPerElem(result, pnt *Point3) {
	result.x = sqrt(pnt.x)
	result.y = sqrt(pnt.y)
	result.z = sqrt(pnt.z)
}

func P3RsqrtPerElem(result, pnt *Point3) {
	result.x = 1.0 / sqrt(pnt.x)
	result.y = 1.0 / sqrt(pnt.y)
	result.z = 1.0 / sqrt(pnt.z)
}

func P3AbsPerElem(result, pnt *Point3) {
	result.x = abs(pnt.x)
	result.y = abs(pnt.y)
	result.z = abs(pnt.z)
}

func P3CopySignPerElem(result, pnt0, pnt1 *Point3) {
	if pnt1.x < 0.0 {
		result.x = -abs(pnt0.x)
	} else {
		result.x = abs(pnt0.x)
	}
	if pnt1.y < 0.0 {
		result.y = -abs(pnt0.y)
	} else {
		result.y = abs(pnt0.y)
	}
	if pnt1.z < 0.0 {
		result.z = -abs(pnt0.z)
	} else {
		result.z = abs(pnt0.z)
	}
}

func P3MaxPerElem(result, pnt0, pnt1 *Point3) {
	result.x = max(pnt0.x, pnt1.x)
	result.y = max(pnt0.y, pnt1.y)
	result.z = max(pnt0.z, pnt1.z)
}

func (p *Point3) MaxElem() float32 {
	var result float32
	result = max(p.x, p.y)
	result = max(p.z, result)
	return result
}

func P3MinPerElem(result, pnt0, pnt1 *Point3) {
	result.x = min(pnt0.x, pnt1.x)
	result.y = min(pnt0.y, pnt1.y)
	result.z = min(pnt0.z, pnt1.z)
}

func (p *Point3) MinElem() float32 {
	var result float32
	result = min(p.x, p.y)
	result = min(p.z, result)
	return result
}

func (p *Point3) Sum() float32 {
	var result float32
	result = p.x + p.y + p.z
	return result
}

func P3Scale(result, pnt *Point3, scaleVal float32) {
	var tmpP3_0 Point3
	P3MakeFromScalar(&tmpP3_0, scaleVal)
	P3MulPerElem(result, pnt, &tmpP3_0)
}

func P3NonUniformScale(result, pnt *Point3, scaleVec *Vector3) {
	var tmpP3_0 Point3
	P3MakeFromV3(&tmpP3_0, scaleVec)
	P3MulPerElem(result, pnt, &tmpP3_0)
}

func (p *Point3) Projection(unitVec *Vector3) float32 {
	result := p.x * unitVec.x
	result += p.y * unitVec.y
	result += p.z * unitVec.z
	return result
}

func (p *Point3) DistSqrFromOrigin() float32 {
	var tmpV3_0 Vector3
	V3MakeFromP3(&tmpV3_0, p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) DistFromOrigin() float32 {
	var tmpV3_0 Vector3
	V3MakeFromP3(&tmpV3_0, p)
	return tmpV3_0.Length()
}

func (p *Point3) DistSqr(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	P3Sub(&tmpV3_0, pnt1, p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) Dist(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	P3Sub(&tmpV3_0, pnt1, p)
	return tmpV3_0.Length()
}

func P3Select(result, pnt0, pnt1 *Point3, select1 int) {
	if select1 != 0 {
		result.x = pnt1.x
		result.y = pnt1.y
		result.z = pnt1.z
	} else {
		result.x = pnt0.x
		result.y = pnt0.y
		result.z = pnt0.z
	}
}

func (p *Point3) String() string {
	return fmt.Sprintf("( %f %f %f )", p.x, p.y, p.z)
}
