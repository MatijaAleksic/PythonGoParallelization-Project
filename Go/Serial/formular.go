package formular

import (
	"math"
)

func FirstAcceleration(t1,t2,m1,m2,L1,L2,G,v1,v2 float64) float64{
	var numerator1 float64= -G * (2 * m1 + m2) * math.Sin(t1)
    var numerator2 float64= -m2 * G * math.Sin(t1 - 2 * t2)
    var numerator3 float64= -2 * math.Sin(t1-t2)
    var numerator4 float64= m2 * ((v2 * v2) * L2 + (v1 * v1) * L1 * math.Cos(t1-t2))
    var numerator float64= numerator1 + numerator2 + (numerator3 * numerator4)
    var denominator float64= L1 * (2 * m1 + m2 - m2 * math.Cos(2 * t1 - 2 * t2))

    return float64(numerator/denominator)

}

func SecondAcceleration(t1, t2, m1, m2, L1, L2, G, v1, v2 float64)float64{
    var numerator1 float64= 2 * math.Sin(t1 - t2)
    var numerator2 float64= (v1 * v1) * L1 * (m1 + m2) + G * (m1 + m2) * math.Cos(t1)
    var numerator3 float64= (v2 * v2) * L2 * m2 * math.Cos(t1-t2)

    var numerator float64= numerator1 * (numerator2 + numerator3)
    var denominator float64= L2 * (2 * m1 + m2 - m2 * math.Cos(2 * t1 - 2 * t2))

    return float64(numerator/denominator)
}