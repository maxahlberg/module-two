package main

import("fmt")

func GenDisplaceFn(a,vel,disp float64) func(float64) float64{
	fn := func(t float64) float64{
		s := 0.5 * a * t * t + vel * t + disp
		return s
	}
	return fn
}
func main(){
	var acc,ini_vel,ini_dis float64

	fmt.Println("Enter accelaration, initial velocity , initial displacement")

	fmt.Scan(&acc,&ini_vel,&ini_dis)

	fn := GenDisplaceFn(acc,ini_vel,ini_dis)

	fmt.Println("Enter t : ")
	var t float64
	fmt.Scanln(&t)

	fmt.Println("Displace after time ",t,"is",fn(t))
}
