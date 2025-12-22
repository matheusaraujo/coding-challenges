import crypto from "node:crypto";

export function getOpenDoors(passcode, path) {
  const hash = crypto.createHash("md5").update(passcode + path).digest("hex");
  return [
    "bcdef".includes(hash[0]),
    "bcdef".includes(hash[1]),
    "bcdef".includes(hash[2]),
    "bcdef".includes(hash[3]),
  ];
}
