from helpers import parse_input
from ortools.sat.python import cp_model


def part2(puzzle_input):
    machines = parse_input(puzzle_input)
    result = 0

    for machine in machines:
        result += solve(machine)

    return int(result)


def solve(machine):
    buttons = machine.buttons
    target = machine.joltage

    n_buttons = len(buttons)
    n_counters = len(target)

    model = cp_model.CpModel()

    x = [model.NewIntVar(0, max(target), f"x{j}") for j in range(n_buttons)]

    for i in range(n_counters):
        model.Add(
            sum((1 if i in buttons[j] else 0) * x[j] for j in range(n_buttons))
            == target[i]
        )

    model.Minimize(sum(x))

    solver = cp_model.CpSolver()
    solver.parameters.max_time_in_seconds = 10
    solver.parameters.num_search_workers = 8

    res = solver.Solve(model)
    if res not in (cp_model.OPTIMAL, cp_model.FEASIBLE):
        return None
    return solver.ObjectiveValue()
