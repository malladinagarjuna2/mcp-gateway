import os

paths = ['internal/controller', 'cmd/main.go']

for p in paths:
    if os.path.isfile(p):
        with open(p, 'r') as f:
            c = f.read()
        if 'v1alpha1' in c:
            with open(p, 'w') as f:
                f.write(c.replace('v1alpha1', 'v1'))
    else:
        for root, dirs, files in os.walk(p):
            for file in files:
                if file.endswith('.go'):
                    path = os.path.join(root, file)
                    with open(path, 'r') as f:
                        c = f.read()
                    if 'v1alpha1' in c:
                        with open(path, 'w') as f:
                            f.write(c.replace('v1alpha1', 'v1'))
