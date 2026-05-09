export function seatId(pass) {
  return parseInt(pass.replace(/[FL]/g, "0").replace(/[BR]/g, "1"), 2);
}
