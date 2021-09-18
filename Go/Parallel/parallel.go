package main

import (
	"fmt"
	"math"
	"formular"
	"os"
	"time"
    "sync"
)

//BROJ KLATANA
var pendulum_number = 100


var length1 float64 = 1
var length2 float64= 1
var mass1 float64= 1
var mass2 float64= 1

var angle_velocity float64= 0
var angle_acceleration float64 = 0

//opcije:
var fps = 100
var g float64 = 9.81
var max_time = 60
var make_for_30fps float64 = 0.03333333
var step float64= 0.001
var differ = 0.03


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
    
    angle1 float64
    angle2 float64
    
    velocity1 float64
    velocity2 float64
    
    acceleration1 float64
    acceleration2 float64
    
    file_name string
    
	p_values []Quad
}

type SafePendulum struct{
    pendulums []Pendulum
    mux sync.Mutex
}

func(c *SafePendulum) calculate_p_values(wg *sync.WaitGroup, index int){
    defer wg.Done()
    c.mux.Lock()
    defer c.mux.Unlock()
    
    var pendulum_temp Pendulum = c.pendulums[index]
    
    p := Quad{pendulum_temp.angle1, pendulum_temp.angle2, pendulum_temp.velocity1, pendulum_temp.velocity1}
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
			

        c.pendulums[index].p_values = append(c.pendulums[index].p_values, p_new)
		//pendulum_temp.p_values = append(pendulum_temp.p_values, p_new) 
			
		p = p_new
							   
	} 
}

func(c *SafePendulum) calculate_x_y_positions(wg *sync.WaitGroup, index int){
    defer wg.Done()
    f, _ := os.Create(fmt.Sprintf("pendulum%d.txt", index+1))

    var counter int = 0
	for j:=0; j < len(c.pendulums[index].p_values); j++{
		if counter%33 == 0{
				
			var x1_value float64= length1 * float64(250) * math.Sin(c.pendulums[index].p_values[j].ang1) + float64(x_offset)
			var y1_value float64= length1 * float64(250) * math.Cos(c.pendulums[index].p_values[j].ang1) + float64(y_offset)
			var x2_value float64= length2 * float64(250) * math.Sin(c.pendulums[index].p_values[j].ang2) + float64(x1_value)
			var y2_value float64= length2 * float64(250) * math.Cos(c.pendulums[index].p_values[j].ang2) + float64(y1_value)
				
			f.WriteString(fmt.Sprintf("(%f,%f)|(%f,%f)\n", x1_value, y1_value,x2_value,y2_value))
				
			counter = 0
		}
		counter = counter + 1
			
	}
    f.Close()
}

func main(){	

	start := time.Now()

	for i := 0; i < pendulum_number; i++ {
    
        var angle1 = float64(math.Pi) / (float64(2.25) + (float64(i)* float64(0.05)))
        var angle2 = float64(math.Pi) / (float64(2.25) + (float64(i)* float64(0.05)))
        
        var pendulum_temp = Pendulum{angle1 : angle1, angle2 : angle2, velocity1: angle_velocity,
        velocity2 : angle_velocity, acceleration1 : angle_acceleration, acceleration2 : angle_acceleration,
        file_name : fmt.Sprintf("pendulum%d.txt", i+1)}
		
        pendulums = append(pendulums, pendulum_temp)
    	
	}
    
    duration0 := time.Since(start)
    fmt.Print("Part1: ")
	fmt.Print(duration0)
    fmt.Print("\n")
    
     
    var wg sync.WaitGroup
	safe_pendulum := SafePendulum{pendulums : pendulums}
	for i:= 0; i < pendulum_number; i++ {
        wg.Add(1)
        go safe_pendulum.calculate_p_values(&wg, i)
	}
    wg.Wait()
    
    duration1 := time.Since(start)
    fmt.Print("Part2: ")
	fmt.Print(duration1)
    fmt.Print("\n")
	
	for i:= 0; i < pendulum_number; i++ {
        wg.Add(1)
		go safe_pendulum.calculate_x_y_positions(&wg,i)
	}
    wg.Wait()
    
	duration := time.Since(start)
	fmt.Print("Program finished in ")
	fmt.Print(duration)
}