# pylint: disable=too-many-locals,invalid-name
import itertools
import re

from helpers import Intcode

DANGEROUS = {
    "infinite loop",
    "molten lava",
    "photons",
    "escape pod",
    "giant electromagnet",
}

OPPOSITE = {"north": "south", "south": "north", "east": "west", "west": "east"}


def read_text(vm):
    return "".join(chr(c) if c < 128 else "?" for c in vm.run())


def send(vm, command):
    vm.feed(command + "\n")
    return read_text(vm)


def parse_room(text):
    headings = list(re.finditer(r"== ([^=]+) ==", text))
    if not headings:
        return None
    last = headings[-1]
    name = last.group(1).strip()
    body = text[last.end() :]

    exits = re.findall(r"^- (north|south|east|west)$", body, re.MULTILINE)

    items = []
    items_section = re.search(r"Items here:\n((?:- .+\n?)+)", body)
    if items_section:
        items = [m.strip() for m in re.findall(r"- (.+)", items_section.group(1))]

    return {"name": name, "exits": exits, "items": items}


def part1(puzzle_input):
    program = [int(x) for x in puzzle_input[0].split(",")]
    vm = Intcode(program)

    starting_room = parse_room(read_text(vm))

    visited = {starting_room["name"]}
    inventory = []
    path_to_checkpoint = []

    def explore(current, path):
        nonlocal path_to_checkpoint

        for item in current["items"]:
            if item not in DANGEROUS:
                send(vm, "take " + item)
                inventory.append(item)

        if current["name"] == "Security Checkpoint":
            path_to_checkpoint = list(path)
            return

        for exit_dir in current["exits"]:
            new_room = parse_room(send(vm, exit_dir))
            if new_room is None:
                continue
            if new_room["name"] in visited:
                if new_room["name"] != current["name"]:
                    send(vm, OPPOSITE[exit_dir])
                continue
            visited.add(new_room["name"])
            explore(new_room, path + [exit_dir])
            send(vm, OPPOSITE[exit_dir])

    explore(starting_room, [])

    text = ""
    for direction in path_to_checkpoint:
        text = send(vm, direction)
    checkpoint = parse_room(text)

    came_from = OPPOSITE[path_to_checkpoint[-1]]
    floor_dir = next(d for d in checkpoint["exits"] if d != came_from)

    for item in inventory:
        send(vm, "drop " + item)

    for size in range(1, len(inventory) + 1):
        for subset in itertools.combinations(inventory, size):
            for item in subset:
                send(vm, "take " + item)
            response = send(vm, floor_dir)
            match = re.search(r"by typing (\d+)", response)
            if match:
                return int(match.group(1))
            for item in subset:
                send(vm, "drop " + item)

    return None
