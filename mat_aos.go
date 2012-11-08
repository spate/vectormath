// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

const g_PI_OVER_2 = 1.570796327

func M3Copy(result *Matrix3, mat *Matrix3) {
	V3Copy(&result.col0, &mat.col0)
	V3Copy(&result.col1, &mat.col1)
	V3Copy(&result.col2, &mat.col2)
}

func M3MakeFromScalar(result *Matrix3, scalar float32) {
	V3MakeFromScalar(&result.col0, scalar)
	V3MakeFromScalar(&result.col1, scalar)
	V3MakeFromScalar(&result.col2, scalar)
}

func M3MakeFromQ(result *Matrix3, unitQuat *Quat) {
	qx := unitQuat.x
	qy := unitQuat.x
	qz := unitQuat.x
	qw := unitQuat.x
	qx2 := qx + qx
	qy2 := qy + qy
	qz2 := qz + qz
	qxqx2 := qx * qx2
	qxqy2 := qx * qy2
	qxqz2 := qx * qz2
	qxqw2 := qw * qx2
	qyqy2 := qy * qy2
	qyqz2 := qy * qz2
	qyqw2 := qw * qy2
	qzqz2 := qz * qz2
	qzqw2 := qw * qz2
	V3MakeFromElems(&result.col0, ((1.0 - qyqy2) - qzqz2), (qxqy2 + qzqw2), (qxqz2 - qyqw2))
	V3MakeFromElems(&result.col1, (qxqy2 - qzqw2), ((1.0 - qxqx2) - qzqz2), (qyqz2 + qxqw2))
	V3MakeFromElems(&result.col2, (qxqz2 + qyqw2), (qyqz2 - qxqw2), ((1.0 - qxqx2) - qyqy2))
}

func M3MakeFromCols(result *Matrix3, col0, col1, col2 *Vector3) {
	V3Copy(&result.col0, col0)
	V3Copy(&result.col1, col1)
	V3Copy(&result.col2, col2)
}

func (m *Matrix3) SetCol0(col0 *Vector3) {
	V3Copy(&m.col0, col0)
}

func (m *Matrix3) SetCol1(col1 *Vector3) {
	V3Copy(&m.col1, col1)
}

func (m *Matrix3) SetCol2(col2 *Vector3) {
	V3Copy(&m.col2, col2)
}

func (m *Matrix3) SetCol(col int, vec *Vector3) {
	switch col {
	case 0:
		V3Copy(&m.col0, vec)
	case 1:
		V3Copy(&m.col1, vec)
	case 2:
		V3Copy(&m.col2, vec)
	}
}

func (m *Matrix3) SetRow(row int, vec *Vector3) {
	m.col0.SetElem(row, vec.GetElem(0))
	m.col1.SetElem(row, vec.GetElem(1))
	m.col2.SetElem(row, vec.GetElem(2))
}

func (m *Matrix3) SetElem(col, row int, val float32) {
	var tmpV3_0 Vector3
	M3GetCol(&tmpV3_0, m, col)
	tmpV3_0.SetElem(row, val)
	m.SetCol(col, &tmpV3_0)
}

func (m *Matrix3) GetElem(col, row int) float32 {
	var tmpV3_0 Vector3
	M3GetCol(&tmpV3_0, m, col)
	return tmpV3_0.GetElem(row)
}

func M3GetCol0(result *Vector3, mat *Matrix3) {
	V3Copy(result, &mat.col0)
}

func M3GetCol1(result *Vector3, mat *Matrix3) {
	V3Copy(result, &mat.col1)
}

func M3GetCol2(result *Vector3, mat *Matrix3) {
	V3Copy(result, &mat.col2)
}

func M3GetCol(result *Vector3, mat *Matrix3, col int) {
	switch col {
	case 0:
		V3Copy(result, &mat.col0)
	case 1:
		V3Copy(result, &mat.col1)
	case 2:
		V3Copy(result, &mat.col2)
	}
}

func M3GetRow(result *Vector3, mat *Matrix3, row int) {
	x := mat.col0.GetElem(row)
	y := mat.col1.GetElem(row)
	z := mat.col2.GetElem(row)
	V3MakeFromElems(result, x, y, z)
}

func M3Transpose(result, mat *Matrix3) {
	var tmpResult Matrix3
	V3MakeFromElems(&tmpResult.col0, mat.col0.x, mat.col1.x, mat.col2.x)
	V3MakeFromElems(&tmpResult.col1, mat.col0.y, mat.col1.y, mat.col2.y)
	V3MakeFromElems(&tmpResult.col2, mat.col0.z, mat.col1.z, mat.col2.z)
	M3Copy(result, &tmpResult)
}

func M3Inverse(result, mat *Matrix3) {
	var tmp0, tmp1, tmp2 Vector3
	V3Cross(&tmp0, &mat.col1, &mat.col2)
	V3Cross(&tmp1, &mat.col2, &mat.col0)
	V3Cross(&tmp2, &mat.col0, &mat.col1)
	detinv := 1.0 / V3Dot(&mat.col2, &tmp2)
	V3MakeFromElems(&result.col0, tmp0.x*detinv, tmp1.x*detinv, tmp2.x*detinv)
	V3MakeFromElems(&result.col1, tmp0.y*detinv, tmp1.y*detinv, tmp2.y*detinv)
	V3MakeFromElems(&result.col2, tmp0.z*detinv, tmp1.z*detinv, tmp2.z*detinv)
}

func (m *Matrix3) Determinant() float32 {
	var tmpV3_0 Vector3
	V3Cross(&tmpV3_0, &m.col0, &m.col1)
	return V3Dot(&m.col2, &tmpV3_0)
}

func M3Add(result, mat0, mat1 *Matrix3) {
	V3Add(&result.col0, &mat0.col0, &mat1.col0)
	V3Add(&result.col1, &mat0.col1, &mat1.col1)
	V3Add(&result.col2, &mat0.col2, &mat1.col2)
}

func M3Sub(result, mat0, mat1 *Matrix3) {
	V3Sub(&result.col0, &mat0.col0, &mat1.col0)
	V3Sub(&result.col1, &mat0.col1, &mat1.col1)
	V3Sub(&result.col2, &mat0.col2, &mat1.col2)
}

func M3Neg(result, mat *Matrix3) {
	V3Neg(&result.col0, &mat.col0)
	V3Neg(&result.col1, &mat.col1)
	V3Neg(&result.col2, &mat.col2)
}

func M3AbsPerElem(result, mat *Matrix3) {
	V3AbsPerElem(&result.col0, &mat.col0)
	V3AbsPerElem(&result.col1, &mat.col1)
	V3AbsPerElem(&result.col2, &mat.col2)
}

func M3ScalarMul(result, mat *Matrix3, scalar float32) {
	V3ScalarMul(&result.col0, &mat.col0, scalar)
	V3ScalarMul(&result.col1, &mat.col1, scalar)
	V3ScalarMul(&result.col2, &mat.col2, scalar)
}

func M3MulV3(result *Vector3, mat *Matrix3, vec *Vector3) {
	tmpX := ((mat.col0.x * vec.x) + (mat.col1.x * vec.y)) + (mat.col2.x * vec.z)
	tmpY := ((mat.col0.y * vec.x) + (mat.col1.y * vec.y)) + (mat.col2.y * vec.z)
	tmpZ := ((mat.col0.z * vec.x) + (mat.col1.z * vec.y)) + (mat.col2.z * vec.z)
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func M3Mul(result, mat0, mat1 *Matrix3) {
	var tmpResult Matrix3
	M3MulV3(&tmpResult.col0, mat0, &mat1.col0)
	M3MulV3(&tmpResult.col1, mat0, &mat1.col1)
	M3MulV3(&tmpResult.col2, mat0, &mat1.col2)
	M3Copy(result, &tmpResult)
}

func M3MulPerElem(result, mat0, mat1 *Matrix3) {
	V3MulPerElem(&result.col0, &mat0.col0, &mat1.col0)
	V3MulPerElem(&result.col1, &mat0.col1, &mat1.col1)
	V3MulPerElem(&result.col2, &mat0.col2, &mat1.col2)
}

func M3MakeIdentity(result *Matrix3) {
	V3MakeXAxis(&result.col0)
	V3MakeYAxis(&result.col1)
	V3MakeZAxis(&result.col2)
}

func M3MakeRotationX(result *Matrix3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeXAxis(&result.col0)
	V3MakeFromElems(&result.col1, 0.0, c, s)
	V3MakeFromElems(&result.col2, 0.0, -s, c)
}

func M3MakeRotationY(result *Matrix3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.col0, c, 0.0, -s)
	V3MakeYAxis(&result.col1)
	V3MakeFromElems(&result.col2, s, 0.0, c)
}

func M3MakeRotationZ(result *Matrix3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.col0, c, s, 0.0)
	V3MakeFromElems(&result.col1, -s, c, 0.0)
	V3MakeZAxis(&result.col2)
}

func M3MakeRotationZYX(result *Matrix3, radiansXYZ *Vector3) {
	sX := sin(radiansXYZ.x)
	cX := cos(radiansXYZ.x)
	sY := sin(radiansXYZ.y)
	cY := cos(radiansXYZ.y)
	sZ := sin(radiansXYZ.z)
	cZ := cos(radiansXYZ.z)
	tmp0 := cZ * sY
	tmp1 := sZ * sY
	V3MakeFromElems(&result.col0, (cZ * cY), (sZ * cY), -sY)
	V3MakeFromElems(&result.col1, ((tmp0 * sX) - (sZ * cX)), ((tmp1 * sX) + (cZ * cX)), (cY * sX))
	V3MakeFromElems(&result.col2, ((tmp0 * cX) + (sZ * sX)), ((tmp1 * cX) - (cZ * sX)), (cY * cX))
}

func M3MakeRotationAxis(result *Matrix3, radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	x := unitVec.x
	y := unitVec.y
	z := unitVec.z
	xy := x * y
	yz := y * z
	zx := z * x
	oneMinusC := 1.0 - c
	V3MakeFromElems(&result.col0, (((x * x) * oneMinusC) + c), ((xy * oneMinusC) + (z * s)), ((zx * oneMinusC) - (y * s)))
	V3MakeFromElems(&result.col1, ((xy * oneMinusC) - (z * s)), (((y * y) * oneMinusC) + c), ((yz * oneMinusC) + (x * s)))
	V3MakeFromElems(&result.col2, ((zx * oneMinusC) + (y * s)), ((yz * oneMinusC) - (x * s)), (((z * z) * oneMinusC) + c))
}

func M3MakeRotationQ(result *Matrix3, unitQuat *Quat) {
	M3MakeFromQ(result, unitQuat)
}

func M3MakeScale(result *Matrix3, scaleVec *Vector3) {
	V3MakeFromElems(&result.col0, scaleVec.x, 0.0, 0.0)
	V3MakeFromElems(&result.col1, 0.0, scaleVec.y, 0.0)
	V3MakeFromElems(&result.col2, 0.0, 0.0, scaleVec.z)
}

func M3AppendScale(result, mat *Matrix3, scaleVec *Vector3) {
	V3ScalarMul(&result.col0, &mat.col0, scaleVec.x)
	V3ScalarMul(&result.col1, &mat.col1, scaleVec.y)
	V3ScalarMul(&result.col2, &mat.col2, scaleVec.z)
}

func M3PrependScale(result *Matrix3, scaleVec *Vector3, mat *Matrix3) {
	V3MulPerElem(&result.col0, &mat.col0, scaleVec)
	V3MulPerElem(&result.col1, &mat.col1, scaleVec)
	V3MulPerElem(&result.col2, &mat.col2, scaleVec)
}

func M3Select(result, mat0, mat1 *Matrix3, select1 int) {
	V3Select(&result.col0, &mat0.col0, &mat1.col0, select1)
	V3Select(&result.col1, &mat0.col1, &mat1.col1, select1)
	V3Select(&result.col2, &mat0.col2, &mat1.col2, select1)
}

func (m *Matrix3) String() string {
	var tmp Matrix3
	M3Transpose(&tmp, m)
	return tmp.col0.String() + tmp.col1.String() + tmp.col2.String()
}

/*******/

func M4Copy(result, mat *Matrix4) {
	V4Copy(&result.col0, &mat.col0)
	V4Copy(&result.col1, &mat.col1)
	V4Copy(&result.col2, &mat.col2)
	V4Copy(&result.col3, &mat.col3)
}

func M4MakeFromScalar(result *Matrix4, scalar float32) {
	V4MakeFromScalar(&result.col0, scalar)
	V4MakeFromScalar(&result.col1, scalar)
	V4MakeFromScalar(&result.col2, scalar)
	V4MakeFromScalar(&result.col3, scalar)
}

func M4MakeFromT3(result *Matrix4, mat *Transform3) {
	V4MakeFromV3Scalar(&result.col0, &mat.col0, 0.0)
	V4MakeFromV3Scalar(&result.col1, &mat.col1, 0.0)
	V4MakeFromV3Scalar(&result.col2, &mat.col2, 0.0)
	V4MakeFromV3Scalar(&result.col3, &mat.col3, 1.0)
}

func M4MakeFromCols(result *Matrix4, col0, col1, col2, col3 *Vector4) {
	V4Copy(&result.col0, col0)
	V4Copy(&result.col1, col1)
	V4Copy(&result.col2, col2)
	V4Copy(&result.col3, col3)
}

func M4MakeFromM3V3(result *Matrix4, mat *Matrix3, translateVec *Vector3) {
	V4MakeFromV3Scalar(&result.col0, &mat.col0, 0.0)
	V4MakeFromV3Scalar(&result.col1, &mat.col1, 0.0)
	V4MakeFromV3Scalar(&result.col2, &mat.col2, 0.0)
	V4MakeFromV3Scalar(&result.col3, translateVec, 1.0)
}

func M4MakeFromQV3(result *Matrix4, unitQuat *Quat, translateVec *Vector3) {
	var mat *Matrix3
	M3MakeFromQ(mat, unitQuat)
	V4MakeFromV3Scalar(&result.col0, &mat.col0, 0.0)
	V4MakeFromV3Scalar(&result.col1, &mat.col1, 0.0)
	V4MakeFromV3Scalar(&result.col2, &mat.col2, 0.0)
	V4MakeFromV3Scalar(&result.col3, translateVec, 1.0)
}

func (m *Matrix4) SetCol0(col0 *Vector4) {
	V4Copy(&m.col0, col0)
}

func (m *Matrix4) SetCol1(col1 *Vector4) {
	V4Copy(&m.col1, col1)
}

func (m *Matrix4) SetCol2(col2 *Vector4) {
	V4Copy(&m.col2, col2)
}

func (m *Matrix4) SetCol3(col3 *Vector4) {
	V4Copy(&m.col3, col3)
}

func (m *Matrix4) SetCol(col int, vec *Vector4) {
	switch col {
	case 0:
		V4Copy(&m.col0, vec)
	case 1:
		V4Copy(&m.col1, vec)
	case 2:
		V4Copy(&m.col2, vec)
	case 3:
		V4Copy(&m.col3, vec)
	}
}

func (m *Matrix4) SetRow(row int, vec *Vector4) {
	m.col0.SetElem(row, vec.x)
	m.col1.SetElem(row, vec.y)
	m.col2.SetElem(row, vec.z)
	m.col3.SetElem(row, vec.w)
}

func (m *Matrix4) SetElem(col, row int, val float32) {
	var tmpV3_0 Vector4
	M4GetCol(&tmpV3_0, m, col)
	tmpV3_0.SetElem(row, val)
	m.SetCol(col, &tmpV3_0)
}

func (m *Matrix4) GetElem(col, row int) float32 {
	var tmpV4_0 Vector4
	M4GetCol(&tmpV4_0, m, col)
	return tmpV4_0.GetElem(row)
}

func M4GetCol0(result *Vector4, mat *Matrix4) {
	V4Copy(result, &mat.col0)
}

func M4GetCol1(result *Vector4, mat *Matrix4) {
	V4Copy(result, &mat.col1)
}

func M4GetCol2(result *Vector4, mat *Matrix4) {
	V4Copy(result, &mat.col2)
}

func M4GetCol3(result *Vector4, mat *Matrix4) {
	V4Copy(result, &mat.col3)
}

func M4GetCol(result *Vector4, mat *Matrix4, col int) {
	switch col {
	case 0:
		V4Copy(result, &mat.col0)
	case 1:
		V4Copy(result, &mat.col1)
	case 2:
		V4Copy(result, &mat.col2)
	case 3:
		V4Copy(result, &mat.col3)
	}
}

func M4GetRow(result *Vector4, mat *Matrix4, row int) {
	V4MakeFromElems(result, mat.col0.GetElem(row), mat.col1.GetElem(row), mat.col2.GetElem(row), mat.col3.GetElem(row))
}

func M4Transpose(result, mat *Matrix4) {
	var tmpResult Matrix4
	V4MakeFromElems(&tmpResult.col0, mat.col0.x, mat.col1.x, mat.col2.x, mat.col3.x)
	V4MakeFromElems(&tmpResult.col1, mat.col0.y, mat.col1.y, mat.col2.y, mat.col3.y)
	V4MakeFromElems(&tmpResult.col2, mat.col0.z, mat.col1.z, mat.col2.z, mat.col3.z)
	V4MakeFromElems(&tmpResult.col3, mat.col0.w, mat.col1.w, mat.col2.w, mat.col3.w)
	M4Copy(result, &tmpResult)
}

func M4Inverse(result, mat *Matrix4) {
	var res0, res1, res2, res3 Vector4
	mA := mat.col0.x
	mB := mat.col0.y
	mC := mat.col0.z
	mD := mat.col0.w
	mE := mat.col1.x
	mF := mat.col1.y
	mG := mat.col1.z
	mH := mat.col1.w
	mI := mat.col2.x
	mJ := mat.col2.y
	mK := mat.col2.z
	mL := mat.col2.w
	mM := mat.col3.x
	mN := mat.col3.y
	mO := mat.col3.z
	mP := mat.col3.w
	tmp0 := ((mK * mD) - (mC * mL))
	tmp1 := ((mO * mH) - (mG * mP))
	tmp2 := ((mB * mK) - (mJ * mC))
	tmp3 := ((mF * mO) - (mN * mG))
	tmp4 := ((mJ * mD) - (mB * mL))
	tmp5 := ((mN * mH) - (mF * mP))
	res0.SetX(((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5))
	res0.SetY(((mN * tmp0) - (mP * tmp2)) - (mO * tmp4))
	res0.SetZ(((mD * tmp3) + (mC * tmp5)) - (mB * tmp1))
	res0.SetW(((mH * tmp2) + (mG * tmp4)) - (mF * tmp0))
	detInv := (1.0 / ((((mA * res0.x) + (mE * res0.y)) + (mI * res0.z)) + (mM * res0.w)))
	res1.SetX(mI * tmp1)
	res1.SetY(mM * tmp0)
	res1.SetZ(mA * tmp1)
	res1.SetW(mE * tmp0)
	res3.SetX(mI * tmp3)
	res3.SetY(mM * tmp2)
	res3.SetZ(mA * tmp3)
	res3.SetW(mE * tmp2)
	res2.SetX(mI * tmp5)
	res2.SetY(mM * tmp4)
	res2.SetZ(mA * tmp5)
	res2.SetW(mE * tmp4)
	tmp0 = ((mI * mB) - (mA * mJ))
	tmp1 = ((mM * mF) - (mE * mN))
	tmp2 = ((mI * mD) - (mA * mL))
	tmp3 = ((mM * mH) - (mE * mP))
	tmp4 = ((mI * mC) - (mA * mK))
	tmp5 = ((mM * mG) - (mE * mO))
	res2.SetX(((mL * tmp1) - (mJ * tmp3)) + res2.x)
	res2.SetY(((mP * tmp0) - (mN * tmp2)) + res2.y)
	res2.SetZ(((mB * tmp3) - (mD * tmp1)) - res2.z)
	res2.SetW(((mF * tmp2) - (mH * tmp0)) - res2.w)
	res3.SetX(((mJ * tmp5) - (mK * tmp1)) + res3.x)
	res3.SetY(((mN * tmp4) - (mO * tmp0)) + res3.y)
	res3.SetZ(((mC * tmp1) - (mB * tmp5)) - res3.z)
	res3.SetW(((mG * tmp0) - (mF * tmp4)) - res3.w)
	res1.SetX(((mK * tmp3) - (mL * tmp5)) - res1.x)
	res1.SetY(((mO * tmp2) - (mP * tmp4)) - res1.y)
	res1.SetZ(((mD * tmp5) - (mC * tmp3)) + res1.z)
	res1.SetW(((mH * tmp4) - (mG * tmp2)) + res1.w)
	V4ScalarMul(&result.col0, &res0, detInv)
	V4ScalarMul(&result.col1, &res1, detInv)
	V4ScalarMul(&result.col2, &res2, detInv)
	V4ScalarMul(&result.col3, &res3, detInv)
}

func M4AffineInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.col0)
	V4GetXYZ(&tmpV3_1, &mat.col1)
	V4GetXYZ(&tmpV3_2, &mat.col2)
	V4GetXYZ(&tmpV3_3, &mat.col3)
	affineMat.SetCol0(&tmpV3_0)
	affineMat.SetCol1(&tmpV3_1)
	affineMat.SetCol2(&tmpV3_2)
	affineMat.SetCol3(&tmpV3_3)
	T3Inverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func M4OrthoInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.col0)
	V4GetXYZ(&tmpV3_1, &mat.col1)
	V4GetXYZ(&tmpV3_2, &mat.col2)
	V4GetXYZ(&tmpV3_3, &mat.col3)
	affineMat.SetCol0(&tmpV3_0)
	affineMat.SetCol1(&tmpV3_1)
	affineMat.SetCol2(&tmpV3_2)
	affineMat.SetCol3(&tmpV3_3)
	T3OrthoInverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func (m *Matrix4) Determinant() float32 {
	mA := m.col0.x
	mB := m.col0.y
	mC := m.col0.z
	mD := m.col0.w
	mE := m.col1.x
	mF := m.col1.y
	mG := m.col1.z
	mH := m.col1.w
	mI := m.col2.x
	mJ := m.col2.y
	mK := m.col2.z
	mL := m.col2.w
	mM := m.col3.x
	mN := m.col3.y
	mO := m.col3.z
	mP := m.col3.w
	tmp0 := ((mK * mD) - (mC * mL))
	tmp1 := ((mO * mH) - (mG * mP))
	tmp2 := ((mB * mK) - (mJ * mC))
	tmp3 := ((mF * mO) - (mN * mG))
	tmp4 := ((mJ * mD) - (mB * mL))
	tmp5 := ((mN * mH) - (mF * mP))
	dx := (((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5))
	dy := (((mN * tmp0) - (mP * tmp2)) - (mO * tmp4))
	dz := (((mD * tmp3) + (mC * tmp5)) - (mB * tmp1))
	dw := (((mH * tmp2) + (mG * tmp4)) - (mF * tmp0))
	return ((((mA * dx) + (mE * dy)) + (mI * dz)) + (mM * dw))
}

func M4Add(result, mat0, mat1 *Matrix4) {
	V4Add(&result.col0, &mat0.col0, &mat1.col0)
	V4Add(&result.col1, &mat0.col1, &mat1.col1)
	V4Add(&result.col2, &mat0.col2, &mat1.col2)
	V4Add(&result.col3, &mat0.col3, &mat1.col3)
}

func M4Sub(result, mat0, mat1 *Matrix4) {
	V4Sub(&result.col0, &mat0.col0, &mat1.col0)
	V4Sub(&result.col1, &mat0.col1, &mat1.col1)
	V4Sub(&result.col2, &mat0.col2, &mat1.col2)
	V4Sub(&result.col3, &mat0.col3, &mat1.col3)
}

func M4Neg(result, mat *Matrix4) {
	V4Neg(&result.col0, &mat.col0)
	V4Neg(&result.col1, &mat.col1)
	V4Neg(&result.col2, &mat.col2)
	V4Neg(&result.col3, &mat.col3)
}

func M4AbsPerElem(result, mat *Matrix4) {
	V4AbsPerElem(&result.col0, &mat.col0)
	V4AbsPerElem(&result.col1, &mat.col1)
	V4AbsPerElem(&result.col2, &mat.col2)
	V4AbsPerElem(&result.col3, &mat.col3)
}

func M4ScalarMul(result, mat *Matrix4, scalar float32) {
	V4ScalarMul(&result.col0, &mat.col0, scalar)
	V4ScalarMul(&result.col1, &mat.col1, scalar)
	V4ScalarMul(&result.col2, &mat.col2, scalar)
	V4ScalarMul(&result.col3, &mat.col3, scalar)
}

func M4MulV4(result *Vector4, mat *Matrix4, vec *Vector4) {
	tmpX := (((mat.col0.x * vec.x) + (mat.col1.x * vec.y)) + (mat.col2.x * vec.z)) + (mat.col3.x * vec.w)
	tmpY := (((mat.col0.y * vec.x) + (mat.col1.y * vec.y)) + (mat.col2.y * vec.z)) + (mat.col3.y * vec.w)
	tmpZ := (((mat.col0.z * vec.x) + (mat.col1.z * vec.y)) + (mat.col2.z * vec.z)) + (mat.col3.z * vec.w)
	tmpW := (((mat.col0.w * vec.x) + (mat.col1.w * vec.y)) + (mat.col2.w * vec.z)) + (mat.col3.w * vec.w)
	V4MakeFromElems(result, tmpX, tmpY, tmpZ, tmpW)
}

func M4MulV3(result *Vector4, mat *Matrix4, vec *Vector3) {
	result.x = ((mat.col0.x * vec.x) + (mat.col1.x * vec.y)) + (mat.col2.x * vec.z)
	result.y = ((mat.col0.y * vec.x) + (mat.col1.y * vec.y)) + (mat.col2.y * vec.z)
	result.z = ((mat.col0.z * vec.x) + (mat.col1.z * vec.y)) + (mat.col2.z * vec.z)
	result.w = ((mat.col0.w * vec.x) + (mat.col1.w * vec.y)) + (mat.col2.w * vec.z)
}

func M4MulP3(result *Vector4, mat *Matrix4, pnt *Point3) {
	result.x = (((mat.col0.x * pnt.x) + (mat.col1.x * pnt.y)) + (mat.col2.x * pnt.z)) + mat.col3.x
	result.y = (((mat.col0.y * pnt.x) + (mat.col1.y * pnt.y)) + (mat.col2.y * pnt.z)) + mat.col3.y
	result.z = (((mat.col0.z * pnt.x) + (mat.col1.z * pnt.y)) + (mat.col2.z * pnt.z)) + mat.col3.z
	result.w = (((mat.col0.w * pnt.x) + (mat.col1.w * pnt.y)) + (mat.col2.w * pnt.z)) + mat.col3.w
}

func M4Mul(result, mat0, mat1 *Matrix4) {
	var tmpResult Matrix4
	M4MulV4(&tmpResult.col0, mat0, &mat1.col0)
	M4MulV4(&tmpResult.col1, mat0, &mat1.col1)
	M4MulV4(&tmpResult.col2, mat0, &mat1.col2)
	M4MulV4(&tmpResult.col3, mat0, &mat1.col3)
	M4Copy(result, &tmpResult)
}

func M4MulT3(result, mat *Matrix4, tfrm1 *Transform3) {
	var tmpResult Matrix4
	var tmpP3_0 Point3
	M4MulV3(&tmpResult.col0, mat, &tfrm1.col0)
	M4MulV3(&tmpResult.col1, mat, &tfrm1.col1)
	M4MulV3(&tmpResult.col2, mat, &tfrm1.col2)
	P3MakeFromV3(&tmpP3_0, &tfrm1.col3)
	M4MulP3(&tmpResult.col3, mat, &tmpP3_0)
	M4Copy(result, &tmpResult)
}

func M4MulPerElem(result, mat0, mat1 *Matrix4) {
	V4MulPerElem(&result.col0, &mat0.col0, &mat1.col0)
	V4MulPerElem(&result.col1, &mat0.col1, &mat1.col1)
	V4MulPerElem(&result.col2, &mat0.col2, &mat1.col2)
	V4MulPerElem(&result.col3, &mat0.col3, &mat1.col3)
}

func M4MakeIdentity(result *Matrix4) {
	V4MakeXAxis(&result.col0)
	V4MakeYAxis(&result.col1)
	V4MakeZAxis(&result.col2)
	V4MakeWAxis(&result.col3)
}

func (m *Matrix4) SetUpper3x3(mat3 *Matrix3) {
	m.col0.SetXYZ(&mat3.col0)
	m.col0.SetXYZ(&mat3.col1)
	m.col0.SetXYZ(&mat3.col2)
}

func M4GetUpper3x3(result *Matrix3, mat *Matrix4) {
	V4GetXYZ(&result.col0, &mat.col0)
	V4GetXYZ(&result.col1, &mat.col1)
	V4GetXYZ(&result.col2, &mat.col2)
}

func (m *Matrix4) SetTranslation(translateVec *Vector3) {
	m.col3.SetXYZ(translateVec)
}

func M4GetTranslation(result *Vector3, mat *Matrix4) {
	V4GetXYZ(result, &mat.col3)
}

func M4MakeRotationX(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeXAxis(&result.col0)
	V4MakeFromElems(&result.col1, 0.0, c, s, 0.0)
	V4MakeFromElems(&result.col2, 0.0, -s, c, 0.0)
	V4MakeWAxis(&result.col3)
}

func M4MakeRotationY(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeFromElems(&result.col0, c, 0.0, -s, 0.0)
	V4MakeYAxis(&result.col1)
	V4MakeFromElems(&result.col2, s, 0.0, c, 0.0)
	V4MakeWAxis(&result.col3)
}

func M4MakeRotationZ(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeFromElems(&result.col0, c, s, 0.0, 0.0)
	V4MakeFromElems(&result.col1, -s, c, 0.0, 0.0)
	V4MakeZAxis(&result.col2)
	V4MakeWAxis(&result.col3)
}

func M4MakeRotationZYX(result *Matrix4, radiansXYZ *Vector3) {
	sX := sin(radiansXYZ.x)
	cX := cos(radiansXYZ.x)
	sY := sin(radiansXYZ.y)
	cY := cos(radiansXYZ.y)
	sZ := sin(radiansXYZ.z)
	cZ := cos(radiansXYZ.z)
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)
	V4MakeFromElems(&result.col0, (cZ * cY), (sZ * cY), -sY, 0.0)
	V4MakeFromElems(&result.col1, ((tmp0 * sX) - (sZ * cX)), ((tmp1 * sX) + (cZ * cX)), (cY * sX), 0.0)
	V4MakeFromElems(&result.col2, ((tmp0 * cX) + (sZ * sX)), ((tmp1 * cX) - (cZ * sX)), (cY * cX), 0.0)
	V4MakeWAxis(&result.col3)
}

func M4MakeRotationAxis(result *Matrix4, radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	x := unitVec.x
	y := unitVec.y
	z := unitVec.z
	xy := x * y
	yz := y * z
	zx := z * x
	oneMinusC := 1.0 - c
	V4MakeFromElems(&result.col0, (((x * x) * oneMinusC) + c), ((xy * oneMinusC) + (z * s)), ((zx * oneMinusC) - (y * s)), 0.0)
	V4MakeFromElems(&result.col1, ((xy * oneMinusC) - (z * s)), (((y * y) * oneMinusC) + c), ((yz * oneMinusC) + (x * s)), 0.0)
	V4MakeFromElems(&result.col2, ((zx * oneMinusC) + (y * s)), ((yz * oneMinusC) - (x * s)), (((z * z) * oneMinusC) + c), 0.0)
	V4MakeWAxis(&result.col3)
}

func M4MakeRotationQ(result *Matrix4, unitQuat *Quat) {
	var tmpT3_0 Transform3
	T3MakeRotationQ(&tmpT3_0, unitQuat)
	M4MakeFromT3(result, &tmpT3_0)
}

func M4MakeScale(result *Matrix4, scaleVec *Vector3) {
	V4MakeFromElems(&result.col0, scaleVec.x, 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.col1, 0.0, scaleVec.y, 0.0, 0.0)
	V4MakeFromElems(&result.col2, 0.0, 0.0, scaleVec.z, 0.0)
	V4MakeWAxis(&result.col3)
}

func M4AppendScale(result, mat *Matrix4, scaleVec *Vector3) {
	V4ScalarMul(&result.col0, &mat.col0, scaleVec.X())
	V4ScalarMul(&result.col1, &mat.col1, scaleVec.Y())
	V4ScalarMul(&result.col2, &mat.col2, scaleVec.Z())
	V4Copy(&result.col3, &mat.col3)
}

func M4PrependScale(result *Matrix4, scaleVec *Vector3, mat *Matrix4) {
	var scale4 Vector4
	V4MakeFromV3Scalar(&scale4, scaleVec, 1.0)
	V4MulPerElem(&result.col0, &mat.col0, &scale4)
	V4MulPerElem(&result.col1, &mat.col1, &scale4)
	V4MulPerElem(&result.col2, &mat.col2, &scale4)
	V4MulPerElem(&result.col3, &mat.col3, &scale4)
}

func M4MakeTranslation(result *Matrix4, translateVec *Vector3) {
	V4MakeXAxis(&result.col0)
	V4MakeYAxis(&result.col1)
	V4MakeZAxis(&result.col2)
	V4MakeFromV3Scalar(&result.col3, translateVec, 1.0)
}

func M4MakeLookAt(result *Matrix4, eyePos, lookAtPos *Point3, upVec *Vector3) {
	var m4EyeFrame Matrix4
	var v3X, v3Y, v3Z, tmpV3_0, tmpV3_1 Vector3
	var tmpV4_0, tmpV4_1, tmpV4_2, tmpV4_3 Vector4
	V3Normalize(&v3Y, upVec)
	P3Sub(&tmpV3_0, eyePos, lookAtPos)
	V3Normalize(&v3Z, &tmpV3_0)
	V3Cross(&tmpV3_1, &v3Y, &v3Z)
	V3Normalize(&v3X, &tmpV3_1)
	V3Cross(&v3Y, &v3Z, &v3X)
	V4MakeFromV3(&tmpV4_0, &v3X)
	V4MakeFromV3(&tmpV4_1, &v3Y)
	V4MakeFromV3(&tmpV4_2, &v3Z)
	V4MakeFromP3(&tmpV4_3, eyePos)
	M4MakeFromCols(&m4EyeFrame, &tmpV4_0, &tmpV4_1, &tmpV4_2, &tmpV4_3)
	M4OrthoInverse(result, &m4EyeFrame)
}

func M4MakePerspective(result *Matrix4, fovyRadians, aspect, zNear, zFar float32) {
	f := tan(g_PI_OVER_2 - (0.5 * fovyRadians))
	rangeInv := 1.0 / (zNear - zFar)
	V4MakeFromElems(&result.col0, (f / aspect), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.col1, 0.0, f, 0.0, 0.0)
	V4MakeFromElems(&result.col2, 0.0, 0.0, ((zNear + zFar) * rangeInv), -1.0)
	V4MakeFromElems(&result.col3, 0.0, 0.0, (((zNear * zFar) * rangeInv) * 2.0), 0.0)
}

func M4MakeFrustum(result *Matrix4, left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))
	n2 := (zNear + zNear)
	V4MakeFromElems(&result.col0, (n2 * inv_rl), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.col1, 0.0, (n2 * inv_tb), 0.0, 0.0)
	V4MakeFromElems(&result.col2, (sum_rl * inv_rl), (sum_tb * inv_tb), (sum_nf * inv_nf), -1.0)
	V4MakeFromElems(&result.col3, 0.0, 0.0, ((n2 * inv_nf) * zFar), 0.0)
}

func M4MakeOrthographic(result *Matrix4, left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))
	V4MakeFromElems(&result.col0, (inv_rl + inv_rl), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.col1, 0.0, (inv_tb + inv_tb), 0.0, 0.0)
	V4MakeFromElems(&result.col2, 0.0, 0.0, (inv_nf + inv_nf), 0.0)
	V4MakeFromElems(&result.col3, (-sum_rl * inv_rl), (-sum_tb * inv_tb), (sum_nf * inv_nf), 1.0)
}

func M4Select(result, mat0, mat1 *Matrix4, select1 int) {
	V4Select(&result.col0, &mat0.col0, &mat1.col0, select1)
	V4Select(&result.col1, &mat0.col1, &mat1.col1, select1)
	V4Select(&result.col2, &mat0.col2, &mat1.col2, select1)
	V4Select(&result.col3, &mat0.col3, &mat1.col3, select1)
}

func (m *Matrix4) String() string {
	var tmp Matrix4
	M4Transpose(&tmp, m)
	return tmp.col0.String() + tmp.col1.String() + tmp.col2.String() + tmp.col3.String()
}

/*******/

func T3Copy(result, tfrm *Transform3) {
	V3Copy(&result.col0, &tfrm.col0)
	V3Copy(&result.col1, &tfrm.col1)
	V3Copy(&result.col2, &tfrm.col2)
	V3Copy(&result.col3, &tfrm.col3)
}

func T3MakeFromScalar(result *Transform3, scalar float32) {
	V3MakeFromScalar(&result.col0, scalar)
	V3MakeFromScalar(&result.col1, scalar)
	V3MakeFromScalar(&result.col2, scalar)
	V3MakeFromScalar(&result.col3, scalar)
}

func T3MakeFromCols(result *Transform3, col0, col1, col2, col3 *Vector3) {
	V3Copy(&result.col0, col0)
	V3Copy(&result.col1, col1)
	V3Copy(&result.col2, col2)
	V3Copy(&result.col3, col3)
}

func T3MakeFromM3V3(result *Transform3, tfrm *Matrix3, translateVec *Vector3) {
	result.SetUpper3x3(tfrm)
	result.SetTranslation(translateVec)
}

func T3MakeFromQV3(result *Transform3, unitQuat *Quat, translateVec *Vector3) {
	var tmpM3_0 Matrix3
	M3MakeFromQ(&tmpM3_0, unitQuat)
	result.SetUpper3x3(&tmpM3_0)
	result.SetTranslation(translateVec)
}

func (t *Transform3) SetCol0(col0 *Vector3) {
	V3Copy(&t.col0, col0)
}

func (t *Transform3) SetCol1(col1 *Vector3) {
	V3Copy(&t.col1, col1)
}

func (t *Transform3) SetCol2(col2 *Vector3) {
	V3Copy(&t.col2, col2)
}

func (t *Transform3) SetCol3(col3 *Vector3) {
	V3Copy(&t.col3, col3)
}

func (t *Transform3) SetCol(col int, vec *Vector3) {
	switch col {
	case 0:
		V3Copy(&t.col0, vec)
	case 1:
		V3Copy(&t.col1, vec)
	case 2:
		V3Copy(&t.col2, vec)
	case 3:
		V3Copy(&t.col3, vec)
	}
}

func (t *Transform3) SetRow(row int, vec *Vector4) {
	t.col0.SetElem(row, vec.GetElem(0))
	t.col1.SetElem(row, vec.GetElem(1))
	t.col2.SetElem(row, vec.GetElem(2))
	t.col3.SetElem(row, vec.GetElem(3))
}

func (t *Transform3) SetElem(col, row int, val float32) {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, t, col)
	tmpV3_0.SetElem(row, val)
	t.SetCol(col, &tmpV3_0)
}

func (t *Transform3) GetElem(col, row int) float32 {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, t, col)
	return tmpV3_0.GetElem(row)
}

func T3GetCol0(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.col0)
}

func T3GetCol1(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.col1)
}

func T3GetCol2(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.col2)
}

func T3GetCol3(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.col3)
}

func T3GetCol(result *Vector3, tfrm *Transform3, col int) {
	switch col {
	case 0:
		V3Copy(result, &tfrm.col0)
	case 1:
		V3Copy(result, &tfrm.col1)
	case 2:
		V3Copy(result, &tfrm.col2)
	case 3:
		V3Copy(result, &tfrm.col3)
	}
}

func T3GetRow(result *Vector4, tfrm *Transform3, row int) {
	V4MakeFromElems(result, tfrm.col0.GetElem(row), tfrm.col1.GetElem(row), tfrm.col2.GetElem(row), tfrm.col3.GetElem(row))
}

func T3Inverse(result, tfrm *Transform3) {
	var tmp0, tmp1, tmp2, inv0, inv1, inv2, tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3, tmpV3_4, tmpV3_5 Vector3
	V3Cross(&tmp0, &tfrm.col1, &tfrm.col2)
	V3Cross(&tmp1, &tfrm.col2, &tfrm.col0)
	V3Cross(&tmp2, &tfrm.col0, &tfrm.col1)
	detinv := (1.0 / V3Dot(&tfrm.col2, &tmp2))
	V3MakeFromElems(&inv0, (tmp0.x * detinv), (tmp1.x * detinv), (tmp2.x * detinv))
	V3MakeFromElems(&inv1, (tmp0.y * detinv), (tmp1.y * detinv), (tmp2.y * detinv))
	V3MakeFromElems(&inv2, (tmp0.z * detinv), (tmp1.z * detinv), (tmp2.z * detinv))
	V3Copy(&result.col0, &inv0)
	V3Copy(&result.col1, &inv1)
	V3Copy(&result.col2, &inv2)
	V3ScalarMul(&tmpV3_0, &inv0, tfrm.col3.x)
	V3ScalarMul(&tmpV3_1, &inv1, tfrm.col3.y)
	V3ScalarMul(&tmpV3_2, &inv2, tfrm.col3.z)
	V3Add(&tmpV3_3, &tmpV3_1, &tmpV3_2)
	V3Add(&tmpV3_4, &tmpV3_0, &tmpV3_3)
	V3Neg(&tmpV3_5, &tmpV3_4)
	V3Copy(&result.col3, &tmpV3_5)
}

func T3OrthoInverse(result, tfrm *Transform3) {
	var inv0, inv1, inv2, tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3, tmpV3_4, tmpV3_5 Vector3
	V3MakeFromElems(&inv0, tfrm.col0.x, tfrm.col1.x, tfrm.col2.x)
	V3MakeFromElems(&inv1, tfrm.col0.y, tfrm.col1.y, tfrm.col2.y)
	V3MakeFromElems(&inv2, tfrm.col0.z, tfrm.col1.z, tfrm.col2.z)
	V3Copy(&result.col0, &inv0)
	V3Copy(&result.col1, &inv1)
	V3Copy(&result.col2, &inv2)
	V3ScalarMul(&tmpV3_0, &inv0, tfrm.col3.x)
	V3ScalarMul(&tmpV3_1, &inv1, tfrm.col3.y)
	V3ScalarMul(&tmpV3_2, &inv2, tfrm.col3.z)
	V3Add(&tmpV3_3, &tmpV3_1, &tmpV3_2)
	V3Add(&tmpV3_4, &tmpV3_0, &tmpV3_3)
	V3Neg(&tmpV3_5, &tmpV3_4)
	V3Copy(&result.col3, &tmpV3_5)
}

func T3AbsPerElem(result, tfrm *Transform3) {
	V3AbsPerElem(&result.col0, &tfrm.col0)
	V3AbsPerElem(&result.col1, &tfrm.col1)
	V3AbsPerElem(&result.col2, &tfrm.col2)
	V3AbsPerElem(&result.col3, &tfrm.col3)
}

func T3MulV3(result *Vector3, tfrm *Transform3, vec *Vector3) {
	tmpX := ((tfrm.col0.x * vec.x) + (tfrm.col1.x * vec.y)) + (tfrm.col2.x * vec.z)
	tmpY := ((tfrm.col0.y * vec.x) + (tfrm.col1.y * vec.y)) + (tfrm.col2.y * vec.z)
	tmpZ := ((tfrm.col0.z * vec.x) + (tfrm.col1.z * vec.y)) + (tfrm.col2.z * vec.z)
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func T3MulP3(result *Point3, tfrm *Transform3, pnt *Point3) {
	tmpX := ((((tfrm.col0.x * pnt.x) + (tfrm.col1.x * pnt.y)) + (tfrm.col2.x * pnt.z)) + tfrm.col3.x)
	tmpY := ((((tfrm.col0.y * pnt.x) + (tfrm.col1.y * pnt.y)) + (tfrm.col2.y * pnt.z)) + tfrm.col3.y)
	tmpZ := ((((tfrm.col0.z * pnt.x) + (tfrm.col1.z * pnt.y)) + (tfrm.col2.z * pnt.z)) + tfrm.col3.z)
	P3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func T3Mul(result, tfrm0, tfrm1 *Transform3) {
	var tmpResult Transform3
	var tmpP3_0, tmpP3_1 Point3
	T3MulV3(&tmpResult.col0, tfrm0, &tfrm1.col0)
	T3MulV3(&tmpResult.col1, tfrm0, &tfrm1.col1)
	T3MulV3(&tmpResult.col2, tfrm0, &tfrm1.col2)
	P3MakeFromV3(&tmpP3_0, &tfrm1.col3)
	T3MulP3(&tmpP3_1, tfrm0, &tmpP3_0)
	V3MakeFromP3(&tmpResult.col3, &tmpP3_1)
	T3Copy(result, &tmpResult)
}

func T3MulPerElem(result, tfrm0, tfrm1 *Transform3) {
	V3MulPerElem(&result.col0, &tfrm0.col0, &tfrm1.col0)
	V3MulPerElem(&result.col1, &tfrm0.col1, &tfrm1.col1)
	V3MulPerElem(&result.col2, &tfrm0.col2, &tfrm1.col2)
	V3MulPerElem(&result.col3, &tfrm0.col3, &tfrm1.col3)
}

func T3MakeIdentity(result *Transform3) {
	V3MakeXAxis(&result.col0)
	V3MakeYAxis(&result.col1)
	V3MakeZAxis(&result.col2)
	V3MakeFromScalar(&result.col3, 0.0)
}

func (m *Transform3) SetUpper3x3(tfrm *Matrix3) {
	V3Copy(&m.col0, &tfrm.col0)
	V3Copy(&m.col1, &tfrm.col1)
	V3Copy(&m.col2, &tfrm.col2)
}

func T3GetUpper3x3(result *Matrix3, tfrm *Transform3) {
	M3MakeFromCols(result, &tfrm.col0, &tfrm.col1, &tfrm.col2)
}

func (t *Transform3) SetTranslation(translateVec *Vector3) {
	V3Copy(&t.col3, translateVec)
}

func T3GetTranslation(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.col3)
}

func T3MakeRotationX(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeXAxis(&result.col0)
	V3MakeFromElems(&result.col1, 0.0, c, s)
	V3MakeFromElems(&result.col2, 0.0, -s, c)
	V3MakeFromScalar(&result.col3, 0.0)
}

func T3MakeRotationY(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.col0, c, 0.0, -s)
	V3MakeYAxis(&result.col1)
	V3MakeFromElems(&result.col2, s, 0.0, c)
	V3MakeFromScalar(&result.col3, 0.0)
}

func T3MakeRotationZ(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.col0, c, s, 0.0)
	V3MakeFromElems(&result.col1, -s, c, 0.0)
	V3MakeZAxis(&result.col2)
	V3MakeFromScalar(&result.col3, 0.0)
}

func T3MakeRotationZYX(result *Transform3, radiansXYZ *Vector3) {
	sX := sin(radiansXYZ.x)
	cX := cos(radiansXYZ.x)
	sY := sin(radiansXYZ.y)
	cY := cos(radiansXYZ.y)
	sZ := sin(radiansXYZ.z)
	cZ := cos(radiansXYZ.z)
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)
	V3MakeFromElems(&result.col0, (cZ * cY), (sZ * cY), -sY)
	V3MakeFromElems(&result.col1, ((tmp0 * sX) - (sZ * cX)), ((tmp1 * sX) + (cZ * cX)), (cY * sX))
	V3MakeFromElems(&result.col2, ((tmp0 * cX) + (sZ * sX)), ((tmp1 * cX) - (cZ * sX)), (cY * cX))
	V3MakeFromScalar(&result.col3, 0.0)
}

func T3MakeRotationAxis(result *Transform3, radians float32, unitVec *Vector3) {
	var tmpM3_0 Matrix3
	var tmpV3_0 Vector3
	M3MakeRotationAxis(&tmpM3_0, radians, unitVec)
	V3MakeFromScalar(&tmpV3_0, 0.0)
	T3MakeFromM3V3(result, &tmpM3_0, &tmpV3_0)
}

func T3MakeRotationQ(result *Transform3, unitQuat *Quat) {
	var tmpM3_0 Matrix3
	var tmpV3_0 Vector3
	M3MakeFromQ(&tmpM3_0, unitQuat)
	V3MakeFromScalar(&tmpV3_0, 0.0)
	T3MakeFromM3V3(result, &tmpM3_0, &tmpV3_0)
}

func T3MakeScale(result *Transform3, scaleVec *Vector3) {
	V3MakeFromElems(&result.col0, scaleVec.x, 0.0, 0.0)
	V3MakeFromElems(&result.col1, 0.0, scaleVec.y, 0.0)
	V3MakeFromElems(&result.col2, 0.0, 0.0, scaleVec.z)
	V3MakeFromScalar(&result.col3, 0.0)
}

func T3AppendScale(result, tfrm *Transform3, scaleVec *Vector3) {
	V3ScalarMul(&result.col0, &tfrm.col0, scaleVec.X())
	V3ScalarMul(&result.col1, &tfrm.col1, scaleVec.Y())
	V3ScalarMul(&result.col2, &tfrm.col2, scaleVec.Z())
	V3Copy(&result.col3, &tfrm.col3)
}

func T3PrependScale(result *Transform3, scaleVec *Vector3, tfrm *Transform3) {
	V3MulPerElem(&result.col0, &tfrm.col0, scaleVec)
	V3MulPerElem(&result.col1, &tfrm.col1, scaleVec)
	V3MulPerElem(&result.col2, &tfrm.col2, scaleVec)
	V3MulPerElem(&result.col3, &tfrm.col3, scaleVec)
}

func T3MakeTranslation(result *Transform3, translateVec *Vector3) {
	V3MakeXAxis(&result.col0)
	V3MakeYAxis(&result.col1)
	V3MakeZAxis(&result.col2)
	V3Copy(&result.col3, translateVec)
}

func T3Select(result, tfrm0, tfrm1 *Transform3, select1 int) {
	V3Select(&result.col0, &tfrm0.col0, &tfrm1.col0, select1)
	V3Select(&result.col1, &tfrm0.col1, &tfrm1.col1, select1)
	V3Select(&result.col2, &tfrm0.col2, &tfrm1.col2, select1)
	V3Select(&result.col3, &tfrm0.col3, &tfrm1.col3, select1)
}

func (t *Transform3) String() string {
	var tmpV4_0, tmpV4_1, tmpV4_2 Vector4
	T3GetRow(&tmpV4_0, t, 0)
	T3GetRow(&tmpV4_1, t, 1)
	T3GetRow(&tmpV4_2, t, 2)
	return tmpV4_0.String() + tmpV4_1.String() + tmpV4_2.String()
}

/*******/

func QMakeFromM3(result *Quat, tfrm *Matrix3) {
	xx := tfrm.col0.x
	yx := tfrm.col0.y
	zx := tfrm.col0.z
	xy := tfrm.col1.x
	yy := tfrm.col1.y
	zy := tfrm.col1.z
	xz := tfrm.col2.x
	yz := tfrm.col2.y
	zz := tfrm.col2.z

	trace := ((xx + yy) + zz)

	negTrace := (trace < 0.0)
	ZgtX := zz > xx
	ZgtY := zz > yy
	YgtX := yy > xx
	largestXorY := (!ZgtX || !ZgtY) && negTrace
	largestYorZ := (YgtX || ZgtX) && negTrace
	largestZorX := (ZgtY || !YgtX) && negTrace

	if largestXorY {
		zz = -zz
		xy = -xy
	}
	if largestYorZ {
		xx = -xx
		yz = -yz
	}
	if largestZorX {
		yy = -yy
		zx = -zx
	}

	radicand := (((xx + yy) + zz) + 1.0)
	scale := (0.5 * (1.0 / sqrt(radicand)))

	tmpx := ((zy - yz) * scale)
	tmpy := ((xz - zx) * scale)
	tmpz := ((yx - xy) * scale)
	tmpw := (radicand * scale)
	qx := tmpx
	qy := tmpy
	qz := tmpz
	qw := tmpw

	if largestXorY {
		qx = tmpw
		qy = tmpz
		qz = tmpy
		qw = tmpx
	}
	if largestYorZ {
		tmpx = qx
		tmpz = qz
		qx = qy
		qy = tmpx
		qz = qw
		qw = tmpz
	}

	result.x = qx
	result.y = qy
	result.z = qz
	result.w = qw
}

func V3Outer(result *Matrix3, tfrm0, tfrm1 *Vector3) {
	V3ScalarMul(&result.col0, tfrm0, tfrm1.X())
	V3ScalarMul(&result.col1, tfrm0, tfrm1.Y())
	V3ScalarMul(&result.col2, tfrm0, tfrm1.Z())
}

func V4Outer(result *Matrix4, tfrm0, tfrm1 *Vector4) {
	V4ScalarMul(&result.col0, tfrm0, tfrm1.X())
	V4ScalarMul(&result.col1, tfrm0, tfrm1.Y())
	V4ScalarMul(&result.col2, tfrm0, tfrm1.Z())
	V4ScalarMul(&result.col3, tfrm0, tfrm1.W())
}

func V3RowMul(result *Vector3, vec *Vector3, mat *Matrix3) {
	tmpX := (((vec.x * mat.col0.x) + (vec.y * mat.col0.y)) + (vec.z * mat.col0.z))
	tmpY := (((vec.x * mat.col1.x) + (vec.y * mat.col1.y)) + (vec.z * mat.col1.z))
	tmpZ := (((vec.x * mat.col2.x) + (vec.y * mat.col2.y)) + (vec.z * mat.col2.z))
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func V3CrossMatrix(result *Matrix3, vec *Vector3) {
	V3MakeFromElems(&result.col0, 0.0, vec.z, -vec.y)
	V3MakeFromElems(&result.col1, -vec.z, 0.0, vec.x)
	V3MakeFromElems(&result.col2, vec.y, -vec.x, 0.0)
}

func V3CrossMatrixMul(result *Matrix3, vec *Vector3, mat *Matrix3) {
	var tmpV3_0, tmpV3_1, tmpV3_2 Vector3
	V3Cross(&tmpV3_0, vec, &mat.col0)
	V3Cross(&tmpV3_1, vec, &mat.col1)
	V3Cross(&tmpV3_2, vec, &mat.col2)
	M3MakeFromCols(result, &tmpV3_0, &tmpV3_1, &tmpV3_2)
}
