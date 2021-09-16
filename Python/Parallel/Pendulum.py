class Pendulum:
    def __init__(self, mass1, mass2, length1, length2, angle1, angle2, angle_velocity1, angle_velocity2,
                 angle_acceleration1, angle_acceleration2, file_name):

        self.mass1 = mass1
        self.mass2 = mass2

        self.length1 = length1
        self.length2 = length2

        self.angle1 = angle1
        self.angle2 = angle2

        self.angle_velocity1 = angle_velocity1
        self.angle_velocity2 = angle_velocity2

        self.angle_acceleration1 = angle_acceleration1
        self.angle_acceleration2 = angle_acceleration2

        self.p_values = []
        self.scatter1 = []
        self.scatter2 = []

        self.file_name = file_name
