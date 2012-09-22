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

func M3SetCol0(result *Matrix3, col0 *Vector3) {
	V3Copy(&result.col0, col0)
}

func M3SetCol1(result *Matrix3, col1 *Vector3) {
	V3Copy(&result.col1, col1)
}

func M3SetCol2(result *Matrix3, col2 *Vector3) {
	V3Copy(&result.col2, col2)
}

func M3SetCol(result *Matrix3, col int, vec *Vector3) {
	switch col {
	case 0:
		V3Copy(&result.col0, vec)
	case 1:
		V3Copy(&result.col1, vec)
	case 2:
		V3Copy(&result.col2, vec)
	}
}

func M3SetRow(result *Matrix3, row int, vec *Vector3) {
	V3SetElem(&result.col0, row, V3GetElem(vec, 0))
	V3SetElem(&result.col1, row, V3GetElem(vec, 1))
	V3SetElem(&result.col2, row, V3GetElem(vec, 2))
}

func M3SetElem(result *Matrix3, col, row int, val float32) {
	var tmpV3_0 Vector3
	M3GetCol(&tmpV3_0, result, col)
	V3SetElem(&tmpV3_0, row, val)
	M3SetCol(result, col, &tmpV3_0)
}

func M3GetElem(mat *Matrix3, col, row int) float32 {
	var tmpV3_0 Vector3
	M3GetCol(&tmpV3_0, mat, col)
	return V3GetElem(&tmpV3_0, row)
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
		V3Copy(result, &mat.col0)
	case 2:
		V3Copy(result, &mat.col0)
	}
}

func M3GetRow(result *Vector3, mat *Matrix3, row int) {
	x := V3GetElem(&mat.col0, row)
	y := V3GetElem(&mat.col1, row)
	z := V3GetElem(&mat.col2, row)
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

func M3Determinant(mat *Matrix3) {
	var tmpV3_0 Vector3
	V3Cross(&tmpV3_0, &mat.col0, &mat.col1)
	V3Dot(&mat.col2, &tmpV3_0)
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

func (mat *Matrix3) String() string {
	var tmp Matrix3
	M3Transpose(&tmp, mat)
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

func M4SetCol0(result *Matrix4, col0 *Vector4) {
	V4Copy(&result.col0, col0)
}

func M4SetCol1(result *Matrix4, col1 *Vector4) {
	V4Copy(&result.col1, col1)
}

func M4SetCol2(result *Matrix4, col2 *Vector4) {
	V4Copy(&result.col2, col2)
}

func M4SetCol3(result *Matrix4, col3 *Vector4) {
	V4Copy(&result.col3, col3)
}

func M4SetCol(result *Matrix4, col int, vec *Vector4) {
	switch col {
	case 0:
		V4Copy(&result.col0, vec)
	case 1:
		V4Copy(&result.col1, vec)
	case 2:
		V4Copy(&result.col2, vec)
	case 3:
		V4Copy(&result.col3, vec)
	}
}

func M4SetRow(result *Matrix4, row int, vec *Vector4) {
	V4SetElem(&result.col0, row, vec.x)
	V4SetElem(&result.col1, row, vec.y)
	V4SetElem(&result.col2, row, vec.z)
	V4SetElem(&result.col3, row, vec.w)
}

func M4SetElem(result *Matrix4, col, row int, val float32) {
	var tmpV3_0 Vector4
	M4GetCol(&tmpV3_0, result, col)
	V4SetElem(&tmpV3_0, row, val)
	M4SetCol(result, col, &tmpV3_0)
}

func M4GetElem(mat *Matrix4, col, row int) float32 {
	var tmpV4_0 Vector4
	M4GetCol(&tmpV4_0, mat, col)
	return V4GetElem(&tmpV4_0, row)
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
	V4MakeFromElems(result, V4GetElem(&mat.col0, row), V4GetElem(&mat.col1, row), V4GetElem(&mat.col2, row), V4GetElem(&mat.col3, row))
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
	V4SetX(&res0, (((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5)))
	V4SetY(&res0, (((mN * tmp0) - (mP * tmp2)) - (mO * tmp4)))
	V4SetZ(&res0, (((mD * tmp3) + (mC * tmp5)) - (mB * tmp1)))
	V4SetW(&res0, (((mH * tmp2) + (mG * tmp4)) - (mF * tmp0)))
	detInv := (1.0 / ((((mA * res0.x) + (mE * res0.y)) + (mI * res0.z)) + (mM * res0.w)))
	V4SetX(&res1, (mI * tmp1))
	V4SetY(&res1, (mM * tmp0))
	V4SetZ(&res1, (mA * tmp1))
	V4SetW(&res1, (mE * tmp0))
	V4SetX(&res3, (mI * tmp3))
	V4SetY(&res3, (mM * tmp2))
	V4SetZ(&res3, (mA * tmp3))
	V4SetW(&res3, (mE * tmp2))
	V4SetX(&res2, (mI * tmp5))
	V4SetY(&res2, (mM * tmp4))
	V4SetZ(&res2, (mA * tmp5))
	V4SetW(&res2, (mE * tmp4))
	tmp0 = ((mI * mB) - (mA * mJ))
	tmp1 = ((mM * mF) - (mE * mN))
	tmp2 = ((mI * mD) - (mA * mL))
	tmp3 = ((mM * mH) - (mE * mP))
	tmp4 = ((mI * mC) - (mA * mK))
	tmp5 = ((mM * mG) - (mE * mO))
	V4SetX(&res2, (((mL * tmp1) - (mJ * tmp3)) + res2.x))
	V4SetY(&res2, (((mP * tmp0) - (mN * tmp2)) + res2.y))
	V4SetZ(&res2, (((mB * tmp3) - (mD * tmp1)) - res2.z))
	V4SetW(&res2, (((mF * tmp2) - (mH * tmp0)) - res2.w))
	V4SetX(&res3, (((mJ * tmp5) - (mK * tmp1)) + res3.x))
	V4SetY(&res3, (((mN * tmp4) - (mO * tmp0)) + res3.y))
	V4SetZ(&res3, (((mC * tmp1) - (mB * tmp5)) - res3.z))
	V4SetW(&res3, (((mG * tmp0) - (mF * tmp4)) - res3.w))
	V4SetX(&res1, (((mK * tmp3) - (mL * tmp5)) - res1.x))
	V4SetY(&res1, (((mO * tmp2) - (mP * tmp4)) - res1.y))
	V4SetZ(&res1, (((mD * tmp5) - (mC * tmp3)) + res1.z))
	V4SetW(&res1, (((mH * tmp4) - (mG * tmp2)) + res1.w))
	V4ScalarMul(&result.col0, &res0, detInv)
	V4ScalarMul(&result.col1, &res1, detInv)
	V4ScalarMul(&result.col2, &res2, detInv)
	V4ScalarMul(&result.col3, &res3, detInv)
}

func M4AffineInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.col0)
	T3SetCol0(&affineMat, &tmpV3_0)
	V4GetXYZ(&tmpV3_1, &mat.col1)
	T3SetCol1(&affineMat, &tmpV3_1)
	V4GetXYZ(&tmpV3_2, &mat.col2)
	T3SetCol2(&affineMat, &tmpV3_2)
	V4GetXYZ(&tmpV3_3, &mat.col3)
	T3SetCol3(&affineMat, &tmpV3_3)
	T3Inverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func M4OrthoInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.col0)
	T3SetCol0(&affineMat, &tmpV3_0)
	V4GetXYZ(&tmpV3_1, &mat.col1)
	T3SetCol1(&affineMat, &tmpV3_1)
	V4GetXYZ(&tmpV3_2, &mat.col2)
	T3SetCol2(&affineMat, &tmpV3_2)
	V4GetXYZ(&tmpV3_3, &mat.col3)
	T3SetCol3(&affineMat, &tmpV3_3)
	T3OrthoInverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func M4Determinant(mat *Matrix4) float32 {
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

func M4SetUpper3x3(result *Matrix4, mat3 *Matrix3) {
	V4SetXYZ(&result.col0, &mat3.col0)
	V4SetXYZ(&result.col1, &mat3.col1)
	V4SetXYZ(&result.col2, &mat3.col2)
}

func M4GetUpper3x3(result *Matrix3, mat *Matrix4) {
	V4GetXYZ(&result.col0, &mat.col0)
	V4GetXYZ(&result.col1, &mat.col1)
	V4GetXYZ(&result.col2, &mat.col2)
}

func M4SetTranslation(result *Matrix4, translateVec *Vector3) {
	V4SetXYZ(&result.col3, translateVec)
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
	V4ScalarMul(&result.col0, &mat.col0, V3GetX(scaleVec))
	V4ScalarMul(&result.col1, &mat.col1, V3GetY(scaleVec))
	V4ScalarMul(&result.col2, &mat.col2, V3GetZ(scaleVec))
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

func (mat *Matrix4) String() string {
	var tmp Matrix4
	M4Transpose(&tmp, mat)
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
	T3SetUpper3x3(result, tfrm)
	T3SetTranslation(result, translateVec)
}

func T3MakeFromQV3(result *Transform3, unitQuat *Quat, translateVec *Vector3) {
	var tmpM3_0 Matrix3
	M3MakeFromQ(&tmpM3_0, unitQuat)
	T3SetUpper3x3(result, &tmpM3_0)
	T3SetTranslation(result, translateVec)
}

func T3SetCol0(result *Transform3, col0 *Vector3) {
	V3Copy(&result.col0, col0)
}

func T3SetCol1(result *Transform3, col1 *Vector3) {
	V3Copy(&result.col1, col1)
}

func T3SetCol2(result *Transform3, col2 *Vector3) {
	V3Copy(&result.col2, col2)
}

func T3SetCol3(result *Transform3, col3 *Vector3) {
	V3Copy(&result.col3, col3)
}

func T3SetCol(result *Transform3, col int, vec *Vector3) {
	switch col {
	case 0:
		V3Copy(&result.col0, vec)
	case 1:
		V3Copy(&result.col1, vec)
	case 2:
		V3Copy(&result.col2, vec)
	case 3:
		V3Copy(&result.col3, vec)
	}
}

func T3SetRow(result *Transform3, row int, vec *Vector4) {
	V3SetElem(&result.col0, row, V4GetElem(vec, 0))
	V3SetElem(&result.col1, row, V4GetElem(vec, 1))
	V3SetElem(&result.col2, row, V4GetElem(vec, 2))
	V3SetElem(&result.col3, row, V4GetElem(vec, 3))
}

func T3SetElem(result *Transform3, col, row int, val float32) {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, result, col)
	V3SetElem(&tmpV3_0, row, val)
	T3SetCol(result, col, &tmpV3_0)
}

func T3GetElem(tfrm *Transform3, col, row int) float32 {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, tfrm, col)
	return V3GetElem(&tmpV3_0, row)
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
	V4MakeFromElems(result, V3GetElem(&tfrm.col0, row), V3GetElem(&tfrm.col1, row), V3GetElem(&tfrm.col2, row), V3GetElem(&tfrm.col3, row))
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

func T3SetUpper3x3(result *Transform3, tfrm *Matrix3) {
	V3Copy(&result.col0, &tfrm.col0)
	V3Copy(&result.col1, &tfrm.col1)
	V3Copy(&result.col2, &tfrm.col2)
}

func T3GetUpper3x3(result *Matrix3, tfrm *Transform3) {
	M3MakeFromCols(result, &tfrm.col0, &tfrm.col1, &tfrm.col2)
}

func T3SetTranslation(result *Transform3, translateVec *Vector3) {
	V3Copy(&result.col3, translateVec)
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
	V3ScalarMul(&result.col0, &tfrm.col0, V3GetX(scaleVec))
	V3ScalarMul(&result.col1, &tfrm.col1, V3GetY(scaleVec))
	V3ScalarMul(&result.col2, &tfrm.col2, V3GetZ(scaleVec))
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

func (tfrm *Transform3) String() string {
	var tmpV4_0, tmpV4_1, tmpV4_2 Vector4
	T3GetRow(&tmpV4_0, tfrm, 0)
	T3GetRow(&tmpV4_1, tfrm, 1)
	T3GetRow(&tmpV4_2, tfrm, 2)
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
	V3ScalarMul(&result.col0, tfrm0, V3GetX(tfrm1))
	V3ScalarMul(&result.col1, tfrm0, V3GetY(tfrm1))
	V3ScalarMul(&result.col2, tfrm0, V3GetZ(tfrm1))
}

func V4Outer(result *Matrix4, tfrm0, tfrm1 *Vector4) {
	V4ScalarMul(&result.col0, tfrm0, V4GetX(tfrm1))
	V4ScalarMul(&result.col1, tfrm0, V4GetY(tfrm1))
	V4ScalarMul(&result.col2, tfrm0, V4GetZ(tfrm1))
	V4ScalarMul(&result.col3, tfrm0, V4GetW(tfrm1))
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
