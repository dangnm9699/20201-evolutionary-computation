import sys

import matplotlib.pyplot as plt
import numpy as np

file = sys.argv[1]
path = sys.argv[2]

path_list = path.split(" ")
path_list = np.array(path_list, dtype=int)

with open(file, "r") as con:
    lines = con.readlines()
    xy = []
    for line in lines:
        as_list = line.split(" ")
        xy.append([as_list[0], as_list[1], as_list[2].replace("\n", "")])

xy = np.array(xy, dtype=int)

plt.figure(figsize=(15, 5))
plt.scatter(xy[:, 1], xy[:, 2])
for i in range(len(xy)):
    plt.annotate(str(i + 1), (xy[i][1], xy[i][2]), color='r', size='8')
arr_x = []
arr_y = []
for i in range(len(path_list)):
    cur = path_list[i]
    arr_x.append(xy[cur - 1][1])
    arr_y.append(xy[cur - 1][2])
arr_x.append(xy[path_list[0] - 1][1])
arr_y.append(xy[path_list[0] - 1][2])
arr_x = np.array(arr_x)
arr_y = np.array(arr_y)
plt.quiver(arr_x[:-1], arr_y[:-1], arr_x[1:] - arr_x[:-1], arr_y[1:] - arr_y[:-1], scale_units='xy', angles='xy',
           scale=1, width=0.0015)
plt.show()
