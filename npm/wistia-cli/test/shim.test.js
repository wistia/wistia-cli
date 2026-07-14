"use strict";

const { test } = require("node:test");
const assert = require("node:assert/strict");
const { spawnSync } = require("node:child_process");
const path = require("node:path");

const SHIM = path.join(__dirname, "..", "bin", "wistia");
const shim = require(SHIM);

// The map must resolve to the packages the meta package actually declares —
// a drifted entry only surfaces as a broken install on that platform, and
// four of the six platforms are never executed by any test environment.
test("platform map matches the meta package optionalDependencies", () => {
  const deps = Object.keys(require("../package.json").optionalDependencies);
  assert.deepEqual(Object.values(shim.PLATFORMS).sort(), deps.sort());
  for (const [key, pkg] of Object.entries(shim.PLATFORMS)) {
    assert.equal(pkg, `@wistia/wistia-cli-${key}`);
  }
});

test("WISTIA_CLI_BINARY overrides resolution", () => {
  const bin = shim.binaryPath({
    platform: "linux",
    arch: "x64",
    env: { WISTIA_CLI_BINARY: "/opt/custom/wistia" },
    resolve: () => {
      throw new Error("resolve must not be called");
    },
  });
  assert.equal(bin, "/opt/custom/wistia");
});

test("resolves the matching platform package binary", () => {
  const bin = shim.binaryPath({
    platform: "darwin",
    arch: "arm64",
    env: {},
    resolve: (spec) => `/resolved/${spec}`,
  });
  assert.equal(bin, "/resolved/@wistia/wistia-cli-darwin-arm64/bin/wistia");
});

test("win32 resolves the .exe binary", () => {
  const bin = shim.binaryPath({
    platform: "win32",
    arch: "x64",
    env: {},
    resolve: (spec) => `/resolved/${spec}`,
  });
  assert.equal(bin, "/resolved/@wistia/wistia-cli-win32-x64/bin/wistia.exe");
});

test("unsupported platform fails with install pointers", () => {
  assert.throws(
    () => shim.binaryPath({ platform: "sunos", arch: "x64", env: {}, resolve: () => "" }),
    (err) => {
      assert.match(err.message, /unsupported platform sunos-x64/);
      assert.match(err.message, /cli-installation/);
      return true;
    },
  );
});

test("missing platform package fails with reinstall guidance", () => {
  assert.throws(
    () =>
      shim.binaryPath({
        platform: "linux",
        arch: "arm64",
        env: {},
        resolve: () => {
          throw new Error("not found");
        },
      }),
    (err) => {
      assert.match(err.message, /omit=optional/);
      assert.match(err.message, /cli-installation/);
      return true;
    },
  );
});

test("spawn path forwards argv, stdout, and exit code", () => {
  // exit inside the write callback: piped stdout may not flush past exit()
  const result = spawnSync(
    process.execPath,
    [SHIM, "-e", "process.stdout.write('passthrough-ok\\n', () => process.exit(7))"],
    {
      env: { ...process.env, WISTIA_CLI_BINARY: process.execPath },
      encoding: "utf8",
    },
  );
  assert.equal(result.status, 7);
  assert.match(result.stdout, /passthrough-ok/);
});

test("unresolvable platform package exits 1 with the shim error", () => {
  const env = { ...process.env };
  delete env.WISTIA_CLI_BINARY;
  const result = spawnSync(process.execPath, [SHIM, "version"], {
    env,
    encoding: "utf8",
  });
  assert.equal(result.status, 1);
  assert.match(result.stderr, /is not installed|unsupported platform/);
});
