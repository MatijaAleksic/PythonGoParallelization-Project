from math import *
from Serial.formular import *
from Serial.Pendulum import Pendulum

import time

from multiprocessing import Pool

#POCETAK IZVRSAVANJA
start_time = time.time()

#BROJ KLATNA
PENDULUM_NUM = 30
PROCESSES_NUM = 30

#parametri:
length1 = 1
length2 = 1
mass1 = 1
mass2 = 1

angle0 = math.pi / 2.25
angle1 = math.pi / 2.25

angle_velocity = 0
angle_acceleration = 0

#opcije:
fps = 30
g = 9.81
max_time = 60
make_for_30fps = 0.03333333 #0.03333333 #0.001  #za 60fps 0.01666667 #za 30fps 0.03333333
step = 0.001

pendulums = []
differ = 0.2

wisteph, height = 1928, 1020
starting_point = (int(wisteph/2), int(height/4))
x_offset = starting_point[0]
y_offset = starting_point[1]


def calculate_pendulum(pendulum):
    all_p_values = []
    all_p_values.append((pendulum.angle1, pendulum.angle2, pendulum.angle_velocity1, pendulum.angle_velocity2))
    #pendulum.p_values.append((pendulum.angle1, pendulum.angle2, pendulum.angle_velocity1, pendulum.angle_velocity2))
    p = all_p_values[0]
    t = 0
    while t < max_time:
        t += step
        alpha = p[0] + step * p[2]
        beta = p[1] + step * p[3]

        gamma = p[2] + step * FirstAcceleration(p[0], p[1], mass1, mass2, length1, length2, g, p[2], p[3])
        delta = p[3] + step * SecondAcceleration(p[0], p[1], mass1, mass2, length1, length2, g, p[2], p[3])

        new_angle1 = (p[0]+step*p[2])
        new_angle2 = (p[1]+step*p[3])
        new_angle_velocity1 = (p[2]+(step/2)*(FirstAcceleration(alpha, beta, mass1, mass2, length1, length2, g, gamma, delta) +
                                              FirstAcceleration(p[0], p[1], mass1, mass2, length1, length2, g, p[2], p[3])))
        new_angle_velocity2 = (p[3]+(step/2)*(SecondAcceleration(alpha, beta, mass1, mass2, length1, length2, g, gamma, delta) +
                                              SecondAcceleration(p[0], p[1], mass1, mass2, length1, length2, g, p[2], p[3])))

        p_new = (new_angle1, new_angle2, new_angle_velocity1, new_angle_velocity2)
        all_p_values.append(p_new)
        p = p_new

    pendulum.p_values = all_p_values
    return pendulum


if __name__ == '__main__':


    for i in range(PENDULUM_NUM):
        pendulums.append(Pendulum(mass1, mass2, length1, length2, angle0, angle1, angle_velocity, angle_velocity,
                                  angle_acceleration, angle_acceleration, f"pendulum{i}.txt"))
        angle1 = math.pi / (2.25+differ)
        angle2 = math.pi / (2.25+differ)
        differ = differ + 0.2

    serial1 = time.time() - start_time
    start_time = time.time()

    #PARALELIZACIJA
    with Pool(processes=PROCESSES_NUM) as p:
        pendulums = p.map(calculate_pendulum, pendulums)
        #p.map(calculate_x_y,pendulums)

    #PARALELNO KRAJ
    print("\n PARALEL PROGRAM FINISHED IN --- %s seconds ---" % (time.time() - start_time))
    start_time = time.time()

    for pendulum in pendulums:
        f = open(f"Data/{pendulum.file_name}", "w")
        counter = 0
        for p in pendulum.p_values:
            if counter == 33:
                x1_value = length1 * 250 * sin(p[0]) + x_offset
                y1_value = length1 * 250 * cos(p[0]) + y_offset
                x2_value = length2 * 250 * sin(p[1]) + x1_value
                y2_value = length2 * 250 * cos(p[1]) + y1_value

                f.write(f"({x1_value},{y1_value})|({x2_value},{y2_value})\n")
                counter = 0
            counter += 1
        f.close()


    serial2 = time.time() - start_time
    #SERIJSKO KRAJ
    print("\n SERIAL PROGRAM FINISHED IN --- %s seconds ---" % (serial1 + serial2))
