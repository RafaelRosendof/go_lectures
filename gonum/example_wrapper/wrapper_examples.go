package examplewrapper

import (
	"fmt"
	"math"
	"math_test/dif"
)

func Derivate_examples() {

	fmt.Println("Tesinting all derivates")

	f := func(x float64) float64 {
		return x * (math.Cos(x*x*x) + math.Log10(x*123*x) - math.Pow(x, math.Log2(x*x*x)))
	}

	f_multi := func(x []float64) float64 {
		x0, x1, x2, x3 := x[0], x[1], x[2], x[3]

		term1 := math.Exp(-(x0*x0 + x1*x1)) * math.Cos(x0*x1*x2)

		innerLog := math.Abs(x2*123*x3) + 1e-9
		term2 := math.Log10(innerLog) * math.Pow(math.Abs(x0), math.Log2(math.Abs(x1*x2*x3)+1e-9))

		term3 := math.Sin(math.Pow(x3, 15)) / (1 + x0*x0 + x1*x1 + x2*x2)

		return term1 + term2 - term3
	}

	fmt.Println("Printing the derivate of the function F as the basic derivation")
	dif.Example_derivative(f, math.Pi/2)

	fmt.Println("Printing the gradient of the function F in all variables")

	point := []float64{1.3, 0.2, 2.9, 0.9}

	grad := dif.Example_gradient(f_multi, point)

	fmt.Println("Printing the gradient result", grad)

	fmt.Println("Printing the hessian of the function F in all variables")

	hess := dif.Example_hessian(f_multi, point)

	fmt.Println("Printing the hessian result", hess)

	fmt.Println("Printing the jacobian of the function F in all variables")

	f_vec := func(dst, x []float64) {
		x0, x1, x2, x3 := x[0], x[1], x[2], x[3]

		dst[0] = math.Sin(x0*x1)*math.Exp(-x2) + math.Cos(x3)
		dst[1] = math.Log1p(math.Abs(x0)) * math.Pow(math.Abs(x1), math.Sin(x2))
		dst[2] = (math.Sqrt(math.Abs(x0*x3)) + 1) / (1 + x1*x1 + x2*x2)
		dst[3] = x0 * (math.Cos(x1*x2*x3) + math.Log10(math.Abs(x0*123)+1) - math.Pow(math.Abs(x1), 0.5))

	}

	point_vec := []float64{1.0, 2.0, 0.5, 1.5}
	res := make([]float64, 4)
	f_vec(res, point_vec)

	res_jacob := dif.Example_jacobian(f_vec, point_vec, 4)

	fmt.Println("Printing the jacobian result", res_jacob)

	fmt.Println("Printing the laplacian of the function F")

	laplacian := dif.Example_laplacian(f_multi, point)

	fmt.Println("Printing the laplacian result", laplacian)

}

func integration_example() {

}

func interpolation_example() {

}
