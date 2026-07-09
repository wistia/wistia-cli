#!/usr/bin/env node
// Assemble the npm packages from downloaded GitHub Release assets.
//
// Usage: node scripts/npm-stage.mjs --version <2026.M.N> --assets <dir> [--out dist/npm]
//
// Verifies every archive against checksums.txt, extracts each platform
// binary into a per-platform package, stamps the meta package's version and
// exact-pinned optionalDependencies, packs everything, and writes
// <out>/manifest.json in publish order (platform packages first, meta last).
// No network access; safe to run against any previously downloaded release.

import { spawnSync } from "node:child_process";
import crypto from "node:crypto";
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");

const TARGETS = [
  { asset: "wistia_Darwin_arm64.tar.gz", key: "darwin-arm64", os: "darwin", cpu: "arm64", label: "macOS ARM64" },
  { asset: "wistia_Darwin_x86_64.tar.gz", key: "darwin-x64", os: "darwin", cpu: "x64", label: "macOS x64" },
  { asset: "wistia_Linux_arm64.tar.gz", key: "linux-arm64", os: "linux", cpu: "arm64", label: "Linux ARM64" },
  { asset: "wistia_Linux_x86_64.tar.gz", key: "linux-x64", os: "linux", cpu: "x64", label: "Linux x64" },
  { asset: "wistia_Windows_arm64.zip", key: "win32-arm64", os: "win32", cpu: "arm64", label: "Windows ARM64" },
  { asset: "wistia_Windows_x86_64.zip", key: "win32-x64", os: "win32", cpu: "x64", label: "Windows x64" },
];

function fail(msg) {
  console.error(`npm-stage: ${msg}`);
  process.exit(1);
}

function parseArgs(argv) {
  const args = {};
  for (let i = 0; i < argv.length; i += 2) {
    const flag = argv[i];
    const value = argv[i + 1];
    if (!flag?.startsWith("--") || value === undefined) fail(`bad arguments near '${flag}'`);
    args[flag.slice(2)] = value;
  }
  return args;
}

function run(cmd, cmdArgs) {
  const result = spawnSync(cmd, cmdArgs, { encoding: "utf8" });
  if (result.status !== 0) {
    fail(`${cmd} ${cmdArgs.join(" ")} failed: ${result.stderr || result.stdout}`);
  }
  return result.stdout;
}

function sha256(file) {
  return crypto.createHash("sha256").update(fs.readFileSync(file)).digest("hex");
}

function parseChecksums(file) {
  const sums = new Map();
  for (const line of fs.readFileSync(file, "utf8").split("\n")) {
    const match = line.match(/^([0-9a-f]{64})\s+\*?(.+)$/);
    if (match) sums.set(match[2].trim(), match[1]);
  }
  return sums;
}

function readJson(file) {
  return JSON.parse(fs.readFileSync(file, "utf8"));
}

function writeJson(file, value) {
  fs.writeFileSync(file, `${JSON.stringify(value, null, 2)}\n`);
}

const args = parseArgs(process.argv.slice(2));
const version = args.version ?? fail("--version is required");
const assetsDir = path.resolve(args.assets ?? fail("--assets is required"));
const outDir = path.resolve(args.out ?? "dist/npm");

if (!/^\d{4}\.\d{1,2}\.\d+$/.test(version)) {
  fail(`version must look like 2026.5.1, got '${version}'`);
}

const checksumsFile = path.join(assetsDir, "checksums.txt");
if (!fs.existsSync(checksumsFile)) fail(`${checksumsFile} not found`);
const checksums = parseChecksums(checksumsFile);

const license = path.join(REPO_ROOT, "LICENSE");
if (!fs.existsSync(license)) fail("LICENSE not found at the repo root");
const platformTemplate = readJson(path.join(REPO_ROOT, "npm", "platform.package.json"));

fs.rmSync(outDir, { recursive: true, force: true });
const tarballsDir = path.join(outDir, "tarballs");
fs.mkdirSync(tarballsDir, { recursive: true });

const manifest = [];

for (const target of TARGETS) {
  const assetPath = path.join(assetsDir, target.asset);
  if (!fs.existsSync(assetPath)) fail(`missing asset ${target.asset}`);
  const expected = checksums.get(target.asset);
  if (!expected) fail(`no checksum entry for ${target.asset}`);
  const actual = sha256(assetPath);
  if (actual !== expected) {
    fail(`checksum mismatch for ${target.asset}: expected ${expected}, got ${actual}`);
  }

  const name = `@wistia/wistia-cli-${target.key}`;
  const pkgDir = path.join(outDir, `wistia-cli-${target.key}`);
  const binDir = path.join(pkgDir, "bin");
  fs.mkdirSync(binDir, { recursive: true });

  const exe = target.os === "win32" ? "wistia.exe" : "wistia";
  if (target.asset.endsWith(".zip")) {
    run("unzip", ["-j", "-o", assetPath, exe, "-d", binDir]);
  } else {
    run("tar", ["-xzf", assetPath, "-C", binDir, exe]);
  }
  fs.chmodSync(path.join(binDir, exe), 0o755);

  writeJson(path.join(pkgDir, "package.json"), {
    ...platformTemplate,
    name,
    version,
    description: `Platform-specific binary for @wistia/wistia-cli (${target.label})`,
    os: [target.os],
    cpu: [target.cpu],
  });
  fs.writeFileSync(
    path.join(pkgDir, "README.md"),
    `# ${name}\n\nPlatform-specific binary for [@wistia/wistia-cli](https://www.npmjs.com/package/@wistia/wistia-cli) (${target.label}).\nDo not install directly — install \`@wistia/wistia-cli\` instead.\n`,
  );
  fs.copyFileSync(license, path.join(pkgDir, "LICENSE"));

  manifest.push({ name, version, dir: pkgDir });
}

// Meta package last: exact-pin every platform package to this version.
const metaSrc = path.join(REPO_ROOT, "npm", "wistia-cli");
const metaDir = path.join(outDir, "wistia-cli");
fs.mkdirSync(path.join(metaDir, "bin"), { recursive: true });
fs.copyFileSync(path.join(metaSrc, "bin", "wistia"), path.join(metaDir, "bin", "wistia"));
fs.chmodSync(path.join(metaDir, "bin", "wistia"), 0o755);
fs.copyFileSync(path.join(metaSrc, "README.md"), path.join(metaDir, "README.md"));
fs.copyFileSync(license, path.join(metaDir, "LICENSE"));

const metaPkg = readJson(path.join(metaSrc, "package.json"));
metaPkg.version = version;
for (const dep of Object.keys(metaPkg.optionalDependencies)) {
  metaPkg.optionalDependencies[dep] = version;
}
writeJson(path.join(metaDir, "package.json"), metaPkg);
manifest.push({ name: metaPkg.name, version, dir: metaDir });

for (const entry of manifest) {
  const stdout = run("npm", ["pack", entry.dir, "--pack-destination", tarballsDir, "--silent"]);
  entry.tarball = path.join(tarballsDir, stdout.trim().split("\n").pop());
}

writeJson(path.join(outDir, "manifest.json"), manifest);
console.log(JSON.stringify(manifest, null, 2));
