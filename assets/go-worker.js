/* In-browser Go judge: the Yaegi Go interpreter compiled to WebAssembly (tools/gojudge), in a Web Worker.
   Nothing is sent to a server. Same protocol as py-worker.js — the page terminates this worker on a time
   limit and starts a fresh one for the next run.
   Messages in:  {id, kind: 'judge', code, test}
   Messages out: {ready, error?} once; {id, started} when a run begins; {id, ok, result|error, dead?} */
importScripts('wasm_exec.js');

let log = '', dead = false;   // the Go runtime's own output (e.g. "fatal error: stack overflow")
const dec = new TextDecoder();
globalThis.fs.writeSync = (fd, buf) => { log = (log + dec.decode(buf)).slice(-4000); return buf.length; };

const text = p => fetch(new URL(p, self.location)).then(r => {
  if (!r.ok) throw new Error(`${p}: HTTP ${r.status}`);
  return r.text();
});
// the module is stored gzipped (GitHub Pages does not compress .wasm); the browser unpacks it as it downloads
const wasm = async p => {
  const r = await fetch(new URL(p, self.location));
  if (!r.ok) throw new Error(`${p}: HTTP ${r.status}`);
  return new Response(r.body.pipeThrough(new DecompressionStream('gzip'))).arrayBuffer();
};

const OVERFLOW = 'Stack overflow — infinite recursion? (The in-browser interpreter allows a few thousand nested calls.)';
const crash = (e = '') => /stack overflow|stack exceeds|call stack size/i.test(log + e) ? OVERFLOW
  : /out of memory/i.test(log + e) ? 'Out of memory'
  : 'The Go runtime crashed' + (log.trim() ? ': ' + log.trim().split('\n').slice(0, 2).join(' · ') : e ? ': ' + e : '');

const boot = (async () => {
  const [lc, bytes] = await Promise.all([text('../go/lc.go'), wasm('go-judge.wasm.gz')]);
  const go = new Go();
  const ready = new Promise(res => { self.goJudgeReady = res; });
  const { instance } = await WebAssembly.instantiate(bytes, go.importObject);
  go.run(instance).then(() => { dead = true; });
  await ready;
  return lc;
})();
boot.then(() => self.postMessage({ ready: true }), e => self.postMessage({ ready: false, error: String(e && e.message || e) }));

self.onmessage = async ({ data: { id, code, test } }) => {
  try {
    const lc = await boot;
    if (dead) throw new Error(crash());
    self.postMessage({ id, started: true });
    log = '';
    let res;
    try { res = self.goJudge(lc, code, test || ''); }
    catch (e) { dead = true; throw new Error(crash(String(e && e.message || e))); }   // the Go runtime's state is gone
    if (typeof res !== 'string') { dead = true; throw new Error(crash()); }
    self.postMessage({ id, ok: true, result: JSON.parse(res) });
  } catch (e) {
    self.postMessage({ id, ok: false, error: String(e && e.message || e), dead });
  }
};
