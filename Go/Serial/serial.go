package main

import (
	"fmt"
	"math"
	"formular"
	"os"
	"time"
)

//BROJ KLATANA
var pendulum_number = 30


var length1 float64 = 1
var length2 float64= 1
var mass1 float64= 1
var mass2 float64= 1

var angle_velocity float64= 0
var angle_acceleration float64 = 0

//opcije:
var fps = 30
var g float64 = 9.81
var max_time = 60
var make_for_30fps float64 = 0.03333333
var step float64= 0.001
var differ = 0.03

var pendulum_angles1 []float64
var pendulum_angles2 []float64

var pendulum_velocities1 []float64
var pendulum_velocities2 []float64

var pendulum_accelerations1 []float64
var pendulum_accelerations2 []float64

var pendulum_file_names []string

var pendulums []Pendulum


//var pendulum_scatter1 []float64
//var pendulum_scatter2 []float64

var wisteph int = 1928
var height int = 1020
var x_offset = int(wisteph/2)
var y_offset = int(height/4)
	
type Touple struct{
	x1 float64
	x2 float64
}

type Quad struct{
	ang1 float64
	ang2 float64
	ang_v1 float64
	ang_v2 float64
}

type Pendulum struct {
	index int
	p_values []Quad
}

func main(){	

	start := time.Now()

	for i := 0; i < pendulum_number; i++ {
		pendulum_angles1 = append(pendulum_angles1, float64(math.Pi) / (float64(2.25) + (float64(i)* float64(0.05))))
		pendulum_angles2 = append(pendulum_angles2, float64(math.Pi) / (float64(2.25) + (float64(i)* float64(0.05))))
		
		pendulum_velocities1 = append(pendulum_velocities1, angle_velocity)
		pendulum_velocities2 = append(pendulum_velocities2, angle_velocity)

		pendulum_accelerations1 = append(pendulum_accelerations1, angle_acceleration)
		pendulum_accelerations2 = append(pendulum_accelerations2, angle_acceleration)

		pendulum_file_names = append(pendulum_file_names, fmt.Sprintf("pendulum%d.txt", i+1)) 
		
		var pendulum_temp = Pendulum{}
		pendulum_temp.index = i
		pendulums = append(pendulums, pendulum_temp)
	}
	
	
	for i:= 0; i < pendulum_number; i++ {
	
		p := Quad{pendulum_angles1[i], pendulum_angles2[i], pendulum_velocities1[i], pendulum_velocities2[i] }
		var t float64 = 0
	
		for t < float64(max_time) {
			t = t + step
		
			var alpha float64 = p.ang1 + step * p.ang_v1
			var beta float64 = p.ang2 + step * p.ang_v2
			
			var gamma float64= p.ang_v1 + step * formular.FirstAcceleration(p.ang1, p.ang2, mass1, mass2, length1, length2, g, p.ang_v1, p.ang_v2)
			var delta float64= p.ang_v2 + step * formular.SecondAcceleration(p.ang1, p.ang2, mass1, mass2, length1, length2, g, p.ang_v1, p.ang_v2)
		
			
			var new_angle1 float64= (p.ang1+step*p.ang_v1)
			var new_angle2 float64= (p.ang2+step*p.ang_v2)
			
			var new_angle_velocity1 float64= (p.ang_v1+(step/2)*(formular.FirstAcceleration(alpha, beta, mass1, mass2, length1, length2, g, gamma, delta) +
                               formular.FirstAcceleration(p.ang1, p.ang2, mass1, mass2, length1, length2, g, p.ang_v1, p.ang_v2)))
							   
			var new_angle_velocity2 float64= (p.ang_v2+(step/2)*(formular.SecondAcceleration(alpha, beta, mass1, mass2, length1, length2, g, gamma, delta) +
                               formular.SecondAcceleration(p.ang1, p.ang2, mass1, mass2, length1, length2, g, p.ang_v1, p.ang_v2)))
				
			
			p_new :=  Quad{new_angle1, new_angle2, new_angle_velocity1, new_angle_velocity2}
			
			pendulums[i].p_values = append(pendulums[i].p_values, p_new) 
			
			p = p_new
							   
		}
	}
	
	var counter int = 0
	for i:= 0; i < pendulum_number; i++ {
		
		f, _ := os.Create(fmt.Sprintf("pendulum%d.txt", i+1))
		defer f.Close()

		for j:=0; j < len(pendulums[i].p_values); j++{
			if counter%33 == 0{
				
				var x1_value float64= length1 * float64(250) * math.Sin(pendulums[i].p_values[j].ang1) + float64(x_offset)
				var y1_value float64= length1 * float64(250) * math.Cos(pendulums[i].p_values[j].ang1) + float64(y_offset)
				var x2_value float64= length2 * float64(250) * math.Sin(pendulums[i].p_values[j].ang2) + float64(x1_value)
				var y2_value float64= length2 * float64(250) * math.Cos(pendulums[i].p_values[j].ang2) + float64(y1_value)
				
				f.WriteString(fmt.Sprintf("(%f,%f)|(%f,%f)\n", x1_value, y1_value,x2_value,y2_value))
				
				counter = 0
			}
			counter = counter + 1
			
		}
	}
	
	duration := time.Since(start)
	fmt.Print("Program finished in ")
	fmt.Print(duration)
}