import z3

with open("days/10/input.txt") as f:
    res = 0
    for line in f.readlines():
        _, *pairs, target = line.split()
        buttons = [list(map(int, p[1:-1].split(","))) for p in pairs]
        joltage = list(map(int, target[1:-1].split(",")))
        # print(buttons, joltage)

        x = [z3.Int(f"x{i}") for i in range(len(buttons))]
        opt = z3.Optimize()

        # Non-negative presses
        for xi in x:
            opt.add(xi >= 0)

        for j in range(len(joltage)):
            total = z3.Sum([x[i] for i in range(len(buttons)) if j in buttons[i]])
            opt.add(total == joltage[j])

        # Minimize total button presses
        total_presses = z3.Sum(x)
        opt.minimize(total_presses)

        # Solve
        if opt.check() == z3.sat:
            m = opt.model()
            presses = [m[x[i]].as_long() for i in range(len(buttons))]
            total = sum(presses)
            res += total
            # print("Minimal number of button presses:", total)
            # print("Press counts:")
            # for i in range(len(buttons)):
            #     print(f"  Button {i}: {presses[i]} times")
    print(res)
