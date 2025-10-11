import matplotlib.pyplot as plt

data = [  # (Community, Masculinity, Maturity)
    ("Java", 3.2, 30),
    ("JavaScript", 3.6, 30),
    ("PHP", 5.5, 30),
    ("Ruby", 5.4, 30),
    ("Spring", 4.8, 23),
    (".Net", 5.7, 23),
    ("Go", 4.4, 16),
    ("TensorFlow", 3.2, 10),
    ("Swift", 4.5, 11),
    ("Flutter", 3.2, 7),
    ("React", 5.4, 12),
]

labls = [point[0] for point in data]
mascu = [point[1] for point in data]
matur = [point[2] for point in data]

plt.figure(figsize=(10, 6))

plt.title("Language maturity and masculinity among Kommunities")
plt.xlabel("Maturity (Years)")
plt.ylabel("Masculinity")
plt.grid(True, linestyle="--", alpha=0.6, zorder=0)
plt.scatter(matur, mascu, zorder=1)

for i, label in enumerate(labls):
    plt.text(matur[i] + 0.3, mascu[i], label, fontsize=9, va="center")

import os

os.makedirs("export", exist_ok=True)
plt.savefig("export/figure.png", dpi=300, facecolor="#ffffff")
