export function fillDisk(a, size) {
  while (a.length < size) {
    const b = a
      .split("")
      .reverse()
      .map((c) => (c === "1" ? "0" : "1"))
      .join("");
    a = a + "0" + b;
  }
  return a.substring(0, size);
}

export function checksum(disk) {
  while (disk.length % 2 === 0) {
    let c = "";
    for (let i = 0; i < disk.length; i += 2) {
      c += disk[i] === disk[i + 1] ? "1" : "0";
    }
    disk = c;
  }

  return disk;
}
